-- Deploy inventory:20230323_alter_table_pricelists_and_pricelist_items to pg

BEGIN;

ALTER TABLE "inventory"."pricelists" ALTER COLUMN is_default SET NOT NULL;

ALTER TABLE "inventory"."pricelist_items" DROP CONSTRAINT pricelist_items_pkey;
ALTER TABLE "inventory"."pricelist_items" ADD COLUMN item_unit_id TEXT NOT NULL DEFAULT '';
ALTER TABLE "inventory"."pricelist_items" ADD PRIMARY KEY (pricelist_id, variant_id, item_unit_id);

COMMIT;
