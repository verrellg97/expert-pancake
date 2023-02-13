-- Revert inventory:20230210_add_unit_categories_and_item_infos_table from pg

BEGIN;

DROP TABLE IF EXISTS "inventory"."unit_categories";

ALTER TABLE "inventory"."units" DROP COLUMN unit_category_id;

DROP TABLE IF EXISTS "inventory"."item_infos";

ALTER TABLE "inventory"."item_variants" DROP COLUMN barcode;
ALTER TABLE "inventory"."item_variants" ADD COLUMN stock BIGINT NOT NULL DEFAULT 0;

ALTER TABLE "inventory"."item_reorders" DROP COLUMN item_unit_id;

COMMIT;
