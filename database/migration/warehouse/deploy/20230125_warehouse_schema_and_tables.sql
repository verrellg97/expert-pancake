-- Deploy warehouse:20230125_warehouse_schema_and_tables to pg

BEGIN;

CREATE SCHEMA IF NOT EXISTS warehouse;

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

CREATE TABLE "warehouse"."racks" (
  "id" text NOT NULL,
  "branch_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "warehouse"."warehouse_racks" (
  "warehouse_id" text NOT NULL,
  "rack_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

COMMIT;
