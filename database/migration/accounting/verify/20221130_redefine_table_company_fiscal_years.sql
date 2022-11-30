-- Verify accounting:20221130_redefine_table_company_fiscal_years on pg

BEGIN;

DO $$ << if_company_fiscal_years_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'accounting'
        AND tablename = 'company_fiscal_years'
) THEN RAISE EXCEPTION 'table accounting.company_fiscal_years not found';
END IF;
END if_company_fiscal_years_table_exist_test $$;

ROLLBACK;
