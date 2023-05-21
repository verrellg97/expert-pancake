-- Deploy sales:20230522_add_table_pos_customer_settings to pg

BEGIN;

CREATE TABLE "sales"."pos_customer_settings" (
  "branch_id" text NOT NULL,
  "contact_book_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("branch_id", "contact_book_id")
);

COMMIT;
