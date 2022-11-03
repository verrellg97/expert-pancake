-- Revert account:20221103_alter_table_users_email_to_not_unique from pg

BEGIN;

ALTER TABLE "account"."users" ADD UNIQUE (email);

COMMIT;
