package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	"github.com/expert-pancake/service/purchasing/impl/client"
	"github.com/expert-pancake/service/purchasing/util"
)

type UpdatePurchaseOrderStatusTrxParams struct {
	PurchaseOrderId string
	Status          string
}

type UpdatePurchaseOrderStatusTrxResult struct {
	Message string
}

func (trx *Trx) UpdatePurchaseOrderStatusTrx(ctx context.Context, arg UpdatePurchaseOrderStatusTrxParams) (UpdatePurchaseOrderStatusTrxResult, error) {
	var result UpdatePurchaseOrderStatusTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		purchaseOrder, err := q.GetPurchaseOrder(context.Background(), arg.PurchaseOrderId)
		if err != nil {
			return err
		}

		result.Message = "OK"

		if !purchaseOrder.IsDeleted {
			if purchaseOrder.Status != arg.Status && purchaseOrder.Status == "created" {

				err := q.UpdatePurchaseOrderStatus(context.Background(), db.UpdatePurchaseOrderStatusParams{
					ID:     arg.PurchaseOrderId,
					Status: arg.Status,
				})
				if err != nil {
					return err
				}

				if arg.Status == "accepted" && purchaseOrder.SecondaryCompanyID != "" && purchaseOrder.KonekinID != "" {
					konekin, err := client.GetKonekinContactBook(
						client.GetKonekinContactBookRequest{
							PrimaryCompanyId:   purchaseOrder.SecondaryCompanyID,
							SecondaryCompanyId: purchaseOrder.CompanyID,
						})
					if err != nil {
						return err
					}

					if konekin.Result.IsCustomer {
						supplierBranchId := ""

						if (!konekin.Result.IsAllBranches && len(konekin.Result.Branches) == 0) || konekin.Result.IsAllBranches {
							branches, err := client.GetCompanyBranches(
								client.GetCompanyBranchesRequest{
									CompanyId: purchaseOrder.SecondaryCompanyID,
								})
							if err != nil {
								return err
							}

							supplierBranchId = branches.Result[0].BranchId
						} else if !konekin.Result.IsAllBranches && len(konekin.Result.Branches) == 1 {
							supplierBranchId = konekin.Result.Branches[0]
						}

						salesOrder, err := client.UpsertSalesOrder(
							client.UpsertSalesOrderRequest{
								PurchaseOrderId:       purchaseOrder.ID,
								PurchaseOrderBranchId: purchaseOrder.BranchID,
								CompanyId:             purchaseOrder.SecondaryCompanyID,
								BranchId:              supplierBranchId,
								TransactionDate:       purchaseOrder.TransactionDate.Format(util.DateLayoutYMD),
								ContactBookId:         konekin.Result.ContactBookId,
								SecondaryCompanyId:    purchaseOrder.CompanyID,
								KonekinId:             konekin.Result.KonekinId,
								CurrencyCode:          purchaseOrder.CurrencyCode,
								IsAllBranches:         konekin.Result.IsAllBranches,
								Branches:              konekin.Result.Branches,
							})
						if err != nil {
							return err
						}

						err = q.UpdateAcceptedPurchaseOrder(context.Background(), db.UpdateAcceptedPurchaseOrderParams{
							ID:           arg.PurchaseOrderId,
							SalesOrderID: salesOrder.Result.TransactionId,
						})
						if err != nil {
							return err
						}

						purchaseOrderItems, err := q.GetPurchaseOrderItems(context.Background(), arg.PurchaseOrderId)
						if err != nil {
							return err
						}

						var salesOrderItemsReq = make([]client.SalesOrderItemsRequest, 0)

						for _, d := range purchaseOrderItems {

							var salesOrderItemReq = client.SalesOrderItemsRequest{
								PurchaseOrderItemId:    d.ID,
								PrimaryItemVariantId:   d.SecondaryItemVariantID,
								SecondaryItemVariantId: d.PrimaryItemVariantID,
								PrimaryItemUnitId:      d.SecondaryItemUnitID,
								SecondaryItemUnitId:    d.PrimaryItemUnitID,
								PrimaryItemUnitValue:   strconv.FormatInt(d.SecondaryItemUnitValue, 10),
								SecondaryItemUnitValue: strconv.FormatInt(d.PrimaryItemUnitValue, 10),
								Amount:                 strconv.FormatInt(d.Amount, 10),
								Price:                  strconv.FormatInt(d.Price, 10),
							}
							salesOrderItemsReq = append(salesOrderItemsReq, salesOrderItemReq)
						}

						salesOrderItems, err := client.UpdateSalesOrderItems(
							client.UpdateSalesOrderItemsRequest{
								SalesOrderId:    salesOrder.Result.TransactionId,
								SalesOrderItems: salesOrderItemsReq,
							})
						if err != nil {
							return err
						}

						for _, d := range salesOrderItems.Result.SalesOrderItems {
							err = q.UpdateAcceptedPurchaseOrderItem(context.Background(), db.UpdateAcceptedPurchaseOrderItemParams{
								ID:               d.PurchaseOrderItemId,
								SalesOrderItemID: d.DetailId,
							})
							if err != nil {
								return err
							}
						}
					}
				}
			} else {
				result.Message = "No data updated"
			}
		} else {
			result.Message = "Transaction has been deleted"
		}

		return err
	})

	return result, err
}
