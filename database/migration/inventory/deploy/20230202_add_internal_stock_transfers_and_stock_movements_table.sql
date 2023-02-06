-- Deploy inventory:20230202_add_internal_stock_transfers_and_stock_movements_table to pg

BEGIN;

CREATE TABLE "inventory"."internal_stock_transfers" (
  "id" text NOT NULL,
  "source_warehouse_id" text NOT NULL,
  "destination_warehouse_id" text NOT NULL,
  "form_number" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
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

COMMIT;
