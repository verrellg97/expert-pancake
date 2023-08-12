-- Deploy sales:20230813_alter_table_sales_orders to pg

BEGIN;

ALTER TABLE "sales"."sales_orders" ADD COLUMN payment_term INT NOT NULL DEFAULT 0;

COMMIT;
