-- Revert account:20230731_alter_table_users from pg

BEGIN;

ALTER TABLE "account"."users" DROP COLUMN image_url;

COMMIT;
