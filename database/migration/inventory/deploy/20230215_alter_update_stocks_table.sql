-- Deploy inventory:20230215_alter_update_stocks_table to pg

BEGIN;

ALTER TABLE "inventory"."update_stocks" ADD COLUMN batch text;
ALTER TABLE "inventory"."update_stocks" ADD COLUMN expired_date date;
ALTER TABLE "inventory"."update_stocks" ADD COLUMN item_barcode_id text NOT NULL DEFAULT '';

COMMIT;
