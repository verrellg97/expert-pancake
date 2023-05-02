-- Revert purchasing:20230502_alter_table_purchase_order from pg

BEGIN;

ALTER TABLE "purchasing"."purchase_orders" DROP COLUMN shipping_date;
ALTER TABLE "purchasing"."purchase_orders" DROP COLUMN receiving_warehouse_id;

COMMIT;
