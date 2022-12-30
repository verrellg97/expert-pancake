package db

import (
	"context"
	"strconv"
	"time"

	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/expert-pancake/service/accounting/util"
	uuid "github.com/satori/go.uuid"
)

type CreateNewJournalBookTrxParams struct {
	CompanyId       string
	Name            string
	StartPeriod     time.Time
	EndPeriod       time.Time
	ChartOfAccounts []model.AddJournalBookAccountRequest
}

type CreateNewJournalBookTrxResult struct {
	JournalBookId string
	CompanyId     string
	Name          string
	StartPeriod   string
	EndPeriod     string
	IsClosed      bool
}

func (trx *Trx) CreateNewJournalBookTrx(ctx context.Context, arg CreateNewJournalBookTrxParams) (CreateNewJournalBookTrxResult, error) {
	var result CreateNewJournalBookTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		transRes, err := q.InsertJournalBook(ctx, db.InsertJournalBookParams{
			ID:          id,
			CompanyID:   arg.CompanyId,
			Name:        arg.Name,
			StartPeriod: arg.StartPeriod,
			EndPeriod:   arg.EndPeriod,
		})
		if err != nil {
			return err
		}

		for _, d := range arg.ChartOfAccounts {
			amount, _ := strconv.ParseInt(d.Amount, 10, 64)
			err = q.InsertJournalBookAccount(ctx, db.InsertJournalBookAccountParams{
				JournalBookID:    id,
				ChartOfAccountID: d.ChartOfAccountId,
				Amount:           amount,
			})
			if err != nil {
				return err
			}
			_, err = q.InsertTransactionJournal(ctx, db.InsertTransactionJournalParams{
				CompanyID:            arg.CompanyId,
				TransactionID:        id,
				TransactionDate:      arg.StartPeriod,
				TransactionReference: "OPENING BALANCE",
				ChartOfAccountID:     d.ChartOfAccountId,
				Amount:               amount,
			})
			if err != nil {
				return err
			}
		}

		result.JournalBookId = transRes.ID
		result.CompanyId = transRes.CompanyID
		result.Name = transRes.Name
		result.StartPeriod = transRes.StartPeriod.Format(util.DateLayoutYMD)
		result.EndPeriod = transRes.EndPeriod.Format(util.DateLayoutYMD)
		result.IsClosed = transRes.IsClosed

		return err
	})

	return result, err
}
