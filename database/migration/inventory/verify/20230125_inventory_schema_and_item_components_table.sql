-- Verify inventory:20230125_inventory_schema_and_item_components_table on pg

BEGIN;

DO $$ << if_inventory_schema_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM information_schema.schemata
    WHERE schema_name = 'inventory'
) THEN RAISE EXCEPTION 'schema inventory not found';
END IF;
END if_inventory_schema_exist_test $$;

DO $$ << if_item_brands_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'item_brands'
) THEN RAISE EXCEPTION 'table inventory.item_brands not found';
END IF;
END if_item_brands_table_exist_test $$;

DO $$ << if_item_groups_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'item_groups'
) THEN RAISE EXCEPTION 'table inventory.item_groups not found';
END IF;
END if_item_groups_table_exist_test $$;

DO $$ << if_item_units_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'item_units'
) THEN RAISE EXCEPTION 'table inventory.item_units not found';
END IF;
END if_item_units_table_exist_test $$;

ROLLBACK;
