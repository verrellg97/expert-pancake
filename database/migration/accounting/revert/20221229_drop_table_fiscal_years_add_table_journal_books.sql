-- Revert accounting:20221229_drop_table_fiscal_years_add_table_journal_books from pg

BEGIN;

DROP TABLE IF EXISTS "accounting"."journal_books";
DROP TABLE IF EXISTS "accounting"."journal_book_accounts";

CREATE TABLE "accounting"."company_fiscal_years" (
  "company_id" text NOT NULL,
  "start_period" date NOT NULL DEFAULT CURRENT_DATE,
  "end_period" date NOT NULL DEFAULT CURRENT_DATE,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("company_id")
);

COMMIT;
