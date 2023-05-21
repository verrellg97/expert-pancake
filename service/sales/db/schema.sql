CREATE SCHEMA IF NOT EXISTS sales;

CREATE TABLE "sales"."point_of_sales" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "warehouse_id" text NOT NULL,
  "form_number" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "contact_book_id" text NOT NULL,
  "secondary_company_id" text NOT NULL DEFAULT '',
  "konekin_id" text NOT NULL DEFAULT '',
  "currency_code" text NOT NULL,
  "chart_of_account_id" text NOT NULL,
  "total_items" bigint NOT NULL DEFAULT 0,
  "total" bigint NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "sales"."point_of_sale_items" (
  "id" text NOT NULL,
  "point_of_sale_id" text NOT NULL,
  "warehouse_rack_id" text NOT NULL,
  "item_variant_id" text NOT NULL,
  "item_unit_id" text NOT NULL,
  "item_unit_value" bigint NOT NULL DEFAULT 0,
  "batch" text,
  "expired_date" date,
  "item_barcode_id" text NOT NULL,
  "amount" bigint NOT NULL DEFAULT 0,
  "price" bigint NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "sales"."pos_user_settings" (
  "user_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "warehouse_id" text NOT NULL,
  "warehouse_rack_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("user_id", "branch_id")
);

CREATE TABLE "sales"."pos_chart_of_account_settings" (
  "branch_id" text NOT NULL,
  "chart_of_account_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("branch_id", "chart_of_account_id")
);

CREATE TABLE "sales"."pos_customer_settings" (
  "branch_id" text NOT NULL,
  "contact_book_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("branch_id", "contact_book_id")
);