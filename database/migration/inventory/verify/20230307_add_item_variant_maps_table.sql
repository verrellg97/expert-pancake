-- Verify inventory:20230307_add_item_variant_maps_table on pg

BEGIN;

DO $$ << if_item_variant_maps_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'item_variant_maps'
) THEN RAISE EXCEPTION 'table inventory.item_variant_maps not found';
END IF;
END if_item_variant_maps_table_exist_test $$;

ROLLBACK;
