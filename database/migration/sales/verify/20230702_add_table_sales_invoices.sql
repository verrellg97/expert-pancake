-- Verify sales:20230702_add_table_sales_invoices on pg

BEGIN;

DO $$ << if_sales_invoices_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'sales_invoices'
) THEN RAISE EXCEPTION 'table sales.sales_invoices not found';
END IF;
END if_sales_invoices_table_exist_test $$;

DO $$ << if_sales_invoice_items_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'sales_invoice_items'
) THEN RAISE EXCEPTION 'table sales.sales_invoice_items not found';
END IF;
END if_sales_invoice_items_table_exist_test $$;

ROLLBACK;
