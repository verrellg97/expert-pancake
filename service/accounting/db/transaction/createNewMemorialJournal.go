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

type CreateNewMemorialJournalTrxParams struct {
	CompanyId       string
	TransactionDate time.Time
	Description     string
	ChartOfAccounts []model.AddMemorialJournalAccountRequest
}

type CreateNewMemorialJournalTrxResult struct {
	MemorialJournalId string
	CompanyId         string
	TransactionDate   string
	Description       string
}

func (trx *Trx) CreateNewMemorialJournalTrx(ctx context.Context, arg CreateNewMemorialJournalTrxParams) (CreateNewMemorialJournalTrxResult, error) {
	var result CreateNewMemorialJournalTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		transRes, err := q.InsertMemorialJournal(ctx, db.InsertMemorialJournalParams{
			ID:              id,
			CompanyID:       arg.CompanyId,
			TransactionDate: arg.TransactionDate,
			Description:     arg.Description,
		})
		if err != nil {
			return err
		}

		for _, d := range arg.ChartOfAccounts {
			debit_amount, _ := strconv.ParseInt(d.DebitAmount, 10, 64)
			credit_amount, _ := strconv.ParseInt(d.CreditAmount, 10, 64)
			err = q.InsertMemorialJournalAccount(ctx, db.InsertMemorialJournalAccountParams{
				MemorialJournalID: id,
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
				CompanyID:            arg.CompanyId,
				TransactionID:        id,
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
