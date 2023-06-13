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

type UpsertDeliveryOrderTrxParams struct {
	Id                 string
	CompanyId          string
	BranchId           string
	TransactionDate    time.Time
	ContactBookId      string
	SecondaryCompanyId string
	KonekinId          string
	SecondaryBranchId  string
}

type UpsertDeliveryOrderTrxResult struct {
	TransactionId      string
	CompanyId          string
	BranchId           string
	FormNumber         string
	TransactionDate    string
	ContactBookId      string
	SecondaryCompanyId string
	KonekinId          string
	SecondaryBranchId  string
	TotalItems         string
	Status             string
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

		headerRes, err := q.UpsertDeliveryOrder(ctx, db.UpsertDeliveryOrderParams{
			ID:                 id,
			CompanyID:          arg.CompanyId,
			BranchID:           arg.BranchId,
			FormNumber:         "DO-" + fmt.Sprintf("%08d", rand.Intn(100000000)),
			TransactionDate:    arg.TransactionDate,
			ContactBookID:      arg.ContactBookId,
			SecondaryCompanyID: arg.SecondaryCompanyId,
			KonekinID:          arg.KonekinId,
			SecondaryBranchID:  arg.SecondaryBranchId,
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
		result.SecondaryCompanyId = headerRes.SecondaryCompanyID
		result.KonekinId = headerRes.KonekinID
		result.SecondaryBranchId = headerRes.SecondaryBranchID
		result.TotalItems = strconv.FormatInt(headerRes.TotalItems, 10)
		result.Status = headerRes.Status

		return err
	})

	return result, err
}
