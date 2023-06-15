-- Verify sales:20230616_add_table_sales_order_branches on pg

BEGIN;

DO $$ << if_sales_order_branches_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'sales_order_branches'
) THEN RAISE EXCEPTION 'table sales.sales_order_branches not found';
END IF;
END if_sales_order_branches_table_exist_test $$;

ROLLBACK;
