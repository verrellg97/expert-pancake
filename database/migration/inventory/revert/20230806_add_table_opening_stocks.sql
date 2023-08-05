-- Revert inventory:20230806_add_table_opening_stocks from pg

BEGIN;

DROP TABLE IF EXISTS "inventory"."opening_stocks";

COMMIT;
