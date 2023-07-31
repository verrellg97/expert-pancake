-- Revert inventory:20230731_alter_table_stock_movements from pg

BEGIN;

ALTER TABLE "inventory"."stock_movements" DROP COLUMN company_id;
ALTER TABLE "inventory"."stock_movements" DROP COLUMN branch_id;

COMMIT;
