-- Revert accounting:20230102_drop_table_journal_book_accounts_alter_table_memorial_journal_accounts from pg

BEGIN;

CREATE TABLE "accounting"."journal_book_accounts" (
  "journal_book_id" text NOT NULL,
  "chart_of_account_id" text NOT NULL,
  "amount" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

ALTER TABLE "accounting"."memorial_journal_accounts" ALTER COLUMN memorial_journal_id DROP DEFAULT;
ALTER TABLE "accounting"."memorial_journal_accounts" DROP COLUMN journal_book_id;

COMMIT;
