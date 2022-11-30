// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	GetCompanyChartOfAccounts(ctx context.Context, arg GetCompanyChartOfAccountsParams) ([]GetCompanyChartOfAccountsRow, error)
	InsertCompanyChartOfAccount(ctx context.Context, arg InsertCompanyChartOfAccountParams) (AccountingCompanyChartOfAccount, error)
	UpdateCompanyChartOfAccount(ctx context.Context, arg UpdateCompanyChartOfAccountParams) (AccountingCompanyChartOfAccount, error)
	UpsertCompanyFiscalYear(ctx context.Context, arg UpsertCompanyFiscalYearParams) (AccountingCompanyFiscalYear, error)
}

var _ Querier = (*Queries)(nil)