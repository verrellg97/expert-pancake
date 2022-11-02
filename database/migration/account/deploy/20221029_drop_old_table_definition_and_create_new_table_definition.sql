-- Deploy account:20221029_drop_old_table_definition_and_create_new_table_definition to pg

BEGIN;

DROP TABLE IF EXISTS "account"."user_infos";

DROP TABLE IF EXISTS "account"."users";

CREATE TABLE IF NOT EXISTS "account"."users" (
  "id" text UNIQUE PRIMARY KEY NOT NULL,
  "fullname" text NOT NULL,
  "nickname" text NOT NULL,
  "email" text UNIQUE,
  "phone_number" text NOT NULL UNIQUE,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "account"."user_passwords" (
  "user_id" text NOT NULL,
  "password" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("user_id")
);

CREATE TABLE IF NOT EXISTS "account"."user_infos" (
  "user_id" text NOT NULL,
  "key" text NOT NULL,
  "value" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("user_id", "key")
);

CREATE UNIQUE INDEX ON "account"."user_infos" ("user_id", "key");

COMMIT;
