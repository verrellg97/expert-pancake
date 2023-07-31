-- Deploy account:20230731_alter_table_users to pg

BEGIN;

ALTER TABLE "account"."users" ADD COLUMN image_url TEXT NOT NULL DEFAULT '';

COMMIT;
