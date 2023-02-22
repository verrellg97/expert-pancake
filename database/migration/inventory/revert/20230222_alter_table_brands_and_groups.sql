-- Revert inventory:20230222_alter_table_brands_and_groups from pg

BEGIN;

ALTER TABLE "inventory"."brands" DROP COLUMN is_deleted;
ALTER TABLE "inventory"."groups" DROP COLUMN is_deleted;

COMMIT;
