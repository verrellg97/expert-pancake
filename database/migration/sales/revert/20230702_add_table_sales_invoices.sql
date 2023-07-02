-- Revert sales:20230702_add_table_sales_invoices from pg

BEGIN;

ALTER TABLE "sales"."sales_orders" DROP COLUMN purchase_order_receiving_warehouse_id;
ALTER TABLE "sales"."sales_order_items" DROP COLUMN amount_invoiced;
ALTER TABLE "sales"."delivery_orders" RENAME COLUMN sales_order_id TO secondary_branch_id;

DROP TABLE IF EXISTS "sales"."sales_invoice_items";
DROP TABLE IF EXISTS "sales"."sales_invoices";

COMMIT;
