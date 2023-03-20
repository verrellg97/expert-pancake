-- Revert inventory:202303020_alter_table_item_units from pg

BEGIN;

ALTER TABLE "inventory"."item_units" DROP COLUMN is_deleted;

COMMIT;
