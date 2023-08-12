-- Deploy purchasing:20230813_alter_table_purchase_orders to pg

BEGIN;

ALTER TABLE "purchasing"."purchase_orders" ADD COLUMN payment_term INT NOT NULL DEFAULT 0;

COMMIT;
