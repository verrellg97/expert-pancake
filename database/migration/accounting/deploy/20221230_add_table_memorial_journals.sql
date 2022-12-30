-- Deploy accounting:20221230_add_table_memorial_journals to pg

BEGIN;

CREATE TABLE "accounting"."memorial_journals" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "description" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "accounting"."memorial_journal_accounts" (
  "memorial_journal_id" text NOT NULL,
  "chart_of_account_id" text NOT NULL,
  "debit_amount" bigint NOT NULL DEFAULT 0,
  "credit_amount" bigint NOT NULL DEFAULT 0,
  "description" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

COMMIT;
