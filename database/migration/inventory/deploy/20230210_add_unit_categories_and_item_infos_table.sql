-- Deploy inventory:20230210_add_unit_categories_and_item_infos_table to pg

BEGIN;

CREATE TABLE "inventory"."unit_categories" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

ALTER TABLE "inventory"."units" ADD COLUMN unit_category_id text NOT NULL DEFAULT '';

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

ALTER TABLE "inventory"."item_variants" ADD COLUMN barcode text NOT NULL DEFAULT '';
ALTER TABLE "inventory"."item_variants" DROP COLUMN stock;

ALTER TABLE "inventory"."item_reorders" ADD COLUMN item_unit_id text NOT NULL DEFAULT '';

COMMIT;
