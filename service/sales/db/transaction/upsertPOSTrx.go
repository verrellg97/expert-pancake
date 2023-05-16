package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
	uuid "github.com/satori/go.uuid"
)

type UpsertPOSTrxParams struct {
	POS      model.POSRequest
	POSItems []model.POSItemRequest
}

type UpsertPOSTrxResult struct {
	Message string
}

func (trx *Trx) UpsertPOSTrx(ctx context.Context, arg UpsertPOSTrxParams) (UpsertPOSTrxResult, error) {
	var result UpsertPOSTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		var id = ""
		if arg.POS.Id == "" {
			id = uuid.NewV4().String()
		} else {
			id = arg.POS.Id
		}

		totalItems := strconv.ParseInt(arg.POS.TotalItems, 10, 64)
		total := strconv.ParseInt(arg.POS.Total, 10, 64)


		headerRes, err := q.UpsertPOS(ctx, db.UpsertPOSParams{
			ID:                 id,
			CompanyID:          arg.POS.CompanyId,
			BranchID:           arg.POS.BranchId,
			WarehouseID:        arg.POS.WarehouseId,
			FormNumber:         arg.POS.FormNumber,
			TransactionDate:    util.StringToDate(arg.POS.TransactionDate),
			ContactBookID:      arg.POS.ContactBookId,
			SecondaryCompanyID: arg.POS.SecondaryCompanyId,
			KonekinID:          arg.POS.KonekinId,
			CurrencyCode:       arg.POS.CurrencyCode,
			ChartOfAccountID:   arg.POS.ChartOfAccountId,
			TotalItems:         totalItems,
			Total:              strconv.ParseInt(arg.POS.Total, 10, 64),
		})
		if err != nil {
			return err
		}

		err = q.DeletePOSItemsPOS(ctx, arg.POS.Id)
		if err != nil {
			return err
		}

		for _, d := range arg.POSItems {

			detailRes, err := q.InsertPOSItem(ctx, db.InsertPOSItemParams{
				ID:              uuid.NewV4().String(),
				PointOfSaleID:   headerRes.ID,
				WarehouseRackID: d.WarehouseRackId,
				ItemVariantID:   d.ItemVariantId,
				ItemUnitID:      d.ItemUnitId,
				ItemUnitValue:   strconv.ParseInt(d.ItemUnitValue, 10, 64),
				Batch:           d.Batch,
				ExpiredDate:     util.StringToDate(d.ExpiredDate),
				ItemBarcodeID:   d.ItemBarcodeId,
				Amount:          strconv.ParseInt(d.Amount, 10, 64),
				Price:           strconv.ParseInt(d.Price, 10, 64),
			})
			if err != nil {
				return err
			}

		}

		return err
	})

	return result, err
}
