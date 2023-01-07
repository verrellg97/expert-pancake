-- Revert business-relation:20230106_business_relation_schema_and_table from pg

BEGIN;

DROP TABLE IF EXISTS "business_relation"."contact_groups";
DROP TABLE IF EXISTS "business_relation"."contact_books";
DROP TABLE IF EXISTS "business_relation"."contact_book_additional_infos";
DROP TABLE IF EXISTS "business_relation"."contact_book_mailing_addresses";
DROP TABLE IF EXISTS "business_relation"."contact_book_shipping_addresses";
DROP TABLE IF EXISTS "business_relation"."contact_book_branches";
DROP TABLE IF EXISTS "business_relation"."contact_book_customer_infos";
DROP TABLE IF EXISTS "business_relation"."contact_book_supplier_infos";

DROP SCHEMA IF EXISTS business_relation;

COMMIT;
