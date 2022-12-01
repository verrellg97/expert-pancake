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
SET 
    account_code = $2,
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
WHERE company_id = $1 AND account_name LIKE $2
AND account_group LIKE $3 
AND CASE WHEN $4 = '0' THEN is_deleted = FALSE 
WHEN $4 = '1' THEN is_deleted = TRUE ELSE 1=1 END;

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

-- name: InsertCashTransaction :one
INSERT INTO accounting.cash_transactions(id, company_id, branch_id, transaction_date, 
transaction_type, type, main_chart_of_account_id, contra_chart_of_account_id, 
amount, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: InsertTransactionJournal :one
INSERT INTO accounting.transactions_journal(company_id, branch_id, transaction_id, 
transaction_date, transaction_reference , transaction_type, chart_of_account_id, 
amount, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetCashTransactions :many
SELECT a.id, a.company_id, a.branch_id, a.transaction_date, 
a.transaction_type, a.type, a.main_chart_of_account_id, a.contra_chart_of_account_id, 
a.amount, a.description, b.account_name AS main_chart_of_account_name, 
c.account_name AS contra_chart_of_account_name
FROM accounting.cash_transactions a
JOIN accounting.company_chart_of_accounts b ON a.main_chart_of_account_id = b.id
LEFT JOIN accounting.company_chart_of_accounts c ON a.contra_chart_of_account_id = c.id
WHERE a.company_id = $1
AND a.branch_id = $2 AND a.type LIKE $3;

-- name: GetCompanyChartOfAccount :one
SELECT *
FROM accounting.company_chart_of_accounts
WHERE id = $1;

-- name: GetCashTransactionsGroupByDate :many
SELECT transaction_date, 
SUM(CASE WHEN type = 'IN' THEN amount ELSE 0 END) AS cash_in, 
SUM(CASE WHEN type = 'OUT' THEN amount ELSE 0 END) AS cash_out
FROM accounting.cash_transactions 
WHERE company_id = $1
AND branch_id = $2
GROUP BY transaction_date;