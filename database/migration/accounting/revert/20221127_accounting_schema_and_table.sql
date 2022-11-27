-- Revert accounting:20221127_accounting_schema_and_table from pg

BEGIN;

DROP TABLE IF EXISTS "accounting"."company_fiscal_years";

DROP TABLE IF EXISTS "accounting"."company_chart_of_accounts";

DROP SCHEMA IF EXISTS accounting;

COMMIT;
