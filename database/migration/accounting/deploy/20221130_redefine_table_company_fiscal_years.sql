-- Deploy accounting:20221130_redefine_table_company_fiscal_years to pg

BEGIN;

DROP TABLE IF EXISTS "accounting"."company_fiscal_years";

CREATE TABLE "accounting"."company_fiscal_years" (
  "company_id" text NOT NULL,
  "start_period" date NOT NULL DEFAULT CURRENT_DATE,
  "end_period" date NOT NULL DEFAULT CURRENT_DATE,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("company_id")
);

COMMIT;
