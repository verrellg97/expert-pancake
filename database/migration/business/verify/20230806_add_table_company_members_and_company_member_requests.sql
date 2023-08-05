-- Verify business:20230806_add_table_company_members_and_company_member_requests on pg

BEGIN;

DO $$ << if_company_members_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business'
        AND tablename = 'company_members'
) THEN RAISE EXCEPTION 'table business.company_members not found';
END IF;
END if_company_members_table_exist_test $$;

DO $$ << if_company_member_requests_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business'
        AND tablename = 'company_member_requests'
) THEN RAISE EXCEPTION 'table business.company_member_requests not found';
END IF;
END if_company_member_requests_table_exist_test $$;

ROLLBACK;
