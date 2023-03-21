-- Revert inventory:20230320_alter_table_item_units from pg

BEGIN;

ALTER TABLE "inventory"."item_units" DROP COLUMN is_deleted;

COMMIT;
