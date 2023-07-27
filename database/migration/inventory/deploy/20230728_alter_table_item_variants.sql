-- Deploy inventory:20230728_alter_table_item_variants to pg

BEGIN;

ALTER TABLE "inventory"."item_variants" ADD COLUMN is_deleted bool DEFAULT false;

COMMIT;
