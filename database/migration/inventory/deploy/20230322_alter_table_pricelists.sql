-- Deploy inventory:20230322_alter_table_pricelists to pg

BEGIN;

ALTER TABLE "inventory"."pricelists" ADD COLUMN is_default bool DEFAULT false;

COMMIT;
