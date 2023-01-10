-- Deploy business-relation:20230107_alter_table_add_primary_key to pg

BEGIN;

ALTER TABLE business_relation.contact_book_customer_infos
ADD CONSTRAINT contact_book_customer_infos_pkey PRIMARY KEY (contact_book_id);

ALTER TABLE business_relation.contact_book_supplier_infos
ADD CONSTRAINT contact_book_supplier_infos_pkey PRIMARY KEY (contact_book_id);

COMMIT;
