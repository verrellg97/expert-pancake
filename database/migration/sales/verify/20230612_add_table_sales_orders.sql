-- Verify sales:20230612_add_table_sales_orders on pg

BEGIN;

DO $$ << if_sales_orders_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'sales_orders'
) THEN RAISE EXCEPTION 'table sales.sales_orders not found';
END IF;
END if_sales_orders_table_exist_test $$;

DO $$ << if_sales_order_items_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'sales_order_items'
) THEN RAISE EXCEPTION 'table sales.sales_order_items not found';
END IF;
END if_sales_order_items_table_exist_test $$;

ROLLBACK;
