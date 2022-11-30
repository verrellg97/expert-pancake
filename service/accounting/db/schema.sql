CREATE SCHEMA IF NOT EXISTS accounting;

CREATE TABLE accounting.company_fiscal_years (
  "company_id" text NOT NULL,
  "start_period" date NOT NULL DEFAULT CURRENT_DATE,
  "end_period" date NOT NULL DEFAULT CURRENT_DATE,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("company_id")
);

CREATE TABLE accounting.company_chart_of_accounts (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "account_code" text NOT NULL,
  "account_name" text NOT NULL,
  "account_group" text NOT NULL,
  "bank_name" text NOT NULL,
  "bank_account_number" text NOT NULL,
  "bank_code" text NOT NULL,
  "opening_balance" bigint NOT NULL DEFAULT 0,
  "is_deleted" boolean NOT NULL DEFAULT FALSE,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);