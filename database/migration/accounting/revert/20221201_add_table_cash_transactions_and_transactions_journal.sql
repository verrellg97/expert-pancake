-- Revert accounting:20221201_add_table_cash_transactions_and_transactions_journal from pg

BEGIN;

DROP TABLE IF EXISTS "accounting"."transactions_journal";
DROP TABLE IF EXISTS "accounting"."cash_transactions";

COMMIT;
