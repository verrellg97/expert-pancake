package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	"github.com/expert-pancake/service/purchasing/impl/client"
	"github.com/expert-pancake/service/purchasing/util"
)

type UpdateReceiptOrderStatusTrxParams struct {
	ReceiptOrderId string
	Status         string
}

type UpdateReceiptOrderStatusTrxResult struct {
	Message string
}

func (trx *Trx) UpdateReceiptOrderStatusTrx(ctx context.Context, arg UpdateReceiptOrderStatusTrxParams) (UpdateReceiptOrderStatusTrxResult, error) {
	var result UpdateReceiptOrderStatusTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		receiptOrder, err := q.GetReceiptOrder(context.Background(), arg.ReceiptOrderId)
		if err != nil {
			return err
		}

		result.Message = "OK"

		if !receiptOrder.IsDeleted && receiptOrder.WarehouseID != "" {
			if receiptOrder.Status != arg.Status && receiptOrder.Status == "created" {

				err := q.UpdateReceiptOrderStatus(context.Background(), db.UpdateReceiptOrderStatusParams{
					ID:     arg.ReceiptOrderId,
					Status: arg.Status,
				})
				if err != nil {
					return err
				}

				if arg.Status == "accepted" {
					receiptOrderItems, err := q.GetReceiptOrderItems(context.Background(), arg.ReceiptOrderId)
					if err != nil {
						return err
					}

					deleteStock, err := client.DeleteStockMovement(
						client.DeleteStockMovementRequest{
							TransactionId:        receiptOrder.ID,
							TransactionReference: "RO",
						})
					if err != nil || deleteStock.Result.Message != "OK" {
						return err
					}

					for _, d := range receiptOrderItems {

						insertStock, err := client.InsertStockMovement(
							client.InsertStockMovementRequest{
								TransactionId:        arg.ReceiptOrderId,
								CompanyId:            receiptOrder.CompanyID,
								BranchId:             receiptOrder.BranchID,
								TransactionCode:      receiptOrder.FormNumber,
								TransactionDate:      receiptOrder.TransactionDate.Format(util.DateLayoutYMD),
								TransactionReference: "RO",
								DetailTransactionId:  d.ID,
								WarehouseId:          receiptOrder.WarehouseID,
								WarehouseRackId:      d.WarehouseRackID,
								VariantId:            d.PrimaryItemVariantID,
								ItemBarcodeId:        d.ItemBarcodeID,
								Amount:               strconv.FormatInt(d.Amount*d.PrimaryItemUnitValue, 10),
							})
						if err != nil || insertStock.Result.Message != "OK" {
							result.Message = "Warehouse rack required"
							return err
						}

						err = q.UpdatePurchaseOrderItemAmountReceived(context.Background(), db.UpdatePurchaseOrderItemAmountReceivedParams{
							ID:             d.PurchaseOrderItemID,
							AmountReceived: d.Amount,
						})
						if err != nil {
							return err
						}
					}
				}
			} else {
				result.Message = "No data updated"
			}
		} else {
			if receiptOrder.IsDeleted {
				result.Message = "Transaction has been deleted"
			} else {
				result.Message = "Warehouse required"
			}
		}

		return err
	})

	return result, err
}
