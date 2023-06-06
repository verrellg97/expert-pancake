-- Revert sales:20230606_add_table_pos_payment_methods from pg

BEGIN;

DROP TABLE IF EXISTS "sales"."pos_payment_methods";

COMMIT;
