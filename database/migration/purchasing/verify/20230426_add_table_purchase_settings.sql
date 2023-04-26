-- Verify purchasing:20230426_add_table_purchase_settings on pg

BEGIN;

DO $$ << if_purchase_settings_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'purchasing'
        AND tablename = 'purchase_settings'
) THEN RAISE EXCEPTION 'table purchasing.purchase_settings not found';
END IF;
END if_purchase_settings_table_exist_test $$;

ROLLBACK;
