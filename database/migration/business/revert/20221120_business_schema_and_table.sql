-- Revert business:20221120_business_schema_and_table from pg

BEGIN;

DROP TABLE IF EXISTS "business"."companies";

DROP TABLE IF EXISTS "business"."company_branches";

DROP SCHEMA IF EXISTS business;

COMMIT;
