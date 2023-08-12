-- Revert sales:20230813_alter_table_sales_orders from pg

BEGIN;

ALTER TABLE "sales"."sales_orders" DROP COLUMN payment_term;

COMMIT;
