-- Verify account:20221011_account_schema_and_table on pg

BEGIN;

DO $$ << if_account_schema_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM information_schema.schemata
    WHERE schema_name = 'account'
) THEN RAISE EXCEPTION 'schema account not found';
END IF;
END if_account_schema_exist_test $$;

DO $$ << if_users_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'account'
        AND tablename = 'users'
) THEN RAISE EXCEPTION 'table account.users not found';
END IF;
END if_users_table_exist_test $$;

DO $$ << if_user_infos_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'account'
        AND tablename = 'user_infos'
) THEN RAISE EXCEPTION 'table account.user_infos not found';
END IF;
END if_user_infos_table_exist_test $$;

ROLLBACK;
