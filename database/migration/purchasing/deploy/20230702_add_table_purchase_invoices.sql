-- Deploy purchasing:20230702_add_table_purchase_invoices to pg

BEGIN;

ALTER TABLE "purchasing"."purchase_order_items" ADD COLUMN amount_received BIGINT NOT NULL DEFAULT 0;
ALTER TABLE "purchasing"."purchase_order_items" ADD COLUMN amount_invoiced BIGINT NOT NULL DEFAULT 0;
ALTER TABLE "purchasing"."receipt_orders" ADD COLUMN warehouse_id text NOT NULL DEFAULT '';
ALTER TABLE "purchasing"."receipt_order_items" DROP COLUMN amount_delivered;
ALTER TABLE "purchasing"."receipt_order_items" ALTER COLUMN purchase_order_item_id SET DEFAULT '';
ALTER TABLE "purchasing"."receipt_order_items" ALTER COLUMN item_barcode_id SET DEFAULT '';

CREATE TABLE "purchasing"."purchase_invoices" (
  "id" text NOT NULL,
  "sales_invoice_id" text NOT NULL DEFAULT '',
  "receipt_order_id" text NOT NULL DEFAULT '',
  "company_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "form_number" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "contact_book_id" text NOT NULL,
  "secondary_company_id" text NOT NULL DEFAULT '',
  "konekin_id" text NOT NULL DEFAULT '',
  "currency_code" text NOT NULL,
  "total_items" bigint NOT NULL DEFAULT 0,
  "total" bigint NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "status" text NOT NULL DEFAULT 'created',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "purchasing"."purchase_invoice_items" (
  "id" text NOT NULL,
  "purchase_order_item_id" text NOT NULL DEFAULT '',
  "sales_order_item_id" text NOT NULL DEFAULT '',
  "sales_invoice_item_id" text NOT NULL DEFAULT '',
  "receipt_order_item_id" text NOT NULL DEFAULT '',
  "purchase_invoice_id" text NOT NULL,
  "primary_item_variant_id" text NOT NULL,
  "secondary_item_variant_id" text NOT NULL DEFAULT '',
  "primary_item_unit_id" text NOT NULL,
  "secondary_item_unit_id" text NOT NULL DEFAULT '',
  "primary_item_unit_value" bigint NOT NULL DEFAULT 0,
  "secondary_item_unit_value" bigint NOT NULL DEFAULT 0,
  "amount" bigint NOT NULL DEFAULT 0,
  "price" bigint NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
