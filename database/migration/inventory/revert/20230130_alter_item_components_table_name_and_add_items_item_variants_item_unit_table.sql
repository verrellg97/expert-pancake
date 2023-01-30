-- Revert inventory:20230130_alter_item_components_table_name_and_add_items_item_variants_item_unit_table from pg

BEGIN;

ALTER TABLE "inventory"."brands" RENAME TO "item_brands";
ALTER TABLE "inventory"."groups" RENAME TO "item_groups";
ALTER TABLE "inventory"."units" RENAME TO "item_units";

DROP TABLE IF EXISTS "inventory"."item_units";
DROP TABLE IF EXISTS "inventory"."item_variants";
DROP TABLE IF EXISTS "inventory"."items";

COMMIT;
