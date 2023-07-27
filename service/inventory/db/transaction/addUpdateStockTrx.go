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

type AddUpdateStockTrxParams struct {
	TransactionDate time.Time
	TransactionCode string
	WarehouseId     string
	WarehouseRackId string
	VariantId       string
	ItemUnitId      string
	ItemUnitValue   string
	BeginningStock  string
	EndingStock     string
	Batch           string
	ExpiredDate     string
}

type AddUpdateStockTrxResult struct {
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
	BeginningStock  string
	EndingStock     string
	Batch           *string
	ExpiredDate     *string
}

func (trx *Trx) AddUpdateStockTrx(ctx context.Context, arg AddUpdateStockTrxParams) (AddUpdateStockTrxResult, error) {
	var result AddUpdateStockTrxResult

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

		item_unit_value, _ := strconv.ParseInt(arg.ItemUnitValue, 10, 64)
		beginning_stock, _ := strconv.ParseInt(arg.BeginningStock, 10, 64)
		ending_stock, _ := strconv.ParseInt(arg.EndingStock, 10, 64)
		err = q.InsertUpdateStock(ctx, db.InsertUpdateStockParams{
			ID:              id,
			FormNumber:      "UPS-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
			TransactionDate: arg.TransactionDate,
			WarehouseID:     arg.WarehouseId,
			WarehouseRackID: arg.WarehouseRackId,
			VariantID:       arg.VariantId,
			ItemUnitID:      arg.ItemUnitId,
			ItemUnitValue:   item_unit_value,
			BeginningStock:  beginning_stock,
			EndingStock:     ending_stock,
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
			TransactionCode:      arg.TransactionCode,
			TransactionReference: "UPDATE STOCK",
			WarehouseID:          arg.WarehouseId,
			WarehouseRackID:      arg.WarehouseRackId,
			VariantID:            arg.VariantId,
			ItemBarcodeID:        barcodeRes,
			Amount:               (ending_stock - beginning_stock) * item_unit_value,
		})
		if err != nil {
			return err
		}

		headerRes, err := q.GetUpdateStock(ctx, id)

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
		result.BeginningStock = strconv.FormatInt(headerRes.BeginningStock, 10)
		result.EndingStock = strconv.FormatInt(headerRes.EndingStock, 10)
		result.Batch = batch
		result.ExpiredDate = expired_date

		return err
	})

	return result, err
}
