-- Revert business:20221128_alter_table_company_and_company_branches_table from pg

BEGIN;

ALTER TABLE "business"."companies" ALTER COLUMN user_id DROP NOT NULL;
ALTER TABLE "business"."company_branches" ALTER COLUMN user_id DROP NOT NULL, ALTER COLUMN company_id DROP NOT NULL;

ALTER TABLE "business"."companies" ALTER COLUMN is_deleted DROP DEFAULT;
ALTER TABLE "business"."companies" ALTER COLUMN is_deleted TYPE INT USING CASE WHEN is_deleted = FALSE THEN 0 ELSE 1 END;
ALTER TABLE "business"."companies" ALTER COLUMN is_deleted SET DEFAULT 0;

ALTER TABLE "business"."company_branches" ALTER COLUMN is_deleted DROP DEFAULT;
ALTER TABLE "business"."company_branches" ALTER COLUMN is_deleted TYPE INT USING CASE WHEN is_deleted = FALSE THEN 0 ELSE 1 END;
ALTER TABLE "business"."company_branches" ALTER COLUMN is_deleted SET DEFAULT 0;

ALTER TABLE "business"."company_branches" DROP COLUMN is_central;

COMMIT;
