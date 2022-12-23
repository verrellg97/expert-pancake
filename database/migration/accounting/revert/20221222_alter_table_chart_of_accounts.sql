-- Revert accounting:20221222_alter_table_chart_of_accounts from pg

BEGIN;

ALTER TABLE "accounting"."company_chart_of_accounts" ADD COLUMN report_type TEXT NOT NULL DEFAULT '';
ALTER TABLE "accounting"."company_chart_of_accounts" ADD COLUMN account_type TEXT NOT NULL DEFAULT '';
ALTER TABLE "accounting"."company_chart_of_accounts" ADD COLUMN account_group TEXT NOT NULL DEFAULT '';
ALTER TABLE "accounting"."company_chart_of_accounts" DROP COLUMN chart_of_account_group_id;

COMMIT;
