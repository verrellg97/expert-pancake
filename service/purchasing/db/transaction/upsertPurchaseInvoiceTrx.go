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

type UpsertPurchaseInvoiceTrxParams struct {
	Id                   string
	SalesInvoiceId       string
	ReceiptOrderId       string
	CompanyId            string
	BranchId             string
	TransactionDate      string
	ContactBookId        string
	SecondaryCompanyId   string
	KonekinId            string
	CurrencyCode         string
	TotalItems           string
	Total                string
	PurchaseInvoiceItems []model.UpsertPurchaseInvoiceItemRequest
}

func (trx *Trx) UpsertPurchaseInvoiceTrx(ctx context.Context, arg UpsertPurchaseInvoiceTrxParams) error {

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		var id = ""
		if arg.Id == "" {
			id = uuid.NewV4().String()
		} else {
			id = arg.Id
		}

		totalItmes, _ := strconv.ParseInt(arg.TotalItems, 10, 64)
		total, _ := strconv.ParseInt(arg.Total, 10, 64)

		headerRes, err := q.UpsertPurchaseInvoice(ctx, db.UpsertPurchaseInvoiceParams{
			ID:                 id,
			FormNumber:         "PI-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
			SalesInvoiceID:     arg.SalesInvoiceId,
			ReceiptOrderID:     arg.ReceiptOrderId,
			CompanyID:          arg.CompanyId,
			BranchID:           arg.BranchId,
			TransactionDate:    util.StringToDate(arg.TransactionDate),
			ContactBookID:      arg.ContactBookId,
			SecondaryCompanyID: arg.SecondaryCompanyId,
			KonekinID:          arg.KonekinId,
			CurrencyCode:       arg.CurrencyCode,
			TotalItems:         totalItmes,
			Total:              total,
		})

		if err != nil {
			return err
		}

		q.DeletePurchaseInvoiceItems(ctx, headerRes.ID)

		for _, d := range arg.PurchaseInvoiceItems {
			primaryItemUnitValue, _ := strconv.ParseInt(d.PrimaryItemUnitValue, 10, 64)
			secondaryItemUnitValue, _ := strconv.ParseInt(d.SecondaryItemUnitValue, 10, 64)
			amount, _ := strconv.ParseInt(d.Amount, 10, 64)
			price, _ := strconv.ParseInt(d.Price, 10, 64)

			err := q.InsertPurchaseInvoiceItem(ctx, db.InsertPurchaseInvoiceItemParams{
				ID:                     uuid.NewV4().String(),
				PurchaseOrderItemID:    d.PurchaseOrderItemId,
				SalesOrderItemID:       d.SalesOrderItemId,
				SalesInvoiceItemID:     d.SalesInvoiceItemId,
				ReceiptOrderItemID:     d.ReceiptOrderItemId,
				PurchaseInvoiceID:      headerRes.ID,
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

		return err
	})

	return err
}
