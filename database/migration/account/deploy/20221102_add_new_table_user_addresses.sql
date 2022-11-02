-- Deploy account:20221102_add_new_table_user_addresses to pg

BEGIN;

CREATE TABLE "account"."user_addresses" (
  "user_id" text UNIQUE PRIMARY KEY NOT NULL,
  "country" text NOT NULL,
  "province" text NOT NULL,
  "regency" text NOT NULL,
  "district" text NOT NULL,
  "full_address" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);


COMMIT;
