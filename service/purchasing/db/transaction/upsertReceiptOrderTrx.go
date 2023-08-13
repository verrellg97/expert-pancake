package db

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	"github.com/expert-pancake/service/purchasing/model"
	"github.com/expert-pancake/service/purchasing/util"
	uuid "github.com/satori/go.uuid"
)

type UpsertReceiptOrderTrxParams struct {
	Id                 string
	DeliveryOrderId    string
	WarehouseId        string
	CompanyId          string
	BranchId           string
	TransactionDate    string
	ContactBookId      string
	SecondaryCompanyId string
	KonekinId          string
	TotalItems         string
	ReceiptOrderItems  []model.UpsertReceiptOrderItemsRequest
}

type UpsertReceiptOrderTrxResult struct {
	TransactionId      string
	DeliveryOrderId    string
	WarehouseId        string
	CompanyId          string
	BranchId           string
	FormNumber         string
	TransactionDate    string
	ContactBookId      string
	SecondaryCompanyId string
	KonekinId          string
	TotalItems         string
	Status             string
	ReceiptOrderItems  []model.ReceiptOrderItem
}

func (trx *Trx) UpsertReceiptOrderTrx(ctx context.Context, arg UpsertReceiptOrderTrxParams) (UpsertReceiptOrderTrxResult, error) {
	var result UpsertReceiptOrderTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		var id = ""
		if arg.Id == "" {
			id = uuid.NewV4().String()
		} else {
			id = arg.Id
		}

		totalItmes, _ := strconv.ParseInt(arg.TotalItems, 10, 64)

		headerRes, err := q.UpsertReceiptOrder(ctx, db.UpsertReceiptOrderParams{
			ID:                 id,
			DeliveryOrderID:    arg.DeliveryOrderId,
			WarehouseID:        arg.WarehouseId,
			FormNumber:         "RCO-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
			CompanyID:          arg.CompanyId,
			BranchID:           arg.BranchId,
			TransactionDate:    util.StringToDate(arg.TransactionDate),
			ContactBookID:      arg.ContactBookId,
			SecondaryCompanyID: arg.SecondaryCompanyId,
			KonekinID:          arg.KonekinId,
			TotalItems:         totalItmes,
		})

		if err != nil {
			return err
		}

		q.DeleteReceiptOrderItems(ctx, headerRes.ID)

		for _, d := range arg.ReceiptOrderItems {
			primaryItemUnitValue, _ := strconv.ParseInt(d.PrimaryItemUnitValue, 10, 64)
			secondaryItemUnitValue, _ := strconv.ParseInt(d.SecondaryItemUnitValue, 10, 64)
			amount, _ := strconv.ParseInt(d.Amount, 10, 64)

			detailRes, err := q.InsertReceiptOrderItem(ctx, db.InsertReceiptOrderItemParams{
				ID:                     uuid.NewV4().String(),
				PurchaseOrderItemID:    d.PurchaseOrderItemId,
				SalesOrderItemID:       d.SalesOrderItemId,
				DeliveryOrderItemID:    d.DeliveryOrderItemId,
				ReceiptOrderID:         headerRes.ID,
				PrimaryItemVariantID:   d.PrimaryItemVariantId,
				WarehouseRackID:        d.WarehouseRackId,
				Batch:                  util.NewNullableString(d.Batch),
				ExpiredDate:            util.NewNullableDate(util.StringToDate(d.ExpiredDate)),
				ItemBarcodeID:          d.ItemBarcodeId,
				SecondaryItemVariantID: d.SecondaryItemVariantId,
				PrimaryItemUnitID:      d.PrimaryItemUnitId,
				SecondaryItemUnitID:    d.SecondaryItemUnitId,
				PrimaryItemUnitValue:   primaryItemUnitValue,
				SecondaryItemUnitValue: secondaryItemUnitValue,
				Amount:                 amount,
			})
			if err != nil {
				return err
			}

			var batch, expired_date *string
			if detailRes.Batch.Valid {
				batch = &detailRes.Batch.String
			}
			if detailRes.ExpiredDate.Valid {
				expired_date = new(string)
				*expired_date = detailRes.ExpiredDate.Time.Format(util.DateLayoutYMD)
			}

			var receiptOrderItem = model.ReceiptOrderItem{
				DetailId:               detailRes.ID,
				PurchaseOrderItemId:    detailRes.PurchaseOrderItemID,
				SalesOrderItemId:       detailRes.SalesOrderItemID,
				DeliveryOrderItemId:    detailRes.DeliveryOrderItemID,
				ReceiptOrderId:         headerRes.ID,
				PrimaryItemVariantId:   detailRes.PrimaryItemVariantID,
				WarehouseRackId:        detailRes.WarehouseRackID,
				Batch:                  batch,
				ExpiredDate:            expired_date,
				ItemBarcodeId:          detailRes.ItemBarcodeID,
				SecondaryItemVariantId: detailRes.SecondaryItemVariantID,
				PrimaryItemUnitId:      detailRes.PrimaryItemUnitID,
				SecondaryItemUnitId:    detailRes.SecondaryItemUnitID,
				PrimaryItemUnitValue:   d.PrimaryItemUnitValue,
				SecondaryItemUnitValue: d.SecondaryItemUnitValue,
				Amount:                 d.Amount,
			}
			result.ReceiptOrderItems = append(result.ReceiptOrderItems, receiptOrderItem)
		}

		result.TransactionId = headerRes.ID
		result.DeliveryOrderId = headerRes.DeliveryOrderID
		result.WarehouseId = headerRes.WarehouseID
		result.CompanyId = headerRes.CompanyID
		result.BranchId = headerRes.BranchID
		result.FormNumber = headerRes.FormNumber
		result.TransactionDate = headerRes.TransactionDate.Format(util.DateLayoutYMD)
		result.ContactBookId = headerRes.ContactBookID
		result.SecondaryCompanyId = headerRes.SecondaryCompanyID
		result.KonekinId = headerRes.KonekinID
		result.TotalItems = arg.TotalItems
		result.Status = headerRes.Status

		return err
	})

	return result, err
}
