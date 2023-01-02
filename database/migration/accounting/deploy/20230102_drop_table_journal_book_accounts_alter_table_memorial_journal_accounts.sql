-- Deploy accounting:20230102_drop_table_journal_book_accounts_alter_table_memorial_journal_accounts to pg

BEGIN;

DROP TABLE IF EXISTS "accounting"."journal_book_accounts";
ALTER TABLE "accounting"."memorial_journal_accounts" ALTER COLUMN memorial_journal_id SET DEFAULT '';
ALTER TABLE "accounting"."memorial_journal_accounts" ADD COLUMN journal_book_id TEXT NOT NULL DEFAULT '';

COMMIT;
