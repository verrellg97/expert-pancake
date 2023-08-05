-- Revert inventory:20230806_alter_table_opening_stocks from pg

BEGIN;

ALTER TABLE "inventory"."opening_stocks" DROP COLUMN batch;
ALTER TABLE "inventory"."opening_stocks" DROP COLUMN expired_date;
ALTER TABLE "inventory"."opening_stocks" DROP COLUMN item_barcode_id;

COMMIT;
