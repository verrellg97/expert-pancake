-- Verify sales:20230613_add_table_delivery_orders on pg

BEGIN;

DO $$ << if_delivery_orders_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'delivery_orders'
) THEN RAISE EXCEPTION 'table sales.delivery_orders not found';
END IF;
END if_delivery_orders_table_exist_test $$;

DO $$ << if_delivery_order_items_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'delivery_order_items'
) THEN RAISE EXCEPTION 'table sales.delivery_order_items not found';
END IF;
END if_delivery_order_items_table_exist_test $$;

ROLLBACK;
