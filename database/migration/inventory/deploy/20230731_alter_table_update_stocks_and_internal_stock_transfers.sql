-- Deploy inventory:20230731_alter_table_update_stocks_and_internal_stock_transfers to pg

BEGIN;

ALTER TABLE "inventory"."internal_stock_transfers" ADD COLUMN company_id TEXT NOT NULL DEFAULT '';
ALTER TABLE "inventory"."internal_stock_transfers" ADD COLUMN branch_id TEXT NOT NULL DEFAULT '';

ALTER TABLE "inventory"."update_stocks" ADD COLUMN company_id TEXT NOT NULL DEFAULT '';
ALTER TABLE "inventory"."update_stocks" ADD COLUMN branch_id TEXT NOT NULL DEFAULT '';

COMMIT;
