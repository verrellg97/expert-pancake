-- Revert inventory:20230322_add_pricelists_and_pricelist_items_table from pg

BEGIN;

DROP TABLE IF EXISTS "inventory"."pricelist_items";
DROP TABLE IF EXISTS "inventory"."pricelists";

COMMIT;
