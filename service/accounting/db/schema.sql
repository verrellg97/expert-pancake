CREATE SCHEMA IF NOT EXISTS accounting;

CREATE TABLE accounting.chart_of_account_groups (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "report_type" text NOT NULL,
  "account_type" text NOT NULL,
  "account_group_name" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE accounting.company_chart_of_accounts (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "currency_code" text NOT NULL,
  "chart_of_account_group_id" text NOT NULL,
  "account_code" text NOT NULL,
  "account_name" text NOT NULL,
  "bank_name" text NOT NULL,
  "bank_account_number" text NOT NULL,
  "bank_code" text NOT NULL,
  "is_deleted" boolean NOT NULL DEFAULT FALSE,
  "is_all_branches" boolean NOT NULL DEFAULT FALSE,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE accounting.chart_of_account_branches (
  "chart_of_account_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE accounting.journal_books (
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

CREATE TABLE accounting.journal_book_accounts (
  "journal_book_id" text NOT NULL,
  "chart_of_account_id" text NOT NULL,
  "amount" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE accounting.memorial_journals (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "description" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE accounting.memorial_journal_accounts (
  "memorial_journal_id" text NOT NULL,
  "chart_of_account_id" text NOT NULL,
  "debit_amount" bigint NOT NULL DEFAULT 0,
  "credit_amount" bigint NOT NULL DEFAULT 0,
  "description" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE accounting.cash_transactions (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "type" text NOT NULL,
  "main_chart_of_account_id" text NOT NULL,
  "contra_chart_of_account_id" text NOT NULL,
  "amount" bigint NOT NULL DEFAULT 0,
  "description" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE accounting.transactions_journal (
  "company_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "transaction_id" text NOT NULL,
  "transaction_date" date NOT NULL DEFAULT CURRENT_DATE,
  "transaction_reference" text NOT NULL,
  "chart_of_account_id" text NOT NULL,
  "amount" bigint NOT NULL DEFAULT 0,
  "description" text NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);