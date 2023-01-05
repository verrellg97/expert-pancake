package db

import (
	"context"
	"strconv"
	"time"

	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/expert-pancake/service/accounting/util"
)

type UpdateMemorialJournalTrxParams struct {
	Id              string
	TransactionDate time.Time
	Description     string
	ChartOfAccounts []model.AddMemorialJournalAccountRequest
}

type UpdateMemorialJournalTrxResult struct {
	MemorialJournalId string
	CompanyId         string
	TransactionDate   string
	Description       string
}

func (trx *Trx) UpdateMemorialJournalTrx(ctx context.Context, arg UpdateMemorialJournalTrxParams) (UpdateMemorialJournalTrxResult, error) {
	var result UpdateMemorialJournalTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		transRes, err := q.UpdateMemorialJournal(ctx, db.UpdateMemorialJournalParams{
			ID:              arg.Id,
			TransactionDate: arg.TransactionDate,
			Description:     arg.Description,
		})
		if err != nil {
			return err
		}

		err = q.DeleteMemorialJournalAccount(ctx, arg.Id)
		if err != nil {
			return err
		}

		err = q.DeleteTransactionJournalByIdRef(ctx, db.DeleteTransactionJournalByIdRefParams{
			TransactionID:        arg.Id,
			TransactionReference: "MEMORIAL JOURNAL",
		})
		if err != nil {
			return err
		}

		for _, d := range arg.ChartOfAccounts {
			debit_amount, _ := strconv.ParseInt(d.DebitAmount, 10, 64)
			credit_amount, _ := strconv.ParseInt(d.CreditAmount, 10, 64)
			err = q.InsertMemorialJournalAccount(ctx, db.InsertMemorialJournalAccountParams{
				MemorialJournalID: arg.Id,
				ChartOfAccountID:  d.ChartOfAccountId,
				DebitAmount:       debit_amount,
				CreditAmount:      credit_amount,
				Description:       d.Description,
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
				TransactionDate:      arg.TransactionDate,
				TransactionReference: "MEMORIAL JOURNAL",
				ChartOfAccountID:     d.ChartOfAccountId,
				Amount:               amount,
			})
			if err != nil {
				return err
			}
		}

		result.MemorialJournalId = transRes.ID
		result.CompanyId = transRes.CompanyID
		result.TransactionDate = transRes.TransactionDate.Format(util.DateLayoutYMD)
		result.Description = transRes.Description

		return err
	})

	return result, err
}
