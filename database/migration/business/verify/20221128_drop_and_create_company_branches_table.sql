-- Verify business:20221128_drop_and_create_company_branches_table on pg

BEGIN;

DO $$ << if_company_branches_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business'
        AND tablename = 'company_branches'
) THEN RAISE EXCEPTION 'table business.company_branches not found';
END IF;
END if_company_branches_table_exist_test $$;

ROLLBACK;
