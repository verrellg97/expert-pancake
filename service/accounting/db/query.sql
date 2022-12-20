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
INSERT INTO accounting.company_chart_of_accounts(id, company_id, currency_code, 
report_type, account_type, account_group, account_code, account_name, 
bank_name, bank_account_number, bank_code, is_all_branches)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: InsertChartOfAccountBranches :exec
INSERT INTO accounting.chart_of_account_branches(chart_of_account_id, branch_id)
VALUES ($1, $2);

-- name: DeleteChartOfAccountBranches :exec
DELETE FROM accounting.chart_of_account_branches
WHERE chart_of_account_id = $1;

-- name: GetChartOfAccountBranches :many
SELECT chart_of_account_id, branch_id FROM accounting.chart_of_account_branches
WHERE chart_of_account_id = $1;

-- name: UpdateCompanyChartOfAccount :one
UPDATE accounting.company_chart_of_accounts
SET 
    currency_code = $2,
    report_type = $3,
    account_type = $4,
    account_group = $5,
    account_code = $6,
    account_name = $7,
    bank_name = $8,
    bank_account_number = $9,
    bank_code = $10,
    is_all_branches = $11,
    is_deleted = $12,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetCompanyChartOfAccounts :many
SELECT id, company_id, currency_code, 
report_type, account_type, account_group, account_code, account_name,
bank_name, bank_account_number, bank_code, is_all_branches, is_deleted
FROM accounting.company_chart_of_accounts
WHERE company_id = $1
AND account_name LIKE $2
AND CASE WHEN @is_filter_journal_type::bool
THEN account_type = ANY(@account_types::text[]) ELSE TRUE END
AND CASE WHEN @is_deleted_filter::bool
THEN is_deleted = $3 ELSE TRUE END;

-- name: GetCompanySettingFiscalYear :one
SELECT company_id, start_period, end_period
FROM accounting.company_fiscal_years
WHERE company_id = $1;

-- name: GetCompanySettingBank :one
SELECT id, company_id, account_type, account_group, account_code, account_name,
bank_name, bank_account_number, bank_code, is_deleted
FROM accounting.company_chart_of_accounts
WHERE company_id = $1 
AND account_group = 'BANK'
ORDER BY created_at LIMIT 1;

-- name: GetCompanySettingCash :one
SELECT id, company_id, account_type, account_group, account_code, account_name,
bank_name, bank_account_number, bank_code, is_deleted
FROM accounting.company_chart_of_accounts
WHERE company_id = $1 
AND account_group = 'KAS'
ORDER BY created_at LIMIT 1;

-- name: InsertCashTransaction :one
INSERT INTO accounting.cash_transactions(id, company_id, branch_id, transaction_date, 
type, main_chart_of_account_id, contra_chart_of_account_id, 
amount, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: InsertTransactionJournal :one
INSERT INTO accounting.transactions_journal(company_id, branch_id, transaction_id, 
transaction_date, transaction_reference , chart_of_account_id, 
amount, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetCashTransactions :many
SELECT a.id, a.company_id, a.branch_id, a.transaction_date, 
a.type, a.main_chart_of_account_id, a.contra_chart_of_account_id, 
a.amount, a.description, b.account_name AS main_chart_of_account_name, 
c.account_name AS contra_chart_of_account_name
FROM accounting.cash_transactions a
JOIN accounting.company_chart_of_accounts b ON a.main_chart_of_account_id = b.id
LEFT JOIN accounting.company_chart_of_accounts c ON a.contra_chart_of_account_id = c.id
WHERE a.company_id = $1
AND a.branch_id = $2 
AND a.type LIKE $3;

-- name: GetCompanyChartOfAccount :one
SELECT *
FROM accounting.company_chart_of_accounts
WHERE company_id = $1;

-- name: GetCashTransactionsGroupByDate :many
SELECT transaction_date, 
SUM(CASE WHEN type = 'IN' THEN amount ELSE 0 END) AS cash_in, 
SUM(CASE WHEN type = 'OUT' THEN amount ELSE 0 END) AS cash_out
FROM accounting.cash_transactions 
WHERE company_id = $1
AND branch_id = $2
GROUP BY transaction_date;