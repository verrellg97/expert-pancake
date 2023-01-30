-- Verify warehouse:20230125_warehouse_schema_and_tables on pg

BEGIN;

DO $$ << if_warehouse_schema_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM information_schema.schemata
    WHERE schema_name = 'warehouse'
) THEN RAISE EXCEPTION 'schema warehouse not found';
END IF;
END if_warehouse_schema_exist_test $$;

DO $$ << if_warehouses_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'warehouse'
        AND tablename = 'warehouses'
) THEN RAISE EXCEPTION 'table warehouse.warehouses not found';
END IF;
END if_warehouses_table_exist_test $$;

DO $$ << if_racks_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'warehouse'
        AND tablename = 'racks'
) THEN RAISE EXCEPTION 'table warehouse.racks not found';
END IF;
END if_racks_table_exist_test $$;

DO $$ << if_warehouse_racks_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'warehouse'
        AND tablename = 'warehouse_racks'
) THEN RAISE EXCEPTION 'table warehouse.warehouse_racks not found';
END IF;
END if_warehouse_racks_table_exist_test $$;

ROLLBACK;
