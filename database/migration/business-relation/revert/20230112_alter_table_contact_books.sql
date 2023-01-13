-- Revert business-relation:20230112_alter_table_contact_books from pg

BEGIN;

ALTER TABLE "business_relation"."contact_books" DROP COLUMN is_default;
ALTER TABLE "business_relation"."contact_books" DROP COLUMN konekin_id;

COMMIT;
