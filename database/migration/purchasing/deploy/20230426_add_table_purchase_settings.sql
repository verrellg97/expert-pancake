-- Deploy purchasing:20230426_add_table_purchase_settings to pg

BEGIN;

CREATE TABLE "purchasing"."purchase_settings" (
  "company_id" text NOT NULL,
  "is_auto_approve_order" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("company_id")
);

COMMIT;
