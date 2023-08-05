-- Deploy business:20230806_add_table_company_members_and_company_member_requests to pg

BEGIN;

CREATE TABLE "business"."company_members" (
  "id" text NOT NULL,
  "user_id" text NOT NULL,
  "company_id" text NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "business"."company_member_requests" (
  "id" text NOT NULL,
  "user_id" text NOT NULL,
  "company_id" text NOT NULL,
  "status" text NOT NULL DEFAULT 'waiting',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);


COMMIT;
