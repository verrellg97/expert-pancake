-- Deploy warehouse:20230201_drop_old_table_definition_and_create_new_table_definition to pg

BEGIN;

DROP TABLE IF EXISTS "warehouse"."warehouse_racks";
DROP TABLE IF EXISTS "warehouse"."racks";
DROP TABLE IF EXISTS "warehouse"."warehouses";

CREATE TABLE "warehouse"."warehouses" (
  "id" text NOT NULL,
  "branch_id" text NOT NULL,
  "code" text NOT NULL,
  "name" text NOT NULL,
  "address" text NOT NULL,
  "type" text NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "warehouse"."warehouse_racks" (
  "id" text NOT NULL,
  "warehouse_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
