-- Revert inventory:20230207_add_item_reorders_table from pg

BEGIN;

DROP TABLE IF EXISTS "inventory"."item_reorders";

COMMIT;
