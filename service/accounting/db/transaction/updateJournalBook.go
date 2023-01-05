package db

import (
	"context"
	"strconv"
	"time"

	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/expert-pancake/service/accounting/util"
)

type UpdateJournalBookTrxParams struct {
	Id              string
	Name            string
	StartPeriod     time.Time
	EndPeriod       time.Time
	ChartOfAccounts []model.AddJournalBookAccountRequest
}

type UpdateJournalBookTrxResult struct {
	JournalBookId string
	CompanyId     string
	Name          string
	StartPeriod   string
	EndPeriod     string
}

func (trx *Trx) UpdateJournalBookTrx(ctx context.Context, arg UpdateJournalBookTrxParams) (UpdateJournalBookTrxResult, error) {
	var result UpdateJournalBookTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		transRes, err := q.UpdateJournalBook(ctx, db.UpdateJournalBookParams{
			ID:          arg.Id,
			Name:        arg.Name,
			StartPeriod: arg.StartPeriod,
			EndPeriod:   arg.EndPeriod,
		})
		if err != nil {
			return err
		}

		err = q.DeleteJournalBookAccount(ctx, arg.Id)
		if err != nil {
			return err
		}

		err = q.DeleteTransactionJournalByIdRef(ctx, db.DeleteTransactionJournalByIdRefParams{
			TransactionID:        arg.Id,
			TransactionReference: "OPENING BALANCE",
		})
		if err != nil {
			return err
		}

		for _, d := range arg.ChartOfAccounts {
			debit_amount, _ := strconv.ParseInt(d.DebitAmount, 10, 64)
			credit_amount, _ := strconv.ParseInt(d.CreditAmount, 10, 64)
			err = q.InsertJournalBookAccount(ctx, db.InsertJournalBookAccountParams{
				JournalBookID:    arg.Id,
				ChartOfAccountID: d.ChartOfAccountId,
				DebitAmount:      debit_amount,
				CreditAmount:     credit_amount,
				Description:      d.Description,
			})
			if err != nil {
				return err
			}
			amount := debit_amount
			if debit_amount == 0 {
				amount = -1 * credit_amount
			}
			_, err = q.InsertTransactionJournal(ctx, db.InsertTransactionJournalParams{
				CompanyID:            transRes.CompanyID,
				TransactionID:        arg.Id,
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

		return err
	})

	return result, err
}
