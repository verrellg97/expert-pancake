-- Verify account:20221102_add_new_table_user_addresses on pg

BEGIN;

DO $$ << if_user_addresses_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'account'
        AND tablename = 'user_addresses'
) THEN RAISE EXCEPTION 'table account.user_addresses not found';
END IF;
END if_user_addresses_table_exist_test $$;

ROLLBACK;
