package db

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/util"
	uuid "github.com/satori/go.uuid"
)

type UpsertSalesOrderTrxParams struct {
	Id                    string
	PurchaseOrderId       string
	PurchaseOrderBranchId string
	CompanyId             string
	BranchId              string
	TransactionDate       time.Time
	ContactBookId         string
	SecondaryCompanyId    string
	KonekinId             string
	CurrencyCode          string
}

type UpsertSalesOrderTrxResult struct {
	TransactionId         string
	PurchaseOrderId       string
	PurchaseOrderBranchId string
	CompanyId             string
	BranchId              string
	FormNumber            string
	TransactionDate       string
	ContactBookId         string
	SecondaryCompanyId    string
	KonekinId             string
	CurrencyCode          string
	TotalItems            string
	Total                 string
	Status                string
}

func (trx *Trx) UpsertSalesOrderTrx(ctx context.Context, arg UpsertSalesOrderTrxParams) (UpsertSalesOrderTrxResult, error) {
	var result UpsertSalesOrderTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		var id = ""
		if arg.Id == "" {
			id = uuid.NewV4().String()
		} else {
			id = arg.Id
		}

		headerRes, err := q.UpsertSalesOrder(ctx, db.UpsertSalesOrderParams{
			ID:                    id,
			PurchaseOrderID:       arg.PurchaseOrderId,
			PurchaseOrderBranchID: arg.PurchaseOrderBranchId,
			CompanyID:             arg.CompanyId,
			BranchID:              arg.BranchId,
			FormNumber:            "SO-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
			TransactionDate:       arg.TransactionDate,
			ContactBookID:         arg.ContactBookId,
			SecondaryCompanyID:    arg.SecondaryCompanyId,
			KonekinID:             arg.KonekinId,
			CurrencyCode:          arg.CurrencyCode,
		})

		if err != nil {
			return err
		}

		result.TransactionId = headerRes.ID
		result.PurchaseOrderId = headerRes.PurchaseOrderID
		result.PurchaseOrderBranchId = headerRes.PurchaseOrderBranchID
		result.CompanyId = headerRes.CompanyID
		result.BranchId = headerRes.BranchID
		result.FormNumber = headerRes.FormNumber
		result.TransactionDate = headerRes.TransactionDate.Format(util.DateLayoutYMD)
		result.ContactBookId = headerRes.ContactBookID
		result.SecondaryCompanyId = headerRes.SecondaryCompanyID
		result.KonekinId = headerRes.KonekinID
		result.CurrencyCode = headerRes.CurrencyCode
		result.TotalItems = strconv.FormatInt(headerRes.TotalItems, 10)
		result.Total = strconv.FormatInt(headerRes.Total, 10)
		result.Status = headerRes.Status

		return err
	})

	return result, err
}
