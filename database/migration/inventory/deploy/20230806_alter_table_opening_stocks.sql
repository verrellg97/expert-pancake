-- Deploy inventory:20230806_alter_table_opening_stocks to pg

BEGIN;

ALTER TABLE "inventory"."opening_stocks" ADD COLUMN batch text;
ALTER TABLE "inventory"."opening_stocks" ADD COLUMN expired_date date;
ALTER TABLE "inventory"."opening_stocks" ADD COLUMN item_barcode_id text NOT NULL DEFAULT '';

COMMIT;
