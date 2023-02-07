-- Verify inventory:20230207_add_item_reorders_table on pg

BEGIN;

DO $$ << if_item_reorders_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'item_reorders'
) THEN RAISE EXCEPTION 'table inventory.item_reorders not found';
END IF;
END if_item_reorders_table_exist_test $$;

ROLLBACK;
