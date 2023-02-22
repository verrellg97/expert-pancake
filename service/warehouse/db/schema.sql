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

CREATE TABLE "warehouse"."warehouse_racks" (
  "id" text NOT NULL,
  "warehouse_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "is_deleted" bool NOT NULL DEFAULT FALSE,
  PRIMARY KEY ("id")
);