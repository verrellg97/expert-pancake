package db

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/util"
	uuid "github.com/satori/go.uuid"
)

type AddOpeningStockTrxParams struct {
	TransactionDate time.Time
	CompanyId       string
	BranchId        string
	WarehouseId     string
	WarehouseRackId string
	VariantId       string
	ItemUnitId      string
	ItemUnitValue   string
	Amount          string
	Price           string
	Batch           string
	ExpiredDate     string
}

type AddOpeningStockTrxResult struct {
	TransactionId   string
	FormNumber      string
	TransactionDate string
	WarehouseId     string
	WarehouseRackId string
	ItemId          string
	ItemName        string
	VariantId       string
	VariantName     string
	ItemUnitId      string
	ItemUnitName    string
	ItemUnitValue   string
	Amount          string
	Price           string
	Batch           *string
	ExpiredDate     *string
}

func (trx *Trx) AddOpeningStockTrx(ctx context.Context, arg AddOpeningStockTrxParams) (AddOpeningStockTrxResult, error) {
	var result AddOpeningStockTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		barcodeRes, err := q.GetItemBarcode(ctx, db.GetItemBarcodeParams{
			VariantID:         arg.VariantId,
			IsNullBatch:       !util.NewNullableString(arg.Batch).Valid,
			Batch:             util.NewNullableString(arg.Batch),
			IsNullExpiredDate: !util.NewNullableDate(util.StringToDate(arg.ExpiredDate)).Valid,
			ExpiredDate:       util.NewNullableDate(util.StringToDate(arg.ExpiredDate)),
		})

		if err != nil {
			if err == sql.ErrNoRows {
				barcodeRes = uuid.NewV4().String()
				err = q.InsertItemBarcode(ctx, db.InsertItemBarcodeParams{
					ID:          barcodeRes,
					VariantID:   arg.VariantId,
					Batch:       util.NewNullableString(arg.Batch),
					ExpiredDate: util.NewNullableDate(util.StringToDate(arg.ExpiredDate)),
				})
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}

		id := uuid.NewV4().String()
		formNumber := "OPS-" + fmt.Sprintf("%08d", rand.Intn(100000000))

		item_unit_value, _ := strconv.ParseInt(arg.ItemUnitValue, 10, 64)
		amount, _ := strconv.ParseInt(arg.Amount, 10, 64)
		price, _ := strconv.ParseInt(arg.Price, 10, 64)
		err = q.InsertOpeningStock(ctx, db.InsertOpeningStockParams{
			ID:              id,
			FormNumber:      formNumber,
			TransactionDate: arg.TransactionDate,
			CompanyID:       arg.CompanyId,
			BranchID:        arg.BranchId,
			WarehouseID:     arg.WarehouseId,
			WarehouseRackID: arg.WarehouseRackId,
			VariantID:       arg.VariantId,
			ItemUnitID:      arg.ItemUnitId,
			ItemUnitValue:   item_unit_value,
			Amount:          amount,
			Price:           price,
			Batch:           util.NewNullableString(arg.Batch),
			ExpiredDate:     util.NewNullableDate(util.StringToDate(arg.ExpiredDate)),
			ItemBarcodeID:   barcodeRes,
		})
		if err != nil {
			return err
		}

		err = q.InsertStockMovement(ctx, db.InsertStockMovementParams{
			ID:                   uuid.NewV4().String(),
			TransactionID:        id,
			TransactionDate:      arg.TransactionDate,
			TransactionCode:      formNumber,
			TransactionReference: "OPENING STOCK",
			CompanyID:            arg.CompanyId,
			BranchID:             arg.BranchId,
			WarehouseID:          arg.WarehouseId,
			WarehouseRackID:      arg.WarehouseRackId,
			VariantID:            arg.VariantId,
			ItemBarcodeID:        barcodeRes,
			Amount:               amount * item_unit_value,
		})
		if err != nil {
			return err
		}

		headerRes, err := q.GetOpeningStock(ctx, id)

		var batch, expired_date *string
		if headerRes.Batch.Valid {
			batch = &headerRes.Batch.String
		}
		if headerRes.ExpiredDate.Valid {
			expired_date = new(string)
			*expired_date = headerRes.ExpiredDate.Time.Format(util.DateLayoutYMD)
		}

		result.TransactionId = headerRes.ID
		result.FormNumber = headerRes.FormNumber
		result.TransactionDate = headerRes.TransactionDate.Format(util.DateLayoutYMD)
		result.WarehouseId = headerRes.WarehouseID
		result.WarehouseRackId = headerRes.WarehouseRackID
		result.ItemId = headerRes.ItemID
		result.ItemName = headerRes.ItemName
		result.VariantId = headerRes.VariantID
		result.VariantName = headerRes.VariantName
		result.ItemUnitId = headerRes.ItemUnitID
		result.ItemUnitName = headerRes.ItemUnitName
		result.ItemUnitValue = strconv.FormatInt(headerRes.ItemUnitValue, 10)
		result.Amount = strconv.FormatInt(headerRes.Amount, 10)
		result.Price = strconv.FormatInt(headerRes.Price, 10)
		result.Batch = batch
		result.ExpiredDate = expired_date

		return err
	})

	return result, err
}
