-- Revert inventory:20230728_alter_table_item_variants from pg

BEGIN;

ALTER TABLE "inventory"."item_variants" DROP COLUMN is_deleted;

COMMIT;
