-- name: UpsertCompanyFiscalYear :one
INSERT INTO accounting.company_fiscal_years(company_id, start_period, end_period)
VALUES ($1, $2, $3)
ON CONFLICT (company_id)
DO UPDATE SET
    start_period = EXCLUDED.start_period,
    end_period = EXCLUDED.end_period,
    updated_at = NOW()
RETURNING *;

-- name: InsertCompanyChartOfAccount :one
INSERT INTO accounting.company_chart_of_accounts(id, company_id, branch_id, 
account_code, account_name, account_group, 
bank_name, bank_account_number, bank_code, opening_balance, is_deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: UpdateCompanyChartOfAccount :one
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
RETURNING *;

-- name: GetCompanyChartOfAccounts :many
SELECT id, company_id, branch_id, account_code, account_name, account_group, bank_name, 
bank_account_number, bank_code, opening_balance, is_deleted
FROM accounting.company_chart_of_accounts
WHERE company_id = $1 AND account_name LIKE $2;

-- name: GetCompanySettingFiscalYear :one
SELECT company_id, start_period, end_period
FROM accounting.company_fiscal_years
WHERE company_id = $1;

-- name: GetCompanySettingBank :one
SELECT id, company_id, branch_id, account_code, account_name, account_group, bank_name, 
bank_account_number, bank_code, opening_balance, is_deleted
FROM accounting.company_chart_of_accounts
WHERE company_id = $1 
AND account_group = 'BANK'
ORDER BY created_at LIMIT 1;

-- name: GetCompanySettingCash :one
SELECT id, company_id, branch_id, account_code, account_name, account_group, bank_name, 
bank_account_number, bank_code, opening_balance, is_deleted
FROM accounting.company_chart_of_accounts
WHERE company_id = $1 
AND account_group = 'KAS'
ORDER BY created_at LIMIT 1;