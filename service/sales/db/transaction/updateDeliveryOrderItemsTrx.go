package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
	uuid "github.com/satori/go.uuid"
)

type UpdateDeliveryOrderItemsTrxParams struct {
	DeliveryOrderId    string
	DeliveryOrderItems []model.DeliveryOrderItemsRequest
}

type UpdateDeliveryOrderItemsTrxResult struct {
	DeliveryOrderItems []model.DeliveryOrderItem
}

func (trx *Trx) UpdateDeliveryOrderItemsTrx(ctx context.Context, arg UpdateDeliveryOrderItemsTrxParams) (UpdateDeliveryOrderItemsTrxResult, error) {
	var result UpdateDeliveryOrderItemsTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		err = q.DeleteDeliveryOrderItems(ctx, arg.DeliveryOrderId)
		if err != nil {
			return err
		}

		for _, d := range arg.DeliveryOrderItems {
			primaryItemUnitValue, _ := strconv.ParseInt(d.PrimaryItemUnitValue, 10, 64)
			secondaryItemUnitValue, _ := strconv.ParseInt(d.SecondaryItemUnitValue, 10, 64)
			amount, _ := strconv.ParseInt(d.Amount, 10, 64)

			// _, err = q.UpdateSalesOrderItemAmountSent(ctx, db.UpdateSalesOrderItemAmountSentParams{
			// 	ID:         d.SalesOrderItemId,
			// 	AmountSent: amount,
			// })
			// if err != nil {
			// 	return err
			// }

			headerRes, err := q.UpsertDeliveryOrderItem(ctx, db.UpsertDeliveryOrderItemParams{
				ID:                     uuid.NewV4().String(),
				PurchaseOrderItemID:    d.PurchaseOrderItemId,
				SalesOrderItemID:       d.SalesOrderItemId,
				DeliveryOrderID:        arg.DeliveryOrderId,
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
				Amount:                 amount,
			})
			if err != nil {
				return err
			}

			var deliveryOrderItem = model.DeliveryOrderItem{
				DetailId:               headerRes.ID,
				PurchaseOrderItemId:    headerRes.PurchaseOrderItemID,
				SalesOrderItemId:       headerRes.SalesOrderItemID,
				DeliveryOrderId:        headerRes.DeliveryOrderID,
				PrimaryItemVariantId:   headerRes.PrimaryItemVariantID,
				WarehouseRackId:        headerRes.WarehouseRackID,
				Batch:                  d.Batch,
				ExpiredDate:            d.ExpiredDate,
				ItemBarcodeId:          headerRes.ItemBarcodeID,
				SecondaryItemVariantId: headerRes.SecondaryItemVariantID,
				PrimaryItemUnitId:      headerRes.PrimaryItemUnitID,
				SecondaryItemUnitId:    headerRes.SecondaryItemUnitID,
				PrimaryItemUnitValue:   strconv.FormatInt(headerRes.PrimaryItemUnitValue, 10),
				SecondaryItemUnitValue: strconv.FormatInt(headerRes.SecondaryItemUnitValue, 10),
				Amount:                 strconv.FormatInt(headerRes.Amount, 10),
			}
			result.DeliveryOrderItems = append(result.DeliveryOrderItems, deliveryOrderItem)
		}

		err = q.UpdateDeliveryOrderTotalItems(ctx, db.UpdateDeliveryOrderTotalItemsParams{
			ID:         arg.DeliveryOrderId,
			TotalItems: int64(len(arg.DeliveryOrderItems)),
		})
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
