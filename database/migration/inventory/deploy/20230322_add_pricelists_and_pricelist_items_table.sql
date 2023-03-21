-- Deploy inventory:20230322_add_pricelists_and_pricelist_items_table to pg

BEGIN;

CREATE TABLE "inventory"."pricelists" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "start_date" date NOT NULL DEFAULT CURRENT_DATE,
  "end_date" date,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."pricelist_items" (
  "pricelist_id" text NOT NULL,
  "variant_id" text NOT NULL,
  "price" bigint NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("pricelist_id", "variant_id")
);

COMMIT;
