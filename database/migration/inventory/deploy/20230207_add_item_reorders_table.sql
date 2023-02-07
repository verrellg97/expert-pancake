-- Deploy inventory:20230207_add_item_reorders_table to pg

BEGIN;

CREATE TABLE "inventory"."item_reorders" (
  "id" text NOT NULL,
  "warehouse_id" text NOT NULL,
  "variant_id" text NOT NULL,
  "minimum_stock" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
