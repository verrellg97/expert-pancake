-- Revert warehouse:20230222_alter_table_warehouse_racks from pg

BEGIN;

ALTER TABLE "warehouse"."warehouse_racks" DROP COLUMN is_deleted;

COMMIT;
