package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	"github.com/expert-pancake/service/purchasing/model"
	uuid "github.com/satori/go.uuid"
)

type UpdatePurchaseOrderItemsTrxParams struct {
	PurchaseOrderId    string
	PurchaseOrderItems []model.PurchaseOrderItemsRequest
}

type UpdatePurchaseOrderItemsTrxResult struct {
	PurchaseOrderItems []model.PurchaseOrderItem
}

func (trx *Trx) UpdatePurchaseOrderItemsTrx(ctx context.Context, arg UpdatePurchaseOrderItemsTrxParams) (UpdatePurchaseOrderItemsTrxResult, error) {
	var result UpdatePurchaseOrderItemsTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		err = q.DeletePurchaseOrderItems(ctx, arg.PurchaseOrderId)
		if err != nil {
			return err
		}

		for _, d := range arg.PurchaseOrderItems {
			primaryItemUnitValue, _ := strconv.ParseInt(d.PrimaryItemUnitValue, 10, 64)
			secondaryItemUnitValue, _ := strconv.ParseInt(d.SecondaryItemUnitValue, 10, 64)
			amount, _ := strconv.ParseInt(d.Amount, 10, 64)
			price, _ := strconv.ParseInt(d.Price, 10, 64)

			headerRes, err := q.UpsertPurchaseOrderItem(ctx, db.UpsertPurchaseOrderItemParams{
				ID:                     uuid.NewV4().String(),
				PurchaseOrderID:        arg.PurchaseOrderId,
				PrimaryItemVariantID:   d.PrimaryItemVariantId,
				SecondaryItemVariantID: d.SecondaryItemVariantId,
				PrimaryItemUnitID:      d.PrimaryItemUnitId,
				SecondaryItemUnitID:    d.SecondaryItemUnitId,
				PrimaryItemUnitValue:   primaryItemUnitValue,
				SecondaryItemUnitValue: secondaryItemUnitValue,
				Amount:                 amount,
				Price:                  price,
			})
			if err != nil {
				return err
			}

			var purchaseOrderItem = model.PurchaseOrderItem{
				DetailId:               headerRes.ID,
				PurchaseOrderId:        headerRes.PurchaseOrderID,
				PrimaryItemVariantId:   headerRes.PrimaryItemVariantID,
				SecondaryItemVariantId: headerRes.SecondaryItemVariantID,
				PrimaryItemUnitId:      headerRes.PrimaryItemUnitID,
				SecondaryItemUnitId:    headerRes.SecondaryItemUnitID,
				PrimaryItemUnitValue:   strconv.FormatInt(headerRes.PrimaryItemUnitValue, 10),
				SecondaryItemUnitValue: strconv.FormatInt(headerRes.SecondaryItemUnitValue, 10),
				Amount:                 strconv.FormatInt(headerRes.Amount, 10),
				Price:                  strconv.FormatInt(headerRes.Price, 10),
			}
			result.PurchaseOrderItems = append(result.PurchaseOrderItems, purchaseOrderItem)
		}

		err = q.UpdatePurchaseOrderAddItem(ctx, arg.PurchaseOrderId)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
