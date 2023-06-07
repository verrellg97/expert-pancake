-- Deploy sales:20230607_alter_table_pos to pg

BEGIN;

ALTER TABLE "sales"."point_of_sales" RENAME COLUMN chart_of_account_id TO pos_payment_method_id;

COMMIT;
