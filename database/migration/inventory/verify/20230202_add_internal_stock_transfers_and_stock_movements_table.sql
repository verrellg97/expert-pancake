-- Verify inventory:20230202_add_internal_stock_transfers_and_stock_movements_table on pg

BEGIN;

DO $$ << if_internal_stock_transfers_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'internal_stock_transfers'
) THEN RAISE EXCEPTION 'table inventory.internal_stock_transfers not found';
END IF;
END if_internal_stock_transfers_table_exist_test $$;

DO $$ << if_internal_stock_transfer_items_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'internal_stock_transfer_items'
) THEN RAISE EXCEPTION 'table inventory.internal_stock_transfer_items not found';
END IF;
END if_internal_stock_transfer_items_table_exist_test $$;

DO $$ << if_item_barcodes_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'item_barcodes'
) THEN RAISE EXCEPTION 'table inventory.item_barcodes not found';
END IF;
END if_item_barcodes_table_exist_test $$;

DO $$ << if_stock_movements_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'inventory'
        AND tablename = 'stock_movements'
) THEN RAISE EXCEPTION 'table inventory.stock_movements not found';
END IF;
END if_stock_movements_table_exist_test $$;

ROLLBACK;
