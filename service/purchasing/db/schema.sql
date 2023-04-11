CREATE SCHEMA IF NOT EXISTS purchasing;

CREATE TABLE "purchasing"."purchase_orders" (
  "id" text NOT NULL,
  "sales_order_id" text NOT NULL DEFAULT '',
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

CREATE TABLE "purchasing"."purchase_order_items" (
  "id" text NOT NULL,
  "sales_order_item_id" text NOT NULL DEFAULT '',
  "purchase_order_id" text NOT NULL,
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