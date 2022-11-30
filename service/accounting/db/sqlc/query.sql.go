// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
	"time"
)

const getCompanyChartOfAccounts = `-- name: GetCompanyChartOfAccounts :many
SELECT id, company_id, branch_id, account_code, account_name, account_group, bank_name, 
bank_account_number, bank_code, opening_balance, is_deleted
FROM accounting.company_chart_of_accounts
WHERE company_id = $1 AND account_name LIKE $2
`

type GetCompanyChartOfAccountsParams struct {
	CompanyID   string `db:"company_id"`
	AccountName string `db:"account_name"`
}

type GetCompanyChartOfAccountsRow struct {
	ID                string `db:"id"`
	CompanyID         string `db:"company_id"`
	BranchID          string `db:"branch_id"`
	AccountCode       string `db:"account_code"`
	AccountName       string `db:"account_name"`
	AccountGroup      string `db:"account_group"`
	BankName          string `db:"bank_name"`
	BankAccountNumber string `db:"bank_account_number"`
	BankCode          string `db:"bank_code"`
	OpeningBalance    int64  `db:"opening_balance"`
	IsDeleted         bool   `db:"is_deleted"`
}

func (q *Queries) GetCompanyChartOfAccounts(ctx context.Context, arg GetCompanyChartOfAccountsParams) ([]GetCompanyChartOfAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, getCompanyChartOfAccounts, arg.CompanyID, arg.AccountName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCompanyChartOfAccountsRow
	for rows.Next() {
		var i GetCompanyChartOfAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.BranchID,
			&i.AccountCode,
			&i.AccountName,
			&i.AccountGroup,
			&i.BankName,
			&i.BankAccountNumber,
			&i.BankCode,
			&i.OpeningBalance,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertCompanyChartOfAccount = `-- name: InsertCompanyChartOfAccount :one
INSERT INTO accounting.company_chart_of_accounts(id, company_id, branch_id, 
account_code, account_name, account_group, 
bank_name, bank_account_number, bank_code, opening_balance, is_deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id, company_id, branch_id, account_code, account_name, account_group, bank_name, bank_account_number, bank_code, opening_balance, is_deleted, created_at, updated_at
`

type InsertCompanyChartOfAccountParams struct {
	ID                string `db:"id"`
	CompanyID         string `db:"company_id"`
	BranchID          string `db:"branch_id"`
	AccountCode       string `db:"account_code"`
	AccountName       string `db:"account_name"`
	AccountGroup      string `db:"account_group"`
	BankName          string `db:"bank_name"`
	BankAccountNumber string `db:"bank_account_number"`
	BankCode          string `db:"bank_code"`
	OpeningBalance    int64  `db:"opening_balance"`
	IsDeleted         bool   `db:"is_deleted"`
}

func (q *Queries) InsertCompanyChartOfAccount(ctx context.Context, arg InsertCompanyChartOfAccountParams) (AccountingCompanyChartOfAccount, error) {
	row := q.db.QueryRowContext(ctx, insertCompanyChartOfAccount,
		arg.ID,
		arg.CompanyID,
		arg.BranchID,
		arg.AccountCode,
		arg.AccountName,
		arg.AccountGroup,
		arg.BankName,
		arg.BankAccountNumber,
		arg.BankCode,
		arg.OpeningBalance,
		arg.IsDeleted,
	)
	var i AccountingCompanyChartOfAccount
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.BranchID,
		&i.AccountCode,
		&i.AccountName,
		&i.AccountGroup,
		&i.BankName,
		&i.BankAccountNumber,
		&i.BankCode,
		&i.OpeningBalance,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCompanyChartOfAccount = `-- name: UpdateCompanyChartOfAccount :one
UPDATE accounting.company_chart_of_accounts
SET account_code = $2,
account_name = $3,
account_group = $4,
bank_name = $5,
bank_account_number = $6,
bank_code = $7,
opening_balance = $8,
is_deleted = $9,
updated_at = NOW()
WHERE id = $1
RETURNING id, company_id, branch_id, account_code, account_name, account_group, bank_name, bank_account_number, bank_code, opening_balance, is_deleted, created_at, updated_at
`

type UpdateCompanyChartOfAccountParams struct {
	ID                string `db:"id"`
	AccountCode       string `db:"account_code"`
	AccountName       string `db:"account_name"`
	AccountGroup      string `db:"account_group"`
	BankName          string `db:"bank_name"`
	BankAccountNumber string `db:"bank_account_number"`
	BankCode          string `db:"bank_code"`
	OpeningBalance    int64  `db:"opening_balance"`
	IsDeleted         bool   `db:"is_deleted"`
}

func (q *Queries) UpdateCompanyChartOfAccount(ctx context.Context, arg UpdateCompanyChartOfAccountParams) (AccountingCompanyChartOfAccount, error) {
	row := q.db.QueryRowContext(ctx, updateCompanyChartOfAccount,
		arg.ID,
		arg.AccountCode,
		arg.AccountName,
		arg.AccountGroup,
		arg.BankName,
		arg.BankAccountNumber,
		arg.BankCode,
		arg.OpeningBalance,
		arg.IsDeleted,
	)
	var i AccountingCompanyChartOfAccount
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.BranchID,
		&i.AccountCode,
		&i.AccountName,
		&i.AccountGroup,
		&i.BankName,
		&i.BankAccountNumber,
		&i.BankCode,
		&i.OpeningBalance,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertCompanyFiscalYear = `-- name: UpsertCompanyFiscalYear :one
INSERT INTO accounting.company_fiscal_years(company_id, start_period, end_period)
VALUES ($1, $2, $3)
ON CONFLICT (company_id)
DO UPDATE SET
    start_period = EXCLUDED.start_period,
    end_period = EXCLUDED.end_period,
    updated_at = NOW()
RETURNING company_id, start_period, end_period, created_at, updated_at
`

type UpsertCompanyFiscalYearParams struct {
	CompanyID   string    `db:"company_id"`
	StartPeriod time.Time `db:"start_period"`
	EndPeriod   time.Time `db:"end_period"`
}

func (q *Queries) UpsertCompanyFiscalYear(ctx context.Context, arg UpsertCompanyFiscalYearParams) (AccountingCompanyFiscalYear, error) {
	row := q.db.QueryRowContext(ctx, upsertCompanyFiscalYear, arg.CompanyID, arg.StartPeriod, arg.EndPeriod)
	var i AccountingCompanyFiscalYear
	err := row.Scan(
		&i.CompanyID,
		&i.StartPeriod,
		&i.EndPeriod,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}