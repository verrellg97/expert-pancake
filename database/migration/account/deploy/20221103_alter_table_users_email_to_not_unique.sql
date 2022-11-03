-- Deploy account:20221103_alter_table_users_email_to_not_unique to pg

BEGIN;

ALTER TABLE "account"."users" DROP CONSTRAINT IF EXISTS users_email_key;

COMMIT;
