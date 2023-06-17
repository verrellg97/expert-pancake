-- Verify purchasing:20230618_add_table_receipt_orders on pg

BEGIN;

DO $$ << if_receipt_orders_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'purchasing'
        AND tablename = 'receipt_orders'
) THEN RAISE EXCEPTION 'table purchasing.receipt_orders not found';
END IF;
END if_receipt_orders_table_exist_test $$;

DO $$ << if_receipt_order_items_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'purchasing'
        AND tablename = 'receipt_order_items'
) THEN RAISE EXCEPTION 'table purchasing.receipt_order_items not found';
END IF;
END if_receipt_order_items_table_exist_test $$;

ROLLBACK;
