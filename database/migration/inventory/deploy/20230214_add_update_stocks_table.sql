-- Deploy inventory:20230214_add_update_stocks_table to pg

BEGIN;

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
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
