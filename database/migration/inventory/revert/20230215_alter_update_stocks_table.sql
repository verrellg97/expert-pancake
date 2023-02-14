-- Revert inventory:20230215_alter_update_stocks_table from pg

BEGIN;

ALTER TABLE "inventory"."update_stocks" DROP COLUMN batch;
ALTER TABLE "inventory"."update_stocks" DROP COLUMN expired_date;
ALTER TABLE "inventory"."update_stocks" DROP COLUMN item_barcode_id;

COMMIT;
