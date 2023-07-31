-- Deploy inventory:20230731_alter_table_stock_movements to pg

BEGIN;

ALTER TABLE "inventory"."stock_movements" ADD COLUMN company_id TEXT NOT NULL DEFAULT '';
ALTER TABLE "inventory"."stock_movements" ADD COLUMN branch_id TEXT NOT NULL DEFAULT '';

COMMIT;
