-- Verify accounting:20221219_add_table_chart_of_account_branches_and_alter_table_company_chart_of_accounts on pg

BEGIN;

DO $$ << if_chart_of_account_branches_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'accounting'
        AND tablename = 'chart_of_account_branches'
) THEN RAISE EXCEPTION 'table accounting.chart_of_account_branches not found';
END IF;
END if_chart_of_account_branches_table_exist_test $$;

ROLLBACK;
