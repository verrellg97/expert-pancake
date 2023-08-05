-- Deploy inventory:20230806_add_table_opening_stocks to pg

BEGIN;

CREATE TABLE "inventory"."opening_stocks" (
  "id" text NOT NULL,
  "form_number" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "company_id" text NOT NULL DEFAULT '',
  "branch_id" text NOT NULL DEFAULT '',
  "warehouse_id" text NOT NULL,
  "warehouse_rack_id" text NOT NULL,
  "variant_id" text NOT NULL,
  "item_unit_id" text NOT NULL,
  "item_unit_value" bigint NOT NULL DEFAULT 0,
  "amount" bigint NOT NULL DEFAULT 0,
  "price" bigint NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
