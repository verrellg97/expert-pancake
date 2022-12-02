-- Deploy accounting:20221201_add_table_cash_transactions_and_transactions_journal to pg

BEGIN;

CREATE TABLE "accounting"."cash_transactions" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "transaction_type" text NOT NULL,
  "type" text NOT NULL,
  "main_chart_of_account_id" text NOT NULL,
  "contra_chart_of_account_id" text NOT NULL,
  "amount" bigint NOT NULL DEFAULT 0,
  "description" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "accounting"."transactions_journal" (
  "company_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "transaction_id" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "transaction_reference" text NOT NULL,
  "transaction_type" text NOT NULL,
  "chart_of_account_id" text NOT NULL,
  "amount" bigint NOT NULL DEFAULT 0,
  "description" text NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

COMMIT;
