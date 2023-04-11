package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	uuid "github.com/satori/go.uuid"
)

type UpsertPurchaseOrderItemTrxParams struct {
	Id                     string
	PurchaseOrderId        string
	PrimaryItemVariantId   string
	SecondaryItemVariantId string
	PrimaryItemUnitId      string
	SecondaryItemUnitId    string
	PrimaryItemUnitValue   string
	SecondaryItemUnitValue string
	Amount                 string
	Price                  string
}

type UpsertPurchaseOrderItemTrxResult struct {
	DetailId               string
	PurchaseOrderId        string
	PrimaryItemVariantId   string
	SecondaryItemVariantId string
	PrimaryItemUnitId      string
	SecondaryItemUnitId    string
	PrimaryItemUnitValue   string
	SecondaryItemUnitValue string
	Amount                 string
	Price                  string
}

func (trx *Trx) UpsertPurchaseOrderItemTrx(ctx context.Context, arg UpsertPurchaseOrderItemTrxParams) (UpsertPurchaseOrderItemTrxResult, error) {
	var result UpsertPurchaseOrderItemTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		var id = ""
		if arg.Id == "" {
			id = uuid.NewV4().String()
		} else {
			id = arg.Id
		}

		primaryItemUnitValue, _ := strconv.ParseInt(arg.PrimaryItemUnitValue, 10, 64)
		secondaryItemUnitValue, _ := strconv.ParseInt(arg.SecondaryItemUnitValue, 10, 64)
		amount, _ := strconv.ParseInt(arg.Amount, 10, 64)
		price, _ := strconv.ParseInt(arg.Price, 10, 64)

		headerRes, err := q.UpsertPurchaseOrderItem(ctx, db.UpsertPurchaseOrderItemParams{
			ID:                     id,
			PurchaseOrderID:        arg.PurchaseOrderId,
			PrimaryItemVariantID:   arg.PrimaryItemVariantId,
			SecondaryItemVariantID: arg.SecondaryItemVariantId,
			PrimaryItemUnitID:      arg.PrimaryItemUnitId,
			SecondaryItemUnitID:    arg.SecondaryItemUnitId,
			PrimaryItemUnitValue:   primaryItemUnitValue,
			SecondaryItemUnitValue: secondaryItemUnitValue,
			Amount:                 amount,
			Price:                  price,
		})
		if err != nil {
			return err
		}

		err = q.UpdatePurchaseOrderAddItem(ctx, arg.PurchaseOrderId)
		if err != nil {
			return err
		}

		result.DetailId = headerRes.ID
		result.PurchaseOrderId = headerRes.PurchaseOrderID
		result.PrimaryItemVariantId = headerRes.PrimaryItemVariantID
		result.SecondaryItemVariantId = headerRes.SecondaryItemVariantID
		result.PrimaryItemUnitId = headerRes.PrimaryItemUnitID
		result.SecondaryItemUnitId = headerRes.SecondaryItemUnitID
		result.PrimaryItemUnitValue = strconv.FormatInt(headerRes.PrimaryItemUnitValue, 10)
		result.SecondaryItemUnitValue = strconv.FormatInt(headerRes.SecondaryItemUnitValue, 10)
		result.Amount = strconv.FormatInt(headerRes.Amount, 10)
		result.Price = strconv.FormatInt(headerRes.Price, 10)

		return err
	})

	return result, err
}
