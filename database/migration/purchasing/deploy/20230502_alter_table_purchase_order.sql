-- Deploy purchasing:20230502_alter_table_purchase_order to pg

BEGIN;

ALTER TABLE "purchasing"."purchase_orders" ADD COLUMN shipping_date DATE NOT NULL DEFAULT (now());
ALTER TABLE "purchasing"."purchase_orders" ADD COLUMN receiving_warehouse_id text NOT NULL DEFAULT '';

COMMIT;
