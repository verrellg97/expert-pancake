-- Deploy inventory:20230310_alter_table_item_variant_maps to pg

BEGIN;

ALTER TABLE "inventory"."item_variant_maps" ADD COLUMN primary_company_id text DEFAULT '';
ALTER TABLE "inventory"."item_variant_maps" ADD COLUMN secondary_company_id text DEFAULT '';

COMMIT;
