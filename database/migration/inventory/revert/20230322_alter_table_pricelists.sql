-- Revert inventory:20230322_alter_table_pricelists from pg

BEGIN;

ALTER TABLE "inventory"."pricelists" DROP COLUMN is_default;

COMMIT;
