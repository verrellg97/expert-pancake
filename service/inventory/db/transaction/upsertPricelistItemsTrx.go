package db

import (
	"context"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
)

type UpsertPricelistItemsTrxParams struct {
	PricelistId    string
	PricelistItems []model.PricelistItemRequest
}

type UpsertPricelistItemsTrxResult struct {
	Message string
}

func (trx *Trx) UpsertPricelistItemsTrx(ctx context.Context, arg UpsertPricelistItemsTrxParams) (UpsertPricelistItemsTrxResult, error) {
	var result UpsertPricelistItemsTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		for _, d := range arg.PricelistItems {

			price, _ := strconv.ParseInt(d.Price, 10, 64)
			arg := db.UpsertPricelistItemParams{
				PricelistID: arg.PricelistId,
				VariantID:   d.VariantId,
				Price:       price,
			}

			err = q.UpsertPricelistItem(ctx, arg)
			if err != nil {
				return errors.NewServerError(model.UpsertPricelistItemsError, err.Error())
			}
		}

		result.Message = "OK"

		return err
	})

	return result, err
}
