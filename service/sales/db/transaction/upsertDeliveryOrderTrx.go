package db

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
	uuid "github.com/satori/go.uuid"
)

type UpsertDeliveryOrderTrxParams struct {
	Id                 string
	CompanyId          string
	BranchId           string
	WarehouseId        string
	TransactionDate    time.Time
	ContactBookId      string
	SecondaryCompanyId string
	KonekinId          string
	SalesOrderId       string
	Items              []model.DeliveryOrderItemsRequest
}

type UpsertDeliveryOrderTrxResult struct {
	Message string
}

func (trx *Trx) UpsertDeliveryOrderTrx(ctx context.Context, arg UpsertDeliveryOrderTrxParams) (UpsertDeliveryOrderTrxResult, error) {
	var result UpsertDeliveryOrderTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		var id = ""
		if arg.Id == "" {
			id = uuid.NewV4().String()
		} else {
			id = arg.Id
		}

		_, err = q.UpsertDeliveryOrder(ctx, db.UpsertDeliveryOrderParams{
			ID:                 id,
			CompanyID:          arg.CompanyId,
			BranchID:           arg.BranchId,
			WarehouseID:        arg.WarehouseId,
			FormNumber:         "DO-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
			TransactionDate:    arg.TransactionDate,
			ContactBookID:      arg.ContactBookId,
			SecondaryCompanyID: arg.SecondaryCompanyId,
			KonekinID:          arg.KonekinId,
			SalesOrderID:       arg.SalesOrderId,
			TotalItems:         int64(len(arg.Items)),
		})

		if err != nil {
			return err
		}

		err = q.DeleteDeliveryOrderItems(ctx, id)
		if err != nil {
			return err
		}

		for _, d := range arg.Items {
			primaryItemUnitValue, _ := strconv.ParseInt(d.PrimaryItemUnitValue, 10, 64)
			secondaryItemUnitValue, _ := strconv.ParseInt(d.SecondaryItemUnitValue, 10, 64)
			amount, _ := strconv.ParseInt(d.Amount, 10, 64)

			_, err := q.UpsertDeliveryOrderItem(ctx, db.UpsertDeliveryOrderItemParams{
				ID:                     uuid.NewV4().String(),
				PurchaseOrderItemID:    d.PurchaseOrderItemId,
				SalesOrderItemID:       d.SalesOrderItemId,
				DeliveryOrderID:        id,
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

		result.Message = "OK"

		return err
	})

	return result, err
}
