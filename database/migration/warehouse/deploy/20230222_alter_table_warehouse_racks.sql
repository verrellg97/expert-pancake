-- Deploy warehouse:20230222_alter_table_warehouse_racks to pg

BEGIN;

ALTER TABLE "warehouse"."warehouse_racks" ADD COLUMN is_deleted bool NOT NULL DEFAULT false;

COMMIT;
