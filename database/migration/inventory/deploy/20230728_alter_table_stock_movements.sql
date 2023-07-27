-- Deploy inventory:20230728_alter_table_stock_movements to pg

BEGIN;

ALTER TABLE "inventory"."stock_movements" ADD COLUMN transaction_code TEXT NOT NULL DEFAULT '';

COMMIT;
