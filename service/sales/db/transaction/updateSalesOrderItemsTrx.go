package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/model"
	uuid "github.com/satori/go.uuid"
)

type UpdateSalesOrderItemsTrxParams struct {
	SalesOrderId    string
	SalesOrderItems []model.SalesOrderItemsRequest
}

type UpdateSalesOrderItemsTrxResult struct {
	SalesOrderItems []model.SalesOrderItem
}

func (trx *Trx) UpdateSalesOrderItemsTrx(ctx context.Context, arg UpdateSalesOrderItemsTrxParams) (UpdateSalesOrderItemsTrxResult, error) {
	var result UpdateSalesOrderItemsTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		err = q.DeleteSalesOrderItems(ctx, arg.SalesOrderId)
		if err != nil {
			return err
		}

		for _, d := range arg.SalesOrderItems {
			primaryItemUnitValue, _ := strconv.ParseInt(d.PrimaryItemUnitValue, 10, 64)
			secondaryItemUnitValue, _ := strconv.ParseInt(d.SecondaryItemUnitValue, 10, 64)
			amount, _ := strconv.ParseInt(d.Amount, 10, 64)
			price, _ := strconv.ParseInt(d.Price, 10, 64)

			headerRes, err := q.UpsertSalesOrderItem(ctx, db.UpsertSalesOrderItemParams{
				ID:                     uuid.NewV4().String(),
				PurchaseOrderItemID:    d.PurchaseOrderItemId,
				SalesOrderID:           arg.SalesOrderId,
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

			var salesOrderItem = model.SalesOrderItem{
				DetailId:               headerRes.ID,
				PurchaseOrderItemId:    headerRes.PurchaseOrderItemID,
				SalesOrderId:           headerRes.SalesOrderID,
				PrimaryItemVariantId:   headerRes.PrimaryItemVariantID,
				SecondaryItemVariantId: headerRes.SecondaryItemVariantID,
				PrimaryItemUnitId:      headerRes.PrimaryItemUnitID,
				SecondaryItemUnitId:    headerRes.SecondaryItemUnitID,
				PrimaryItemUnitValue:   strconv.FormatInt(headerRes.PrimaryItemUnitValue, 10),
				SecondaryItemUnitValue: strconv.FormatInt(headerRes.SecondaryItemUnitValue, 10),
				Amount:                 strconv.FormatInt(headerRes.Amount, 10),
				Price:                  strconv.FormatInt(headerRes.Price, 10),
			}
			result.SalesOrderItems = append(result.SalesOrderItems, salesOrderItem)
		}

		err = q.UpdateSalesOrderAddItem(ctx, arg.SalesOrderId)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
