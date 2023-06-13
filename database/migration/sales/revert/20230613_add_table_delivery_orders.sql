-- Revert sales:20230613_add_table_delivery_orders from pg

BEGIN;

DROP TABLE IF EXISTS "sales"."delivery_orders";
DROP TABLE IF EXISTS "sales"."delivery_order_items";

COMMIT;
