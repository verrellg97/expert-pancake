-- Revert purchasing:20230702_add_table_purchase_invoices from pg

BEGIN;

ALTER TABLE "purchasing"."purchase_order_items" DROP COLUMN amount_received;
ALTER TABLE "purchasing"."purchase_order_items" DROP COLUMN amount_invoiced;
ALTER TABLE "purchasing"."receipt_orders" DROP COLUMN warehouse_id;
ALTER TABLE "purchasing"."receipt_order_items" ADD COLUMN amount_delivered BIGINT NOT NULL DEFAULT 0;
ALTER TABLE "purchasing"."receipt_order_items" ALTER COLUMN purchase_order_item_id DROP DEFAULT;
ALTER TABLE "purchasing"."receipt_order_items" ALTER COLUMN item_barcode_id DROP DEFAULT;

DROP TABLE IF EXISTS "purchasing"."purchase_invoice_items";
DROP TABLE IF EXISTS "purchasing"."purchase_invoices";

COMMIT;
