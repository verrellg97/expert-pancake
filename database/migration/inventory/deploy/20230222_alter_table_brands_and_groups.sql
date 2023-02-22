-- Deploy inventory:20230222_alter_table_brands_and_groups to pg

BEGIN;

ALTER TABLE "inventory"."brands" ADD COLUMN is_deleted bool NOT NULL DEFAULT false;
ALTER TABLE "inventory"."groups" ADD COLUMN is_deleted bool NOT NULL DEFAULT false;

COMMIT;
