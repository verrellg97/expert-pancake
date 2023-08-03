-- Revert business:20230804_alter_table_companies from pg

BEGIN;

ALTER TABLE "business"."companies" DROP COLUMN image_url;

COMMIT;
