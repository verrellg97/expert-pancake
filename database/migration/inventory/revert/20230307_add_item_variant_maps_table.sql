-- Revert inventory:20230307_add_item_variant_maps_table from pg

BEGIN;

DROP TABLE IF EXISTS "inventory"."item_variant_maps";

COMMIT;
