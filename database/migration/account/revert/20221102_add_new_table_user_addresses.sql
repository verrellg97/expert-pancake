-- Revert account:20221102_add_new_table_user_addresses from pg

BEGIN;

DROP TABLE IF EXISTS "account"."user_addresses";

COMMIT;
