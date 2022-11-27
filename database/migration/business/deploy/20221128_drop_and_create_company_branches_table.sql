-- Deploy business:20221128_drop_and_create_company_branches_table to pg

BEGIN;

DROP TABLE IF EXISTS "business"."company_branches";

CREATE TABLE "business"."company_branches" (
  "id" text NOT NULL,
  "user_id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "address" text NOT NULL,
  "phone_number" text NOT NULL,
  "is_central" int DEFAULT 0 NOT NULL,
  "is_deleted" int DEFAULT 0 NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
