-- Revert sales:20230612_add_table_sales_orders from pg

BEGIN;

DROP TABLE IF EXISTS "sales"."sales_orders";
DROP TABLE IF EXISTS "sales"."sales_order_items";

COMMIT;
