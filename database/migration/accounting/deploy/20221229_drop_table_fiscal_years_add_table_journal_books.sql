-- Deploy accounting:20221229_drop_table_fiscal_years_add_table_journal_books to pg

BEGIN;

DROP TABLE IF EXISTS "accounting"."company_fiscal_years";

CREATE TABLE "accounting"."journal_books" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "name" text NOT NULL,
  "start_period" date NOT NULL DEFAULT CURRENT_DATE,
  "end_period" date NOT NULL DEFAULT CURRENT_DATE,
  "is_closed" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "accounting"."journal_book_accounts" (
  "journal_book_id" text NOT NULL,
  "chart_of_account_id" text NOT NULL,
  "amount" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

COMMIT;
