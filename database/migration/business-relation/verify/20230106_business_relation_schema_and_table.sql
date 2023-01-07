-- Verify business-relation:20230106_business_relation_schema_and_table on pg

BEGIN;

DO $$ << if_business_relation_schema_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM information_schema.schemata
    WHERE schema_name = 'business_relation'
) THEN RAISE EXCEPTION 'schema business_relation not found';
END IF;
END if_business_relation_schema_exist_test $$;

DO $$ << if_contact_groups_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business_relation'
        AND tablename = 'contact_groups'
) THEN RAISE EXCEPTION 'table business_relation.contact_groups not found';
END IF;
END if_contact_groups_table_exist_test $$;

DO $$ << if_contact_books_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business_relation'
        AND tablename = 'contact_books'
) THEN RAISE EXCEPTION 'table business_relation.contact_books not found';
END IF;
END if_contact_books_table_exist_test $$;

DO $$ << if_contact_book_additional_infos_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business_relation'
        AND tablename = 'contact_book_additional_infos'
) THEN RAISE EXCEPTION 'table business_relation.contact_book_additional_infos not found';
END IF;
END if_contact_book_additional_infos_table_exist_test $$;

DO $$ << if_contact_book_mailing_addresses_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business_relation'
        AND tablename = 'contact_book_mailing_addresses'
) THEN RAISE EXCEPTION 'table business_relation.contact_book_mailing_addresses not found';
END IF;
END if_contact_book_mailing_addresses_table_exist_test $$;

DO $$ << if_contact_book_shipping_addresses_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business_relation'
        AND tablename = 'contact_book_shipping_addresses'
) THEN RAISE EXCEPTION 'table business_relation.contact_book_shipping_addresses not found';
END IF;
END if_contact_book_shipping_addresses_table_exist_test $$;

DO $$ << if_contact_book_branches_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business_relation'
        AND tablename = 'contact_book_branches'
) THEN RAISE EXCEPTION 'table business_relation.contact_book_branches not found';
END IF;
END if_contact_book_branches_table_exist_test $$;

DO $$ << if_contact_book_customer_infos_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business_relation'
        AND tablename = 'contact_book_customer_infos'
) THEN RAISE EXCEPTION 'table business_relation.contact_book_customer_infos not found';
END IF;
END if_contact_book_customer_infos_table_exist_test $$;

DO $$ << if_contact_book_supplier_infos_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business_relation'
        AND tablename = 'contact_book_supplier_infos'
) THEN RAISE EXCEPTION 'table business_relation.contact_book_supplier_infos not found';
END IF;
END if_contact_book_supplier_infos_table_exist_test $$;

ROLLBACK;
