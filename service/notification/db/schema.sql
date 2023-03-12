CREATE SCHEMA IF NOT EXISTS notification;
CREATE TABLE "notification"."notifications" (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "branch_id" text NOT NULL,
  "title" text NOT NULL,
  "content" text NOT NULL,
  "type" text NOT NULL,
  "is_read" bool NOT NULL DEFAULT false,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);