-- Verify warehouse:20230201_drop_old_table_definition_and_create_new_table_definition on pg

BEGIN;

DO $$ << if_warehouses_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'warehouse'
        AND tablename = 'warehouses'
) THEN RAISE EXCEPTION 'table warehouse.warehouses not found';
END IF;
END if_warehouses_table_exist_test $$;

DO $$ << if_warehouse_racks_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'warehouse'
        AND tablename = 'warehouse_racks'
) THEN RAISE EXCEPTION 'table warehouse.warehouse_racks not found';
END IF;
END if_warehouse_racks_table_exist_test $$;

ROLLBACK;
