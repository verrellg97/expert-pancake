-- Deploy purchasing:20230411_alter_table_purchase_order to pg

BEGIN;

ALTER TABLE "purchasing"."purchase_orders" ADD COLUMN secondary_company_id text NOT NULL DEFAULT '';

COMMIT;
