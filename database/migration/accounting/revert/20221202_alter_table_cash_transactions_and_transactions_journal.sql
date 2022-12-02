-- Revert accounting:20221202_alter_table_cash_transactions_and_transactions_journal from pg

BEGIN;

ALTER TABLE "accounting"."transactions_journal" ADD COLUMN transaction_type TEXT NOT NULL DEFAULT '';
ALTER TABLE "accounting"."cash_transactions" ADD COLUMN transaction_type TEXT NOT NULL DEFAULT '';

COMMIT;
