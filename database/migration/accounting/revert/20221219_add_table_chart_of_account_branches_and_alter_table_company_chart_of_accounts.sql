-- Revert accounting:20221219_add_table_chart_of_account_branches_and_alter_table_company_chart_of_accounts from pg

BEGIN;

ALTER TABLE "accounting"."company_chart_of_accounts" ADD COLUMN branch_id TEXT NOT NULL DEFAULT '';
ALTER TABLE "accounting"."company_chart_of_accounts" ADD COLUMN opening_balance BIGINT NOT NULL DEFAULT 0;
ALTER TABLE "accounting"."company_chart_of_accounts" DROP COLUMN currency_code;
ALTER TABLE "accounting"."company_chart_of_accounts" DROP COLUMN report_type;
ALTER TABLE "accounting"."company_chart_of_accounts" DROP COLUMN account_type;
ALTER TABLE "accounting"."company_chart_of_accounts" DROP COLUMN is_all_branches;

DROP TABLE IF EXISTS "accounting"."chart_of_account_branches";

COMMIT;
