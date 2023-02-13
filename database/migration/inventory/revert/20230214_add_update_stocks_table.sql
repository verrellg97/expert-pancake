-- Revert inventory:20230214_add_update_stocks_table from pg

BEGIN;

DROP TABLE IF EXISTS "inventory"."update_stocks";

COMMIT;
