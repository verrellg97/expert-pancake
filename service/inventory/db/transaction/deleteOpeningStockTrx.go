package db

import (
	"context"

	db "github.com/expert-pancake/service/inventory/db/sqlc"
)

type DeleteOpeningStockTrxParams struct {
	Id              string
}

func (trx *Trx) DeleteOpeningStockTrx(ctx context.Context, arg DeleteOpeningStockTrxParams) (error) {

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		opening_stock, err := q.GetOpeningStock(ctx, arg.Id)
		
		err = q.DeleteOpeningStock(ctx, opening_stock.ID)
		if err != nil {
			return err
		}

		deleteStockMovement := q.DeleteStockMovement(ctx, db.DeleteStockMovementParams{
			TransactionID: opening_stock.FormNumber,
			TransactionReference: "OPENING STOCK",
		})
		if deleteStockMovement != nil {
			return deleteStockMovement
		}


		return err
	})

	return err
}
