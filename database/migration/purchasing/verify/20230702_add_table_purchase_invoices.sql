-- Verify purchasing:20230702_add_table_purchase_invoices on pg

BEGIN;

DO $$ << if_purchase_invoices_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'purchasing'
        AND tablename = 'purchase_invoices'
) THEN RAISE EXCEPTION 'table purchasing.purchase_invoices not found';
END IF;
END if_purchase_invoices_table_exist_test $$;

DO $$ << if_purchase_invoice_items_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'purchasing'
        AND tablename = 'purchase_invoice_items'
) THEN RAISE EXCEPTION 'table purchasing.purchase_invoice_items not found';
END IF;
END if_purchase_invoice_items_table_exist_test $$;

ROLLBACK;
