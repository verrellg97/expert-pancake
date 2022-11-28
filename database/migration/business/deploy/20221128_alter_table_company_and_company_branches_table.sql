-- Deploy business:20221128_alter_table_company_and_company_branches_table to pg

BEGIN;

ALTER TABLE "business"."companies" ALTER COLUMN user_id SET NOT NULL;
ALTER TABLE "business"."company_branches" ALTER COLUMN user_id SET NOT NULL, ALTER COLUMN company_id SET NOT NULL;

ALTER TABLE "business"."companies" ALTER COLUMN is_deleted DROP DEFAULT;
ALTER TABLE "business"."companies" ALTER COLUMN is_deleted TYPE BOOLEAN USING CASE WHEN is_deleted = 0 THEN FALSE ELSE TRUE END;
ALTER TABLE "business"."companies" ALTER COLUMN is_deleted SET DEFAULT FALSE;

ALTER TABLE "business"."company_branches" ALTER COLUMN is_deleted DROP DEFAULT;
ALTER TABLE "business"."company_branches" ALTER COLUMN is_deleted TYPE BOOLEAN USING CASE WHEN is_deleted = 0 THEN FALSE ELSE TRUE END;
ALTER TABLE "business"."company_branches" ALTER COLUMN is_deleted SET DEFAULT FALSE;

ALTER TABLE "business"."company_branches" ADD COLUMN is_central BOOLEAN NOT NULL DEFAULT FALSE;

COMMIT;
