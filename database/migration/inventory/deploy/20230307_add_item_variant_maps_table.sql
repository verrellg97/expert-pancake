-- Deploy inventory:20230307_add_item_variant_maps_table to pg

BEGIN;

CREATE TABLE "inventory"."item_variant_maps" (
  "id" text NOT NULL,
  "primary_item_variant_id" text NOT NULL,
  "secondary_item_variant_id" text NOT NULL,
  "primary_item_unit_id" text NOT NULL,
  "secondary_item_unit_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
