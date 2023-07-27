package db

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	db "github.com/expert-pancake/service/inventory/db/sqlc"
	"github.com/expert-pancake/service/inventory/model"
	"github.com/expert-pancake/service/inventory/util"
	uuid "github.com/satori/go.uuid"
)

type AddInternalStockTransferTrxParams struct {
	SourceWarehouseId      string
	DestinationWarehouseId string
	TransactionCode        string
	TransactionDate        time.Time
	Items                  []model.InternalStockTransferItemRequest
}

type AddInternalStockTransferTrxResult struct {
	TransactionId          string
	SourceWarehouseId      string
	DestinationWarehouseId string
	FormNumber             string
	TransactionDate        string
}

func (trx *Trx) AddInternalStockTransferTrx(ctx context.Context, arg AddInternalStockTransferTrxParams) (AddInternalStockTransferTrxResult, error) {
	var result AddInternalStockTransferTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		headerRes, err := q.InsertInternalStockTransfer(ctx, db.InsertInternalStockTransferParams{
			ID:                     id,
			SourceWarehouseID:      arg.SourceWarehouseId,
			DestinationWarehouseID: arg.DestinationWarehouseId,
			FormNumber:             "IST-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
			TransactionDate:        arg.TransactionDate,
		})
		if err != nil {
			return err
		}

		for _, d := range arg.Items {

			barcodeRes, err := q.GetItemBarcode(ctx, db.GetItemBarcodeParams{
				VariantID:         d.VariantId,
				IsNullBatch:       !util.NewNullableString(d.Batch).Valid,
				Batch:             util.NewNullableString(d.Batch),
				IsNullExpiredDate: !util.NewNullableDate(util.StringToDate(d.ExpiredDate)).Valid,
				ExpiredDate:       util.NewNullableDate(util.StringToDate(d.ExpiredDate)),
			})

			if err != nil {
				if err == sql.ErrNoRows {
					barcodeRes = uuid.NewV4().String()
					err = q.InsertItemBarcode(ctx, db.InsertItemBarcodeParams{
						ID:          barcodeRes,
						VariantID:   d.VariantId,
						Batch:       util.NewNullableString(d.Batch),
						ExpiredDate: util.NewNullableDate(util.StringToDate(d.ExpiredDate)),
					})
					if err != nil {
						return err
					}
				} else {
					return err
				}

			}

			item_unit_value, _ := strconv.ParseInt(d.ItemUnitValue, 10, 64)
			amount, _ := strconv.ParseInt(d.Amount, 10, 64)
			detailId := uuid.NewV4().String()
			err = q.InsertInternalStockTransferItem(ctx, db.InsertInternalStockTransferItemParams{
				ID:                      detailId,
				InternalStockTransferID: id,
				WarehouseRackID:         d.WarehouseRackId,
				VariantID:               d.VariantId,
				ItemUnitID:              d.ItemUnitId,
				ItemUnitValue:           item_unit_value,
				Amount:                  amount,
				Batch:                   util.NewNullableString(d.Batch),
				ExpiredDate:             util.NewNullableDate(util.StringToDate(d.ExpiredDate)),
				ItemBarcodeID:           barcodeRes,
			})
			if err != nil {
				return err
			}

			err = q.InsertStockMovement(ctx, db.InsertStockMovementParams{
				ID:                   uuid.NewV4().String(),
				TransactionID:        id,
				TransactionDate:      arg.TransactionDate,
				TransactionCode:      arg.TransactionCode,
				TransactionReference: "INTERNAL TRANSFER",
				DetailTransactionID:  detailId,
				WarehouseID:          arg.SourceWarehouseId,
				WarehouseRackID:      d.WarehouseRackId,
				VariantID:            d.VariantId,
				ItemBarcodeID:        barcodeRes,
				Amount:               -amount * item_unit_value,
			})
			if err != nil {
				return err
			}
		}

		result.TransactionId = headerRes.ID
		result.SourceWarehouseId = arg.SourceWarehouseId
		result.DestinationWarehouseId = arg.DestinationWarehouseId
		result.FormNumber = headerRes.FormNumber
		result.TransactionDate = headerRes.TransactionDate.Format(util.DateLayoutYMD)

		return err
	})

	return result, err
}
