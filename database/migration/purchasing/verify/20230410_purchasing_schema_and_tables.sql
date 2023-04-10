-- Verify purchasing:20230410_purchasing_schema_and_tables on pg

BEGIN;

DO $$ << if_purchasing_schema_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM information_schema.schemata
    WHERE schema_name = 'purchasing'
) THEN RAISE EXCEPTION 'schema purchasing not found';
END IF;
END if_purchasing_schema_exist_test $$;

DO $$ << if_purchase_orders_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'purchasing'
        AND tablename = 'purchase_orders'
) THEN RAISE EXCEPTION 'table purchasing.purchase_orders not found';
END IF;
END if_purchase_orders_table_exist_test $$;

DO $$ << if_purchase_order_items_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'purchasing'
        AND tablename = 'purchase_order_items'
) THEN RAISE EXCEPTION 'table purchasing.purchase_order_items not found';
END IF;
END if_purchase_order_items_table_exist_test $$;

ROLLBACK;
