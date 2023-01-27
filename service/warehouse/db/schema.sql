CREATE SCHEMA IF NOT EXISTS warehouse;

CREATE TABLE warehouse.racks (
  "id" text NOT NULL,
  "branch_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);