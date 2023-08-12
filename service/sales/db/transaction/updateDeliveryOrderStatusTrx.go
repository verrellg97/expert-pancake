package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/util"
)

type UpdateDeliveryOrderStatusTrxParams struct {
	DeliveryOrderId string
	Status          string
}

type UpdateDeliveryOrderStatusTrxResult struct {
	Message string
}

func (trx *Trx) UpdateDeliveryOrderStatusTrx(ctx context.Context, arg UpdateDeliveryOrderStatusTrxParams) (UpdateDeliveryOrderStatusTrxResult, error) {
	var result UpdateDeliveryOrderStatusTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		deliveryOrder, err := q.GetDeliveryOrder(context.Background(), arg.DeliveryOrderId)
		if err != nil {
			return err
		}

		salesOrder, err := q.GetSalesOrder(context.Background(), deliveryOrder.SalesOrderID)
		if err != nil {
			return err
		}

		result.Message = "OK"

		if !deliveryOrder.IsDeleted {
			if deliveryOrder.Status != arg.Status && deliveryOrder.Status == "created" {

				err := q.UpdateDeliveryOrderStatus(context.Background(), db.UpdateDeliveryOrderStatusParams{
					ID:     arg.DeliveryOrderId,
					Status: arg.Status,
				})
				if err != nil {
					return err
				}

				if arg.Status == "accepted" && deliveryOrder.SecondaryCompanyID != "" && deliveryOrder.KonekinID != "" {
					konekin, err := client.GetKonekinContactBook(
						client.GetKonekinContactBookRequest{
							PrimaryCompanyId:   deliveryOrder.SecondaryCompanyID,
							SecondaryCompanyId: deliveryOrder.CompanyID,
						})
					if err != nil {
						return err
					}

					deliveryOrderItems, err := q.GetDeliveryOrderItems(context.Background(), arg.DeliveryOrderId)
					if err != nil {
						return err
					}

					var deliveryOrderItemsReq = make([]client.UpsertReceiptOrderItemsRequest, 0)

					for _, d := range deliveryOrderItems {
						var deliveryOrderItemReq = client.UpsertReceiptOrderItemsRequest{
							PurchaseOrderItemId:    d.PurchaseOrderItemID,
							SalesOrderItemId:       d.SalesOrderItemID,
							DeliveryOrderItemId:    d.ID,
							PrimaryItemVariantId:   d.SecondaryItemVariantID,
							WarehouseRackId:        "",
							Batch:                  d.Batch.String,
							ExpiredDate:            d.ExpiredDate.Time.String(),
							ItemBarcodeId:          d.ItemBarcodeID,
							SecondaryItemVariantId: d.PrimaryItemVariantID,
							PrimaryItemUnitId:      d.SecondaryItemUnitID,
							SecondaryItemUnitId:    d.PrimaryItemUnitID,
							PrimaryItemUnitValue:   strconv.FormatInt(d.SecondaryItemUnitValue, 10),
							SecondaryItemUnitValue: strconv.FormatInt(d.PrimaryItemUnitValue, 10),
							Amount:                 strconv.FormatInt(d.Amount, 10),
						}
						err = q.UpdateSalesOrderItemAmountSent(context.Background(), db.UpdateSalesOrderItemAmountSentParams{
							ID:         d.SalesOrderItemID,
							AmountSent: d.Amount,
						})
						if err != nil {
							return err
						}
						deliveryOrderItemsReq = append(deliveryOrderItemsReq, deliveryOrderItemReq)
					}

					_, err = client.UpsertReceiptOrder(
						client.UpsertReceiptOrderRequest{
							DeliveryOrderId:                deliveryOrder.ID,
							WarehouseId:                    salesOrder.PurchaseOrderReceivingWarehouseID,
							CompanyId:                      deliveryOrder.SecondaryCompanyID,
							BranchId:                       salesOrder.PurchaseOrderBranchID,
							TransactionDate:                deliveryOrder.TransactionDate.Format(util.DateLayoutYMD),
							ContactBookId:                  konekin.Result.ContactBookId,
							SecondaryCompanyId:             deliveryOrder.CompanyID,
							KonekinId:                      deliveryOrder.KonekinID,
							TotalItems:                     strconv.FormatInt(int64(len(deliveryOrderItemsReq)), 10),
							UpsertReceiptOrderItemsRequest: deliveryOrderItemsReq,
						})
					if err != nil {
						return err
					}

					// for _, d := range salesOrderItems.Result.SalesOrderItems {
					// 	err = q.UpdateAcceptedPurchaseOrderItem(context.Background(), db.UpdateAcceptedPurchaseOrderItemParams{
					// 		ID:               d.PurchaseOrderItemId,
					// 		SalesOrderItemID: d.DetailId,
					// 	})
					// 	if err != nil {
					// 		return err
					// 	}
					// }
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
