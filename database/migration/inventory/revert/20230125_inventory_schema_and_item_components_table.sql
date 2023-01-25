-- Revert inventory:20230125_inventory_schema_and_item_components_table from pg

BEGIN;

DROP TABLE IF EXISTS "inventory"."item_brands";
DROP TABLE IF EXISTS "inventory"."item_groups";
DROP TABLE IF EXISTS "inventory"."item_units";

DROP SCHEMA IF EXISTS inventory;

COMMIT;
