-- Verify account:20221029_drop_old_table_definition_and_create_new_table_definition on pg

BEGIN;

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

DO $$ << if_user_passwords_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'account'
        AND tablename = 'user_passwords'
) THEN RAISE EXCEPTION 'table account.user_passwords not found';
END IF;
END if_user_passwords_table_exist_test $$;

ROLLBACK;
