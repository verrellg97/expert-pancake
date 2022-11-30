-- Revert accounting:20221130_redefine_table_company_fiscal_years from pg

BEGIN;

DROP TABLE IF EXISTS "accounting"."company_fiscal_years";

CREATE TABLE "accounting"."company_fiscal_years" (
  "company_id" text NOT NULL,
  "start_month" int NOT NULL DEFAULT 0,
  "start_year" int NOT NULL DEFAULT 0,
  "end_month" int NOT NULL DEFAULT 0,
  "end_year" int NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("company_id")
);

COMMIT;
