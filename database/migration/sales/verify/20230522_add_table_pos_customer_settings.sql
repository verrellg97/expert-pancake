-- Verify sales:20230522_add_table_pos_customer_settings on pg

BEGIN;

DO $$ << if_pos_customer_settings_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'pos_customer_settings'
) THEN RAISE EXCEPTION 'table sales.pos_customer_settings not found';
END IF;
END if_pos_customer_settings_table_exist_test $$;

ROLLBACK;
