-- Revert warehouse:20230125_warehouse_schema_and_tables from pg

BEGIN;

DROP TABLE IF EXISTS "warehouse"."warehouse_racks";
DROP TABLE IF EXISTS "warehouse"."racks";
DROP TABLE IF EXISTS "warehouse"."warehouses";

DROP SCHEMA IF EXISTS warehouse;

COMMIT;
