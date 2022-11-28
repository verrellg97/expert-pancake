-- Deploy accounting:20221128_alter_table_company_chart_of_accounts to pg

BEGIN;

ALTER TABLE "accounting"."company_chart_of_accounts" ALTER COLUMN is_deleted DROP DEFAULT;
ALTER TABLE "accounting"."company_chart_of_accounts" ALTER COLUMN is_deleted TYPE BOOLEAN USING CASE WHEN is_deleted = 0 THEN FALSE ELSE TRUE END;
ALTER TABLE "accounting"."company_chart_of_accounts" ALTER COLUMN is_deleted SET DEFAULT FALSE;

COMMIT;
