-- Deploy sales:20230815_alter_table_delivery_orders to pg

BEGIN;

ALTER TABLE "sales"."delivery_orders" ADD COLUMN warehouse_id text NOT NULL DEFAULT '';

COMMIT;
