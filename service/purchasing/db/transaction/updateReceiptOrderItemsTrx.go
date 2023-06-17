package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	"github.com/expert-pancake/service/purchasing/model"
	"github.com/expert-pancake/service/purchasing/util"
	uuid "github.com/satori/go.uuid"
)

type UpdateReceiptOrderItemsTrxParams struct {
	ReceiptOrderId    string
	ReceiptOrderItems []model.ReceiptOrderItemsRequest
}

func (trx *Trx) UpdateReceiptOrderItemsTrx(ctx context.Context, arg UpdateReceiptOrderItemsTrxParams) (error) {
	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		err = q.DeleteReceiptOrderItems(ctx, arg.ReceiptOrderId)
		if err != nil {
			return err
		}

		for _, d := range arg.ReceiptOrderItems {
			primaryItemUnitValue, _ := strconv.ParseInt(d.PrimaryItemUnitValue, 10, 64)
			secondaryItemUnitValue, _ := strconv.ParseInt(d.SecondaryItemUnitValue, 10, 64)
			amount, _ := strconv.ParseInt(d.Amount, 10, 64)
			amountDelivered, _ := strconv.ParseInt(d.AmountDelivered, 10, 64)

			err := q.InsertReceiptOrderItem(ctx, db.InsertReceiptOrderItemParams{
				ID:                     uuid.NewV4().String(),
				ReceiptOrderID:         arg.ReceiptOrderId,
				PurchaseOrderItemID:    d.PurchaseOrderItemId,
				SalesOrderItemID:       d.SalesOrderItemId,
				DeliveryOrderItemID:    d.DeliveryOrderItemId,
				PrimaryItemVariantID:   d.PrimaryItemVariantId,
				WarehouseRackID:        d.WarehouseRackId,
				Batch:                  util.NewNullableString(d.Batch),
				ExpiredDate:            util.NewNullableDate(util.StringToDate(d.ExpiredDate)),
				ItemBarcodeID:          d.ItemBarcodeId,
				SecondaryItemVariantID: d.SecondaryItemVariantId,
				PrimaryItemUnitID:      d.PrimaryItemUnitId,
				SecondaryItemUnitID:    d.SecondaryItemUnitId,
				PrimaryItemUnitValue:   primaryItemUnitValue,
				SecondaryItemUnitValue: secondaryItemUnitValue,
				AmountDelivered:        amountDelivered,
				Amount:                 amount,
			})
			if err != nil {
				return err
			}
		}

		return err
	})

	return err
}
