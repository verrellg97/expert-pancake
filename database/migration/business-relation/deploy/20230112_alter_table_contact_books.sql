-- Deploy business-relation:20230112_alter_table_contact_books to pg

BEGIN;

ALTER TABLE "business_relation"."contact_books" ADD COLUMN is_default bool NOT NULL DEFAULT false;
ALTER TABLE "business_relation"."contact_books" ADD COLUMN konekin_id text NOT NULL DEFAULT '';

COMMIT;
