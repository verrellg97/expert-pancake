package db

import (
	"context"
	"strconv"
	"fmt"
	"math/rand"

	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
	uuid "github.com/satori/go.uuid"
)

type UpsertPOSTrxParams struct {
	Id                 string
	CompanyId          string
	BranchId           string
	WarehouseId        string
	FormNumber         string
	TransactionDate    string
	ContactBookId      string
	SecondaryCompanyId string
	KonekinId          string
	CurrencyCode       string
	ChartOfAccountId   string
	TotalItems         string
	Total              string
	POSItems           []model.POSItemRequest
}

type UpsertPOSTrxResult struct {
	Message string
}

func (trx *Trx) UpsertPOSTrx(ctx context.Context, arg UpsertPOSTrxParams) (UpsertPOSTrxResult, error) {
	var result UpsertPOSTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		var id = ""
		if arg.Id == "" {
			id = uuid.NewV4().String()
		} else {
			id = arg.Id
		}

		var formNumber = ""
		if arg.FormNumber == "" {
			formNumber = "POS-" + fmt.Sprintf("%08d", rand.Intn(100000000))
		} else {
			formNumber = arg.FormNumber
		}

		totalItems, _ := strconv.ParseInt(arg.TotalItems, 10, 64)
		total, _ := strconv.ParseInt(arg.Total, 10, 64)

		headerRes, err := q.UpsertPOS(ctx, db.UpsertPOSParams{
			ID:                 id,
			CompanyID:          arg.CompanyId,
			BranchID:           arg.BranchId,
			WarehouseID:        arg.WarehouseId,
			FormNumber:         formNumber,
			TransactionDate:    util.StringToDate(arg.TransactionDate),
			ContactBookID:      arg.ContactBookId,
			SecondaryCompanyID: arg.SecondaryCompanyId,
			KonekinID:          arg.KonekinId,
			CurrencyCode:       arg.CurrencyCode,
			ChartOfAccountID:   arg.ChartOfAccountId,
			TotalItems:         totalItems,
			Total:              total,
		})
		if err != nil {
			return err
		}

		err = q.DeletePOSItemsPOS(ctx, arg.Id)
		if err != nil {
			return err
		}

		for _, d := range arg.POSItems {
			itemUnitValue, _ := strconv.ParseInt(d.ItemUnitValue, 10, 64)
			amount, _ := strconv.ParseInt(d.Amount, 10, 64)
			price, _ := strconv.ParseInt(d.Price, 10, 64)

			_, err := q.InsertPOSItem(ctx, db.InsertPOSItemParams{
				ID:              uuid.NewV4().String(),
				PointOfSaleID:   headerRes.ID,
				WarehouseRackID: d.WarehouseRackId,
				ItemVariantID:   d.ItemVariantId,
				ItemUnitID:      d.ItemUnitId,
				ItemUnitValue:   itemUnitValue,
				Batch:           util.NewNullableString(d.Batch),
				ExpiredDate:     util.NewNullableDate(util.StringToDate(d.ExpiredDate)),
				ItemBarcodeID:   d.ItemBarcodeId,
				Amount:          amount,
				Price:           price,
			})
			if err != nil {
				return err
			}

		}

		result.Message = "OK"

		return err
	})

	return result, err
}
