CREATE SCHEMA IF NOT EXISTS business;

CREATE TABLE business.companies (
  "id" text NOT NULL,
  "user_id" text,
  "name" text NOT NULL,
  "initial_name" text NOT NULL,
  "type" text NOT NULL,
  "responsible_person" text NOT NULL,
  "is_deleted" int DEFAULT 0 NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE business.company_branches (
  "id" text NOT NULL,
  "user_id" text,
  "company_id" text,
  "name" text NOT NULL,
  "address" text NOT NULL,
  "phone_number" text NOT NULL,
  "is_deleted" int DEFAULT 0 NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);