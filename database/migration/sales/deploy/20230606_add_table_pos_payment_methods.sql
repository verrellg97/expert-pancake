-- Deploy sales:20230606_add_table_pos_payment_methods to pg

BEGIN;

CREATE TABLE "sales"."pos_payment_methods" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "chart_of_account_id" text NOT NULL,
  "name" text NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
