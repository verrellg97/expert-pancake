-- Revert sales:20230514_sales_schema_and_tables from pg

BEGIN;

DROP TABLE IF EXISTS "sales"."pos_chart_of_account_settings";
DROP TABLE IF EXISTS "sales"."pos_user_settings";
DROP TABLE IF EXISTS "sales"."point_of_sale_items";
DROP TABLE IF EXISTS "sales"."point_of_sales";

DROP SCHEMA IF EXISTS sales;

COMMIT;
