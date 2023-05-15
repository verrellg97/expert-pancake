-- Verify sales:20230514_sales_schema_and_tables on pg

BEGIN;

DO $$ << if_sales_schema_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM information_schema.schemata
    WHERE schema_name = 'sales'
) THEN RAISE EXCEPTION 'schema sales not found';
END IF;
END if_sales_schema_exist_test $$;

DO $$ << if_point_of_sales_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'point_of_sales'
) THEN RAISE EXCEPTION 'table sales.point_of_sales not found';
END IF;
END if_point_of_sales_table_exist_test $$;

DO $$ << if_point_of_sale_items_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'point_of_sale_items'
) THEN RAISE EXCEPTION 'table sales.point_of_sale_items not found';
END IF;
END if_point_of_sale_items_table_exist_test $$;

DO $$ << if_pos_user_settings_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'pos_user_settings'
) THEN RAISE EXCEPTION 'table sales.pos_user_settings not found';
END IF;
END if_pos_user_settings_table_exist_test $$;

DO $$ << if_pos_chart_of_account_settings_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'sales'
        AND tablename = 'pos_chart_of_account_settings'
) THEN RAISE EXCEPTION 'table sales.pos_chart_of_account_settings not found';
END IF;
END if_pos_chart_of_account_settings_table_exist_test $$;

ROLLBACK;
