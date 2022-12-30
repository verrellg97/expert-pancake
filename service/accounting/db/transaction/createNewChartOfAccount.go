package db

import (
	"context"

	db "github.com/expert-pancake/service/accounting/db/sqlc"
	uuid "github.com/satori/go.uuid"
)

type CreateNewChartOfAccountTrxParams struct {
	CompanyId             string
	CurrencyCode          string
	ChartOfAccountGroupId string
	AccountCode           string
	AccountName           string
	BankName              string
	BankAccountNumber     string
	BankCode              string
	IsAllBranches         bool
	Branches              []string
}

type CreateNewChartOfAccountTrxResult struct {
	ChartOfAccountId      string
	CompanyId             string
	CurrencyCode          string
	ChartOfAccountGroupId string
	AccountCode           string
	AccountName           string
	BankName              string
	BankAccountNumber     string
	BankCode              string
	IsAllBranches         bool
	Branches              []string
	IsDeleted             bool
}

func (trx *Trx) CreateNewChartOfAccountTrx(ctx context.Context, arg CreateNewChartOfAccountTrxParams) (CreateNewChartOfAccountTrxResult, error) {
	var result CreateNewChartOfAccountTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		transRes, err := q.InsertCompanyChartOfAccount(ctx, db.InsertCompanyChartOfAccountParams{
			ID:                    id,
			CompanyID:             arg.CompanyId,
			CurrencyCode:          arg.CurrencyCode,
			ChartOfAccountGroupID: arg.ChartOfAccountGroupId,
			AccountCode:           arg.AccountCode,
			AccountName:           arg.AccountName,
			BankName:              arg.BankName,
			BankAccountNumber:     arg.BankAccountNumber,
			BankCode:              arg.BankCode,
			IsAllBranches:         arg.IsAllBranches,
		})
		if err != nil {
			return err
		}

		if !arg.IsAllBranches {
			for _, d := range arg.Branches {
				err = q.InsertChartOfAccountBranches(ctx, db.InsertChartOfAccountBranchesParams{
					ChartOfAccountID: id,
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
		result.ChartOfAccountGroupId = transRes.ChartOfAccountGroupID
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
