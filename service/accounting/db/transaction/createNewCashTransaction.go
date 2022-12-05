package db

import (
	"context"
	"strconv"

	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/util"
	uuid "github.com/satori/go.uuid"
)

type CreateNewCashTransactionTrxResult struct {
	CompanyId              string
	BranchId               string
	TransactionId          string
	TransactionDate        string
	Type                   string
	MainChartOfAccountId   string
	ContraChartOfAccountId string
	Amount                 string
	Description            string
}

func (trx *Trx) CreateNewCashTransactionTrx(ctx context.Context, arg db.InsertCashTransactionParams) (CreateNewCashTransactionTrxResult, error) {
	var result CreateNewCashTransactionTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		transRes, err := q.InsertCashTransaction(ctx, db.InsertCashTransactionParams{
			ID:                     id,
			CompanyID:              arg.CompanyID,
			BranchID:               arg.BranchID,
			TransactionDate:        arg.TransactionDate,
			Type:                   arg.Type,
			MainChartOfAccountID:   arg.MainChartOfAccountID,
			ContraChartOfAccountID: arg.ContraChartOfAccountID,
			Amount:                 arg.Amount,
			Description:            arg.Description,
		})
		if err != nil {
			return err
		}

		mainAmount := arg.Amount
		if arg.Type == "OUT" {
			mainAmount = mainAmount * -1
		}
		_, err = q.InsertTransactionJournal(ctx, db.InsertTransactionJournalParams{
			CompanyID:            arg.CompanyID,
			BranchID:             arg.BranchID,
			TransactionID:        id,
			TransactionDate:      arg.TransactionDate,
			TransactionReference: "CASH " + arg.Type,
			ChartOfAccountID:     arg.MainChartOfAccountID,
			Amount:               mainAmount,
			Description:          arg.Description,
		})
		if err != nil {
			return err
		}
		if arg.Type == "OUT" {
			_, err = q.InsertTransactionJournal(ctx, db.InsertTransactionJournalParams{
				CompanyID:            arg.CompanyID,
				BranchID:             arg.BranchID,
				TransactionID:        id,
				TransactionDate:      arg.TransactionDate,
				TransactionReference: "CASH " + arg.Type,
				ChartOfAccountID:     arg.ContraChartOfAccountID,
				Amount:               arg.Amount,
				Description:          arg.Description,
			})
			if err != nil {
				return err
			}
		}

		result.CompanyId = transRes.CompanyID
		result.BranchId = transRes.BranchID
		result.TransactionId = id
		result.TransactionDate = transRes.TransactionDate.Format(util.DateLayoutYMD)
		result.Type = transRes.Type
		result.MainChartOfAccountId = transRes.MainChartOfAccountID
		result.ContraChartOfAccountId = transRes.ContraChartOfAccountID
		result.Amount = strconv.FormatInt(transRes.Amount, 10)
		result.Description = transRes.Description

		return err
	})

	return result, err
}
