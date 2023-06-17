-- Revert purchasing:20230618_add_table_receipt_orders from pg

BEGIN;

DROP TABLE IF EXISTS "purchasing"."receipt_order_items";
DROP TABLE IF EXISTS "purchasing"."receipt_orders";

COMMIT;
