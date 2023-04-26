-- Revert purchasing:20230426_add_table_purchase_settings from pg

BEGIN;

DROP TABLE IF EXISTS "purchasing"."purchase_settings";

COMMIT;
