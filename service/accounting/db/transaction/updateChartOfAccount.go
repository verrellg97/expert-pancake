package db

import (
	"context"

	db "github.com/expert-pancake/service/accounting/db/sqlc"
)

type UpdateChartOfAccountTrxParams struct {
	ID                string
	CompanyID         string
	CurrencyCode      string
	ReportType        string
	AccountType       string
	AccountGroup      string
	AccountCode       string
	AccountName       string
	BankName          string
	BankAccountNumber string
	BankCode          string
	IsAllBranches     bool
	Branches          []string
	IsDeleted         bool
}

func (trx *Trx) UpdateChartOfAccountTrx(ctx context.Context, arg UpdateChartOfAccountTrxParams) (CreateNewChartOfAccountTrxResult, error) {
	var result CreateNewChartOfAccountTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		transRes, err := q.UpdateCompanyChartOfAccount(ctx, db.UpdateCompanyChartOfAccountParams{
			ID:                arg.ID,
			CurrencyCode:      arg.CurrencyCode,
			ReportType:        arg.ReportType,
			AccountType:       arg.AccountType,
			AccountGroup:      arg.AccountGroup,
			AccountCode:       arg.AccountCode,
			AccountName:       arg.AccountName,
			BankName:          arg.BankName,
			BankAccountNumber: arg.BankAccountNumber,
			BankCode:          arg.BankCode,
			IsAllBranches:     arg.IsAllBranches,
			IsDeleted:         arg.IsDeleted,
		})
		if err != nil {
			return err
		}

		err = q.DeleteChartOfAccountBranches(ctx, arg.ID)
		if err != nil {
			return err
		}

		if !arg.IsAllBranches {
			for _, d := range arg.Branches {
				err = q.InsertChartOfAccountBranches(ctx, db.InsertChartOfAccountBranchesParams{
					ChartOfAccountID: arg.ID,
					BranchID:         d,
				})
				if err != nil {
					return err
				}
			}
		}

		result.ChartOfAccountId = transRes.ID
		result.CompanyId = transRes.CompanyID
		result.CurrencyCode = transRes.CurrencyCode
		result.ReportType = transRes.ReportType
		result.AccountType = transRes.AccountType
		result.AccountGroup = transRes.AccountGroup
		result.AccountCode = transRes.AccountCode
		result.AccountName = transRes.AccountName
		result.BankName = transRes.BankName
		result.BankAccountNumber = transRes.BankAccountNumber
		result.BankCode = transRes.BankCode
		result.IsAllBranches = transRes.IsAllBranches
		result.Branches = arg.Branches
		result.IsDeleted = transRes.IsDeleted

		return err
	})

	return result, err
}
