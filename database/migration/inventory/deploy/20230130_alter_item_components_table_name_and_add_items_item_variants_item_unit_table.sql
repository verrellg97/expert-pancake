-- Deploy inventory:20230130_alter_item_components_table_name_and_add_items_item_variants_item_unit_table to pg

BEGIN;

ALTER TABLE "inventory"."item_brands" RENAME TO "brands";
ALTER TABLE "inventory"."item_groups" RENAME TO "groups";
ALTER TABLE "inventory"."item_units" RENAME TO "units";

CREATE TABLE "inventory"."items" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "image_url" text NOT NULL DEFAULT '',
  "code" text NOT NULL,
  "name" text NOT NULL,
  "brand_id" text NOT NULL,
  "group_id" text NOT NULL,
  "tag" text NOT NULL DEFAULT '',
  "description" text NOT NULL DEFAULT '',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "inventory"."item_variants" (
  "id" text NOT NULL,
  "item_id" text NOT NULL,
  "image_url" text NOT NULL DEFAULT '',
  "name" text NOT NULL,
  "price" bigint NOT NULL DEFAULT 0,
  "stock" bigint NOT NULL DEFAULT 0,
  "is_default" bool NOT NULL DEFAULT false,
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
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
