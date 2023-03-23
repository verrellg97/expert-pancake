-- Revert inventory:20230323_alter_table_pricelists_and_pricelist_items from pg

BEGIN;


ALTER TABLE "inventory"."pricelists" ALTER COLUMN is_default DROP NOT NULL;

ALTER TABLE "inventory"."pricelist_items" DROP CONSTRAINT pricelist_items_pkey;
ALTER TABLE "inventory"."pricelist_items" DROP COLUMN item_unit_id;
ALTER TABLE "inventory"."pricelist_items" ADD PRIMARY KEY (pricelist_id, variant_id);

COMMIT;
