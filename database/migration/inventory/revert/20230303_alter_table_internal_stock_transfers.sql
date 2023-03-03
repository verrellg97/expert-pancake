-- Revert inventory:20230303_alter_table_internal_stock_transfers from pg

BEGIN;

ALTER TABLE "inventory"."internal_stock_transfers" DROP COLUMN is_received;

COMMIT;
