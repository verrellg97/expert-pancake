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
  "pos_payment_method_id" text NOT NULL,
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

CREATE TABLE "sales"."pos_payment_methods" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "chart_of_account_id" text NOT NULL,
  "name" text NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "sales"."sales_orders" (
  "id" text NOT NULL,
  "purchase_order_id" text NOT NULL DEFAULT '',
  "purchase_order_branch_id" text NOT NULL DEFAULT '',
  "purchase_order_receiving_warehouse_id" text NOT NULL DEFAULT '',
  "company_id" text NOT NULL,
  "branch_id" text NOT NULL DEFAULT '',
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
  "is_all_branches" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "sales"."sales_order_branches" (
  "sales_order_id" text NOT NULL,
  "company_branch_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "sales"."sales_order_items" (
  "id" text NOT NULL,
  "purchase_order_item_id" text NOT NULL DEFAULT '',
  "sales_order_id" text NOT NULL,
  "primary_item_variant_id" text NOT NULL,
  "secondary_item_variant_id" text NOT NULL DEFAULT '',
  "primary_item_unit_id" text NOT NULL,
  "secondary_item_unit_id" text NOT NULL DEFAULT '',
  "primary_item_unit_value" bigint NOT NULL DEFAULT 0,
  "secondary_item_unit_value" bigint NOT NULL DEFAULT 0,
  "amount" bigint NOT NULL DEFAULT 0,
  "amount_sent" bigint NOT NULL DEFAULT 0,
  "amount_invoiced" bigint NOT NULL DEFAULT 0,
  "price" bigint NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "sales"."delivery_orders" (
  "id" text NOT NULL,
  "sales_order_id" text NOT NULL DEFAULT '',
  "receipt_order_id" text NOT NULL DEFAULT '',
  "company_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "form_number" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "contact_book_id" text NOT NULL,
  "secondary_company_id" text NOT NULL DEFAULT '',
  "konekin_id" text NOT NULL DEFAULT '',
  "total_items" bigint NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "status" text NOT NULL DEFAULT 'created',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "sales"."delivery_order_items" (
  "id" text NOT NULL,
  "purchase_order_item_id" text NOT NULL DEFAULT '',
  "sales_order_item_id" text NOT NULL,
  "receipt_order_item_id" text NOT NULL DEFAULT '',
  "delivery_order_id" text NOT NULL,
  "primary_item_variant_id" text NOT NULL,
  "warehouse_rack_id" text NOT NULL,
  "batch" text,
  "expired_date" date,
  "item_barcode_id" text NOT NULL,
  "secondary_item_variant_id" text NOT NULL DEFAULT '',
  "primary_item_unit_id" text NOT NULL,
  "secondary_item_unit_id" text NOT NULL DEFAULT '',
  "primary_item_unit_value" bigint NOT NULL DEFAULT 0,
  "secondary_item_unit_value" bigint NOT NULL DEFAULT 0,
  "amount" bigint NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

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