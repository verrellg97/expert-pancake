-- Deploy inventory:20230320_alter_table_item_units to pg

BEGIN;

ALTER TABLE "inventory"."item_units" ADD COLUMN is_deleted bool NOT NULL DEFAULT false;

COMMIT;
