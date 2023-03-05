-- Deploy inventory:20230305_alter_table_items to pg

BEGIN;

ALTER TABLE "inventory"."items" ALTER COLUMN brand_id SET DEFAULT '';

COMMIT;
