-- Revert business-relation:20230107_alter_table_add_primary_key from pg

BEGIN;

ALTER TABLE business_relation.contact_book_customer_infos DROP CONSTRAINT contact_book_customer_infos_pkey;
ALTER TABLE business_relation.contact_book_supplier_infos DROP CONSTRAINT contact_book_supplier_infos_pkey;

COMMIT;
