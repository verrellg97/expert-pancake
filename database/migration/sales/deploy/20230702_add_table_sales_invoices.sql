-- Deploy sales:20230702_add_table_sales_invoices to pg

BEGIN;

ALTER TABLE "sales"."sales_orders" ADD COLUMN purchase_order_receiving_warehouse_id TEXT NOT NULL DEFAULT '';
ALTER TABLE "sales"."sales_order_items" ADD COLUMN amount_invoiced BIGINT NOT NULL DEFAULT 0;
ALTER TABLE "sales"."delivery_orders" RENAME COLUMN secondary_branch_id TO sales_order_id;

CREATE TABLE "sales"."sales_invoices" (
  "id" text NOT NULL,
  "sales_order_id" text NOT NULL,
  "purchase_invoice_id" text NOT NULL DEFAULT '',
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

CREATE TABLE "sales"."sales_invoice_items" (
  "id" text NOT NULL,
  "purchase_order_item_id" text NOT NULL DEFAULT '',
  "sales_order_item_id" text NOT NULL,
  "purchase_invoice_item_id" text NOT NULL DEFAULT '',
  "sales_invoice_id" text NOT NULL,
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
