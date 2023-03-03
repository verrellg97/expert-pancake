-- Deploy inventory:20230303_alter_table_internal_stock_transfers to pg

BEGIN;

ALTER TABLE "inventory"."internal_stock_transfers" ADD COLUMN is_received bool NOT NULL DEFAULT false;

COMMIT;
