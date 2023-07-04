package db

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
	uuid "github.com/satori/go.uuid"
)

type UpsertSalesInvoiceTrxParams struct {
	Id                 string
	SalesOrderId       string
	CompanyId          string
	BranchId           string
	TransactionDate    string
	ContactBookId      string
	SecondaryCompanyId string
	KonekinId          string
	CurrencyCode       string
	SalesInvoiceItems  []model.UpsertSalesInvoiceItemRequest
}

func (trx *Trx) UpsertSalesInvoiceTrx(ctx context.Context, arg UpsertSalesInvoiceTrxParams) error {

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		var id = ""
		if arg.Id == "" {
			id = uuid.NewV4().String()
		} else {
			id = arg.Id
		}

		headerRes, err := q.UpsertSalesInvoice(ctx, db.UpsertSalesInvoiceParams{
			ID:                 id,
			FormNumber:         "SI-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
			SalesOrderID:       arg.SalesOrderId,
			CompanyID:          arg.CompanyId,
			BranchID:           arg.BranchId,
			TransactionDate:    util.StringToDate(arg.TransactionDate),
			ContactBookID:      arg.ContactBookId,
			SecondaryCompanyID: arg.SecondaryCompanyId,
			KonekinID:          arg.KonekinId,
			CurrencyCode:       arg.CurrencyCode,
		})

		if err != nil {
			return err
		}

		q.DeleteSalesInvoiceItems(ctx, headerRes.ID)

		for _, d := range arg.SalesInvoiceItems {
			primaryItemUnitValue, _ := strconv.ParseInt(d.PrimaryItemUnitValue, 10, 64)
			secondaryItemUnitValue, _ := strconv.ParseInt(d.SecondaryItemUnitValue, 10, 64)
			amount, _ := strconv.ParseInt(d.Amount, 10, 64)
			price, _ := strconv.ParseInt(d.Price, 10, 64)

			err := q.InsertSalesInvoiceItem(ctx, db.InsertSalesInvoiceItemParams{
				ID:                     uuid.NewV4().String(),
				PurchaseOrderItemID:    d.PurchaseOrderItemId,
				SalesOrderItemID:       d.SalesOrderItemId,
				SalesInvoiceID:         headerRes.ID,
				PrimaryItemVariantID:   d.PrimaryItemVariantId,
				SecondaryItemVariantID: d.SecondaryItemVariantId,
				PrimaryItemUnitID:      d.PrimaryItemUnitId,
				SecondaryItemUnitID:    d.SecondaryItemUnitId,
				PrimaryItemUnitValue:   primaryItemUnitValue,
				SecondaryItemUnitValue: secondaryItemUnitValue,
				Amount:                 amount,
				Price:                  price,
			})
			if err != nil {
				return err
			}
		}

		q.UpdateSalesInvoiceAddItem(ctx, headerRes.ID)

		return err
	})

	return err
}
