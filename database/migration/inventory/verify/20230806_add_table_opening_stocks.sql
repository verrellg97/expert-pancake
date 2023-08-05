-- Verify inventory:20230806_add_table_opening_stocks on pg

BEGIN;

DO $$ << if_opening_stocks_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'opening_stocks'
) THEN RAISE EXCEPTION 'table inventory.opening_stocks not found';
END IF;
END if_opening_stocks_table_exist_test $$;

ROLLBACK;
