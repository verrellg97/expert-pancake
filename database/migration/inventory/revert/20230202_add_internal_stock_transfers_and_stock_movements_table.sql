-- Revert inventory:20230202_add_internal_stock_transfers_and_stock_movements_table from pg

BEGIN;

DROP TABLE IF EXISTS "inventory"."stock_movements";
DROP TABLE IF EXISTS "inventory"."item_barcodes";
DROP TABLE IF EXISTS "inventory"."internal_stock_transfer_items";
DROP TABLE IF EXISTS "inventory"."internal_stock_transfers";

COMMIT;
