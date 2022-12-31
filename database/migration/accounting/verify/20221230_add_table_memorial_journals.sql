-- Verify accounting:20221230_add_table_memorial_journals on pg

BEGIN;

DO $$ << if_memorial_journals_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'accounting'
        AND tablename = 'memorial_journals'
) THEN RAISE EXCEPTION 'table accounting.memorial_journals not found';
END IF;
END if_memorial_journals_table_exist_test $$;

DO $$ << if_memorial_journal_accounts_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'accounting'
        AND tablename = 'memorial_journal_accounts'
) THEN RAISE EXCEPTION 'table accounting.memorial_journal_accounts not found';
END IF;
END if_memorial_journal_accounts_table_exist_test $$;

ROLLBACK;
