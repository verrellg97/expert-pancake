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

func (trx *Trx) UpsertReceiptOrderTrx(ctx context.Context, arg UpsertReceiptOrderTrxParams) (error) {

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

			err := q.InsertReceiptOrderItem(ctx, db.InsertReceiptOrderItemParams{
				ID:                     uuid.NewV4().String(),
				PurchaseOrderItemID:    d.PurchaseOrderItemId,
				SalesOrderItemID:       d.SalesOrderItemId,
				DeliveryOrderItemID:       d.DeliveryOrderItemId,
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
		}

		return err
	})

	return err
}
