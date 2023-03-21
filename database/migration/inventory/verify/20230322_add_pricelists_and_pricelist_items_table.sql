-- Verify inventory:20230322_add_pricelists_and_pricelist_items_table on pg

BEGIN;

DO $$ << if_pricelists_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'pricelists'
) THEN RAISE EXCEPTION 'table inventory.pricelists not found';
END IF;
END if_pricelists_table_exist_test $$;

DO $$ << if_pricelist_items_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'pricelist_items'
) THEN RAISE EXCEPTION 'table inventory.pricelist_items not found';
END IF;
END if_pricelist_items_table_exist_test $$;

ROLLBACK;
