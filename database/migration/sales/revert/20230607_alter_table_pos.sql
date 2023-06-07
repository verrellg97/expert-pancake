-- Revert sales:20230607_alter_table_pos from pg

BEGIN;

ALTER TABLE "sales"."point_of_sales" RENAME COLUMN pos_payment_method_id TO chart_of_account_id;

COMMIT;
