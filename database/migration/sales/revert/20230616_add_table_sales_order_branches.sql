-- Revert sales:20230616_add_table_sales_order_branches from pg

BEGIN;

ALTER TABLE "sales"."sales_orders" DROP COLUMN is_all_branches;
DROP TABLE IF EXISTS "sales"."sales_order_branches";

COMMIT;
