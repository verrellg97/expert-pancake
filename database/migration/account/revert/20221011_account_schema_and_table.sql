-- Revert account:20221011_account_schema_and_table from pg
BEGIN;

DROP TABLE IF EXISTS "account"."user_infos";

DROP TABLE IF EXISTS "account"."users";

DROP SCHEMA IF EXISTS account;

COMMIT;