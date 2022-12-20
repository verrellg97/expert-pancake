-- Deploy accounting:20221219_add_table_chart_of_account_branches_and_alter_table_company_chart_of_accounts to pg

BEGIN;

ALTER TABLE "accounting"."company_chart_of_accounts" DROP COLUMN branch_id;
ALTER TABLE "accounting"."company_chart_of_accounts" DROP COLUMN opening_balance;
ALTER TABLE "accounting"."company_chart_of_accounts" ADD COLUMN currency_code TEXT NOT NULL DEFAULT '';
ALTER TABLE "accounting"."company_chart_of_accounts" ADD COLUMN report_type TEXT NOT NULL DEFAULT '';
ALTER TABLE "accounting"."company_chart_of_accounts" ADD COLUMN account_type TEXT NOT NULL DEFAULT '';
ALTER TABLE "accounting"."company_chart_of_accounts" ADD COLUMN is_all_branches BOOL NOT NULL DEFAULT FALSE;

CREATE TABLE "accounting"."chart_of_account_branches" (
  "chart_of_account_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

COMMIT;
