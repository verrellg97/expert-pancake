-- name: UpsertCompanyFiscalYear :one
INSERT INTO accounting.company_fiscal_years(company_id, start_month, start_year, end_month, end_year)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (company_id)
DO UPDATE SET
    start_month = EXCLUDED.start_month,
    start_year = EXCLUDED.start_year,
    end_month = EXCLUDED.end_month,
    end_year = EXCLUDED.end_year,
    updated_at = NOW()
RETURNING *;

-- name: GetCompanyChartOfAccounts :many
SELECT id, company_id, branch_id, account_code, account_name, account_group, bank_name, 
bank_account_number, bank_code, opening_balance, is_deleted
FROM accounting.company_chart_of_accounts
WHERE company_id = $1 AND account_name LIKE $2;