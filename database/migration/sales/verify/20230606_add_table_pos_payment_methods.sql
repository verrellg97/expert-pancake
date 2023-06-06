-- Verify sales:20230606_add_table_pos_payment_methods on pg

BEGIN;

DO $$ << if_pos_payment_methods_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'pos_payment_methods'
) THEN RAISE EXCEPTION 'table sales.pos_payment_methods not found';
END IF;
END if_pos_payment_methods_table_exist_test $$;

ROLLBACK;
