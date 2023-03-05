-- Revert inventory:20230305_alter_table_items from pg

BEGIN;

ALTER TABLE "inventory"."items" ALTER COLUMN brand_id DROP DEFAULT;

COMMIT;
