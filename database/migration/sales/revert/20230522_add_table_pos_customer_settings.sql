-- Revert sales:20230522_add_table_pos_customer_settings from pg

BEGIN;

DROP TABLE IF EXISTS "sales"."pos_customer_settings";

COMMIT;
