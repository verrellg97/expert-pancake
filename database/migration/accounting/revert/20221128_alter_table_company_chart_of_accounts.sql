-- Revert accounting:20221128_alter_table_company_chart_of_accounts from pg

BEGIN;

ALTER TABLE "accounting"."company_chart_of_accounts" ALTER COLUMN is_deleted DROP DEFAULT;
ALTER TABLE "accounting"."company_chart_of_accounts" ALTER COLUMN is_deleted TYPE INT USING CASE WHEN is_deleted = FALSE THEN 0 ELSE 1 END;
ALTER TABLE "accounting"."company_chart_of_accounts" ALTER COLUMN is_deleted SET DEFAULT 0;

ALTER TABLE "accounting"."company_chart_of_accounts" ALTER COLUMN opening_balance TYPE FLOAT;

COMMIT;
