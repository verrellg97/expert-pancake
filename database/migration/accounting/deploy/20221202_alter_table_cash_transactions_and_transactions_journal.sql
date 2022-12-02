-- Deploy accounting:20221202_alter_table_cash_transactions_and_transactions_journal to pg

BEGIN;

ALTER TABLE "accounting"."transactions_journal" DROP COLUMN transaction_type;
ALTER TABLE "accounting"."cash_transactions" DROP COLUMN transaction_type;

COMMIT;
