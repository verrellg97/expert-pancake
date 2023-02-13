-- Verify inventory:20230210_add_unit_categories_and_item_infos_table on pg

BEGIN;

DO $$ << if_unit_categories_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'unit_categories'
) THEN RAISE EXCEPTION 'table inventory.unit_categories not found';
END IF;
END if_unit_categories_table_exist_test $$;

DO $$ << if_item_infos_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'item_infos'
) THEN RAISE EXCEPTION 'table inventory.item_infos not found';
END IF;
END if_item_infos_table_exist_test $$;

ROLLBACK;
