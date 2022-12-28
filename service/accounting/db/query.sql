-- name: UpsertCompanyFiscalYear :one
INSERT INTO accounting.company_fiscal_years(company_id, start_period, end_period)
VALUES ($1, $2, $3)
ON CONFLICT (company_id)
DO UPDATE SET
    start_period = EXCLUDED.start_period,
    end_period = EXCLUDED.end_period,
    updated_at = NOW()
RETURNING *;

-- name: InsertChartOfAccountGroup :one
INSERT INTO accounting.chart_of_account_groups(id, company_id, 
report_type, account_type, account_group_name)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateChartOfAccountGroup :one
UPDATE accounting.chart_of_account_groups
SET 
    report_type = $2,
    account_type = $3,
    account_group_name = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetChartOfAccountGroups :many
SELECT id, company_id, report_type, account_type, account_group_name 
FROM accounting.chart_of_account_groups
WHERE company_id = $1;

-- name: GetChartOfAccountGroup :one
SELECT id, company_id, report_type, account_type, account_group_name 
FROM accounting.chart_of_account_groups
WHERE id = $1;

-- name: GetChartOfAccountGroupByAccTypeAccGroup :one
SELECT id, company_id, report_type, account_type, account_group_name 
FROM accounting.chart_of_account_groups
WHERE company_id = $1 AND account_type = $2 AND account_group_name = $3;

-- name: InsertCompanyChartOfAccount :one
INSERT INTO accounting.company_chart_of_accounts(id, company_id, currency_code, 
chart_of_account_group_id, account_code, account_name, 
bank_name, bank_account_number, bank_code, is_all_branches)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
    chart_of_account_group_id = $3,
    account_code = $4,
    account_name = $5,
    bank_name = $6,
    bank_account_number = $7,
    bank_code = $8,
    is_all_branches = $9,
    is_deleted = $10,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetCompanyChartOfAccounts :many
SELECT a.id, a.company_id, a.currency_code, a.chart_of_account_group_id,
b.report_type, b.account_type, b.account_group_name, a.account_code, a.account_name,
a.bank_name, a.bank_account_number, a.bank_code, a.is_all_branches, a.is_deleted
FROM accounting.company_chart_of_accounts a
JOIN accounting.chart_of_account_groups b ON a.chart_of_account_group_id = b.id
WHERE a.company_id = $1
AND a.account_name LIKE $2
AND CASE WHEN @is_filter_journal_type::bool
THEN b.account_type = ANY(@account_types::text[]) ELSE TRUE END
AND CASE WHEN @is_deleted_filter::bool
THEN a.is_deleted = $3 ELSE TRUE END;

-- name: GetCompanySettingFiscalYear :one
SELECT company_id, start_period, end_period
FROM accounting.company_fiscal_years
WHERE company_id = $1;

-- name: GetCompanySettingChartOfAccount :one
SELECT a.id, a.company_id, a.currency_code, a.chart_of_account_group_id,
b.report_type, b.account_type, b.account_group_name, a.account_code, a.account_name,
a.bank_name, a.bank_account_number, a.bank_code, a.is_all_branches, a.is_deleted
FROM accounting.company_chart_of_accounts a
JOIN accounting.chart_of_account_groups b ON a.chart_of_account_group_id = b.id
WHERE a.company_id = $1
AND b.account_group_name = $2
ORDER BY a.created_at LIMIT 1;

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