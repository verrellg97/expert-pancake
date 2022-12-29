-- Verify accounting:20221229_drop_table_fiscal_years_add_table_journal_books on pg

BEGIN;

DO $$ << if_journal_books_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'accounting'
        AND tablename = 'journal_books'
) THEN RAISE EXCEPTION 'table accounting.journal_books not found';
END IF;
END if_journal_books_table_exist_test $$;

DO $$ << if_journal_book_accounts_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'accounting'
        AND tablename = 'journal_book_accounts'
) THEN RAISE EXCEPTION 'table accounting.journal_book_accounts not found';
END IF;
END if_journal_book_accounts_table_exist_test $$;

ROLLBACK;
