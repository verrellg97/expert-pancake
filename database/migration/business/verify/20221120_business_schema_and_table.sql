-- Verify business:20221120_business_schema_and_table on pg

BEGIN;

DO $$ << if_business_schema_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM information_schema.schemata
    WHERE schema_name = 'business'
) THEN RAISE EXCEPTION 'schema business not found';
END IF;
END if_business_schema_exist_test $$;

DO $$ << if_companies_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business'
        AND tablename = 'companies'
) THEN RAISE EXCEPTION 'table business.companies not found';
END IF;
END if_companies_table_exist_test $$;

DO $$ << if_company_branches_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business'
        AND tablename = 'company_branches'
) THEN RAISE EXCEPTION 'table business.company_branches not found';
END IF;
END if_company_branches_table_exist_test $$;

ROLLBACK;
