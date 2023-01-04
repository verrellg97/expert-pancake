package db

import (
	"context"

	db "github.com/expert-pancake/service/accounting/db/sqlc"
	uuid "github.com/satori/go.uuid"
)

type CloseJournalBookTrxResult struct {
	Message string
}

func (trx *Trx) CloseJournalBookTrx(ctx context.Context, arg string) (CloseJournalBookTrxResult, error) {
	var result CloseJournalBookTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		err = q.CloseJournalBook(ctx, arg)
		if err != nil {
			return err
		}

		journalBookRes, err := q.GetJournalBook(ctx, arg)
		if err != nil {
			return err
		}

		id := uuid.NewV4().String()

		transRes, err := q.InsertJournalBook(ctx, db.InsertJournalBookParams{
			ID:          id,
			CompanyID:   journalBookRes.CompanyID,
			Name:        "NEW JOURNAL BOOK",
			StartPeriod: journalBookRes.EndPeriod.AddDate(0, 0, 1),
			EndPeriod:   journalBookRes.EndPeriod.AddDate(1, 0, 0),
		})
		if err != nil {
			return err
		}

		journalBookAccountsRes, err := q.GetCompanyChartOfAccountBalance(ctx, db.GetCompanyChartOfAccountBalanceParams{
			CompanyID:   transRes.CompanyID,
			StartPeriod: journalBookRes.StartPeriod,
			EndPeriod:   journalBookRes.EndPeriod,
		})
		if err != nil {
			return err
		}

		for _, d := range journalBookAccountsRes {
			var debit_amount, credit_amount int64
			if d.Balance > 0 {
				debit_amount = d.Balance
			} else {
				credit_amount = d.Balance * -1
			}
			err = q.InsertJournalBookAccount(ctx, db.InsertJournalBookAccountParams{
				JournalBookID:    id,
				ChartOfAccountID: d.ChartOfAccountID,
				DebitAmount:      debit_amount,
				CreditAmount:     credit_amount,
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
