-- Revert inventory:20230310_alter_table_item_variant_maps from pg

BEGIN;

ALTER TABLE "inventory"."item_variant_maps" DROP COLUMN primary_company_id;
ALTER TABLE "inventory"."item_variant_maps" DROP COLUMN secondary_company_id;

COMMIT;
