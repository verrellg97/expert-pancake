-- Revert inventory:20230731_alter_table_update_stocks_and_internal_stock_transfers from pg

BEGIN;

ALTER TABLE "inventory"."internal_stock_transfers" DROP COLUMN company_id;
ALTER TABLE "inventory"."internal_stock_transfers" DROP COLUMN branch_id;

ALTER TABLE "inventory"."update_stocks" DROP COLUMN company_id;
ALTER TABLE "inventory"."update_stocks" DROP COLUMN branch_id;

COMMIT;
