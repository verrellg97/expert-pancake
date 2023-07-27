CREATE SCHEMA IF NOT EXISTS inventory;

CREATE TABLE "inventory"."brands" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "is_deleted" bool NOT NULL DEFAULT FALSE,
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."groups" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "is_deleted" bool NOT NULL DEFAULT FALSE,
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."units" (
  "id" text NOT NULL,
  "unit_category_id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."items" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "image_url" text NOT NULL DEFAULT '',
  "code" text NOT NULL,
  "name" text NOT NULL,
  "brand_id" text NOT NULL DEFAULT '',
  "group_id" text NOT NULL,
  "tag" text NOT NULL DEFAULT '',
  "description" text NOT NULL DEFAULT '',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."item_infos" (
  "item_id" text NOT NULL,
  "is_purchase" bool NOT NULL DEFAULT false,
  "is_sale" bool NOT NULL DEFAULT false,
  "is_raw_material" bool NOT NULL DEFAULT false,
  "is_asset" bool NOT NULL DEFAULT false,
  "purchase_chart_of_account_id" text NOT NULL,
  "sale_chart_of_account_id" text NOT NULL,
  "purchase_item_unit_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("item_id")
);

CREATE TABLE "inventory"."item_variants" (
  "id" text NOT NULL,
  "item_id" text NOT NULL,
  "image_url" text NOT NULL DEFAULT '',
  "barcode" text NOT NULL DEFAULT '',
  "name" text NOT NULL,
  "price" bigint NOT NULL DEFAULT 0,
  "is_default" bool NOT NULL DEFAULT false,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."item_variant_maps" (
  "id" text NOT NULL,
  "primary_company_id" text NOT NULL,
  "secondary_company_id" text NOT NULL,
  "primary_item_variant_id" text NOT NULL,
  "secondary_item_variant_id" text NOT NULL,
  "primary_item_unit_id" text NOT NULL,
  "secondary_item_unit_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."item_units" (
  "id" text NOT NULL,
  "item_id" text NOT NULL,
  "unit_id" text NOT NULL,
  "value" bigint NOT NULL DEFAULT 0,
  "is_default" bool NOT NULL DEFAULT false,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."internal_stock_transfers" (
  "id" text NOT NULL,
  "source_warehouse_id" text NOT NULL,
  "destination_warehouse_id" text NOT NULL,
  "form_number" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "is_received" bool NOT NULL DEFAULT false,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."internal_stock_transfer_items" (
  "id" text NOT NULL,
  "internal_stock_transfer_id" text NOT NULL,
  "warehouse_rack_id" text NOT NULL,
  "variant_id" text NOT NULL,
  "item_unit_id" text NOT NULL,
  "item_unit_value" bigint NOT NULL DEFAULT 0,
  "amount" bigint NOT NULL DEFAULT 0,
  "batch" text,
  "expired_date" date,
  "item_barcode_id" text NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."item_barcodes" (
  "id" text NOT NULL,
  "variant_id" text NOT NULL,
  "batch" text,
  "expired_date" date,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."stock_movements" (
  "id" text NOT NULL,
  "transaction_id" text NOT NULL,
  "transaction_code" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "transaction_reference" text NOT NULL,
  "detail_transaction_id" text NOT NULL DEFAULT '',
  "warehouse_id" text NOT NULL,
  "warehouse_rack_id" text NOT NULL,
  "variant_id" text NOT NULL,
  "item_barcode_id" text NOT NULL,
  "amount" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."update_stocks" (
  "id" text NOT NULL,
  "form_number" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "warehouse_id" text NOT NULL,
  "warehouse_rack_id" text NOT NULL,
  "variant_id" text NOT NULL,
  "item_unit_id" text NOT NULL,
  "item_unit_value" bigint NOT NULL DEFAULT 0,
  "beginning_stock" bigint NOT NULL DEFAULT 0,
  "ending_stock" bigint NOT NULL DEFAULT 0,
  "batch" text,
  "expired_date" date,
  "item_barcode_id" text NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."item_reorders" (
  "id" text NOT NULL,
  "warehouse_id" text NOT NULL,
  "item_unit_id" text NOT NULL,
  "variant_id" text NOT NULL,
  "minimum_stock" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."unit_categories" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."pricelists" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "start_date" date NOT NULL DEFAULT CURRENT_DATE,
  "end_date" date,
  "is_default" bool NOT NULL DEFAULT false,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."pricelist_items" (
  "pricelist_id" text NOT NULL,
  "variant_id" text NOT NULL,
  "item_unit_id" text NOT NULL,
  "price" bigint NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("pricelist_id", "variant_id", "item_unit_id")
);