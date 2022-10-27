-- Deploy account:20221011_account_schema_and_table to pg
BEGIN;

CREATE SCHEMA IF NOT EXISTS account;

CREATE TABLE IF NOT EXISTS "account"."users" (
  "id" TEXT PRIMARY KEY NOT NULL,
  "first_name" TEXT NOT NULL,
  "last_name" TEXT NOT NULL,
  "email" TEXT NOT NULL,
  "phone_number" TEXT NOT NULL,
  "password" TEXT NOT NULL,
  "profile_picture_url" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ DEFAULT now(),
  "updated_at" TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "account"."user_infos" (
  "user_id" TEXT,
  "key" TEXT NOT NULL,
  "value" TEXT NOT NULL
);

ALTER TABLE
  "account"."user_infos"
ADD
  FOREIGN KEY ("user_id") REFERENCES "account"."users" ("id");

COMMIT;