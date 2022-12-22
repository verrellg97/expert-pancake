-- Deploy accounting:20221222_alter_table_chart_of_accounts to pg

BEGIN;

ALTER TABLE "accounting"."company_chart_of_accounts" DROP COLUMN report_type;
ALTER TABLE "accounting"."company_chart_of_accounts" DROP COLUMN account_type;
ALTER TABLE "accounting"."company_chart_of_accounts" DROP COLUMN account_group;
ALTER TABLE "accounting"."company_chart_of_accounts" ADD COLUMN chart_of_account_group_id TEXT NOT NULL DEFAULT '';

COMMIT;
