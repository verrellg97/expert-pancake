-- Revert purchasing:20230813_alter_table_purchase_orders from pg

BEGIN;

ALTER TABLE "purchasing"."purchase_orders" DROP COLUMN payment_term;

COMMIT;
