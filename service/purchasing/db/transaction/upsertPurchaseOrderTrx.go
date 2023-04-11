package db

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	db "github.com/expert-pancake/service/purchasing/db/sqlc"
	"github.com/expert-pancake/service/purchasing/util"
	uuid "github.com/satori/go.uuid"
)

type UpsertPurchaseOrderTrxParams struct {
	Id              string
	CompanyId       string
	BranchId        string
	TransactionDate time.Time
	ContactBookId   string
	KonekinId       string
	CurrencyCode    string
}

type UpsertPurchaseOrderTrxResult struct {
	TransactionId   string
	CompanyId       string
	BranchId        string
	FormNumber      string
	TransactionDate string
	ContactBookId   string
	KonekinId       string
	CurrencyCode    string
	TotalItems      string
	Total           string
	Status          string
}

func (trx *Trx) UpsertPurchaseOrderTrx(ctx context.Context, arg UpsertPurchaseOrderTrxParams) (UpsertPurchaseOrderTrxResult, error) {
	var result UpsertPurchaseOrderTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		var id = ""
		if arg.Id == "" {
			id = uuid.NewV4().String()
		} else {
			id = arg.Id
		}

		headerRes, err := q.UpsertPurchaseOrder(ctx, db.UpsertPurchaseOrderParams{
			ID:              id,
			CompanyID:       arg.CompanyId,
			BranchID:        arg.BranchId,
			FormNumber:      "PO-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
			TransactionDate: arg.TransactionDate,
			ContactBookID:   arg.ContactBookId,
			KonekinID:       arg.KonekinId,
			CurrencyCode:    arg.CurrencyCode,
		})

		if err != nil {
			return err
		}

		result.TransactionId = headerRes.ID
		result.CompanyId = headerRes.CompanyID
		result.BranchId = headerRes.BranchID
		result.FormNumber = headerRes.FormNumber
		result.TransactionDate = headerRes.TransactionDate.Format(util.DateLayoutYMD)
		result.ContactBookId = headerRes.ContactBookID
		result.KonekinId = headerRes.KonekinID
		result.CurrencyCode = headerRes.CurrencyCode
		result.TotalItems = strconv.FormatInt(headerRes.TotalItems, 10)
		result.Total = strconv.FormatInt(headerRes.Total, 10)
		result.Status = headerRes.Status

		return err
	})

	return result, err
}
