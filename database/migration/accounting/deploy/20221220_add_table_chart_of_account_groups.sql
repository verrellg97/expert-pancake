-- Deploy accounting:20221220_add_table_chart_of_account_groups to pg

BEGIN;

CREATE TABLE "accounting"."chart_of_account_groups" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "report_type" text NOT NULL,
  "account_type" text NOT NULL,
  "account_group_name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
