-- Verify business-relation:20230111_add_contact_invitations_table_and_alter_contact_groups_table on pg

BEGIN;

DO $$ << if_contact_invitations_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'business_relation'
        AND tablename = 'contact_invitations'
) THEN RAISE EXCEPTION 'table business_relation.contact_invitations not found';
END IF;
END if_contact_invitations_table_exist_test $$;

ROLLBACK;
