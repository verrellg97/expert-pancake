-- Revert sales:20230815_alter_table_delivery_orders from pg

BEGIN;

ALTER TABLE "sales"."delivery_orders" DROP COLUMN warehouse_id;

COMMIT;
