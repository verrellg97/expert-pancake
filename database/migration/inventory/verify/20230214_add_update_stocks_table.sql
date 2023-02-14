-- Verify inventory:20230214_add_update_stocks_table on pg

BEGIN;

DO $$ << if_update_stocks_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'update_stocks'
) THEN RAISE EXCEPTION 'table inventory.update_stocks not found';
END IF;
END if_update_stocks_table_exist_test $$;

ROLLBACK;
