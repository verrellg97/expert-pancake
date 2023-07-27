-- Revert inventory:20230728_alter_table_stock_movements from pg

BEGIN;

ALTER TABLE "inventory"."stock_movements" DROP COLUMN transaction_code;

COMMIT;
