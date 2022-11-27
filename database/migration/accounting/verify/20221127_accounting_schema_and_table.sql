-- Verify accounting:20221127_accounting_schema_and_table on pg

BEGIN;

DO $$ << if_accounting_schema_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM information_schema.schemata
    WHERE schema_name = 'accounting'
) THEN RAISE EXCEPTION 'schema accounting not found';
END IF;
END if_accounting_schema_exist_test $$;

DO $$ << if_company_fiscal_years_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'accounting'
        AND tablename = 'company_fiscal_years'
) THEN RAISE EXCEPTION 'table accounting.company_fiscal_years not found';
END IF;
END if_company_fiscal_years_table_exist_test $$;

DO $$ << if_company_chart_of_accounts_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'accounting'
        AND tablename = 'company_chart_of_accounts'
) THEN RAISE EXCEPTION 'table accounting.company_chart_of_accounts not found';
END IF;
END if_company_chart_of_accounts_table_exist_test $$;

ROLLBACK;
