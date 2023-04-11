-- Revert purchasing:20230411_alter_table_purchase_order from pg

BEGIN;

ALTER TABLE "purchasing"."purchase_orders" DROP COLUMN secondary_company_id;

COMMIT;
