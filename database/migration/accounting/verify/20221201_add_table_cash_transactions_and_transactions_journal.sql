-- Verify accounting:20221201_add_table_cash_transactions_and_transactions_journal on pg

BEGIN;

DO $$ << if_cash_transactions_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'accounting'
        AND tablename = 'cash_transactions'
) THEN RAISE EXCEPTION 'table accounting.cash_transactions not found';
END IF;
END if_cash_transactions_table_exist_test $$;

DO $$ << if_transactions_journal_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'accounting'
        AND tablename = 'transactions_journal'
) THEN RAISE EXCEPTION 'table accounting.transactions_journal not found';
END IF;
END if_transactions_journal_table_exist_test $$;

ROLLBACK;
