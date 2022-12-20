-- Verify accounting:20221220_add_table_chart_of_account_groups on pg

BEGIN;

DO $$ << if_chart_of_account_groups_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'accounting'
        AND tablename = 'chart_of_account_groups'
) THEN RAISE EXCEPTION 'table accounting.chart_of_account_groups not found';
END IF;
END if_chart_of_account_groups_table_exist_test $$;

ROLLBACK;
