-- Verify inventory:20230130_alter_item_components_table_name_and_add_items_item_variants_item_unit_table on pg

BEGIN;

DO $$ << if_items_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'items'
) THEN RAISE EXCEPTION 'table inventory.items not found';
END IF;
END if_items_table_exist_test $$;

DO $$ << if_item_variants_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'item_variants'
) THEN RAISE EXCEPTION 'table inventory.item_variants not found';
END IF;
END if_item_variants_table_exist_test $$;

DO $$ << if_item_units_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'item_units'
) THEN RAISE EXCEPTION 'table inventory.item_units not found';
END IF;
END if_item_units_table_exist_test $$;

ROLLBACK;
