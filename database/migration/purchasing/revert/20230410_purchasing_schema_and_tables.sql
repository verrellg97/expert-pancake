-- Revert purchasing:20230410_purchasing_schema_and_tables from pg

BEGIN;

DROP TABLE IF EXISTS "purchasing"."purchase_order_items";
DROP TABLE IF EXISTS "purchasing"."purchase_orders";

DROP SCHEMA IF EXISTS purchasing;

COMMIT;
