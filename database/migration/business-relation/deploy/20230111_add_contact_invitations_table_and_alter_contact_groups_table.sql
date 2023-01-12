-- Deploy business-relation:20230111_add_contact_invitations_table_and_alter_contact_groups_table to pg

BEGIN;

ALTER TABLE "business_relation"."contact_groups" ADD COLUMN image_url TEXT NOT NULL DEFAULT '';
ALTER TABLE "business_relation"."contact_groups" ADD COLUMN description TEXT NOT NULL DEFAULT '';

CREATE TABLE "business_relation"."contact_invitations" (
  "id" text NOT NULL,
  "primary_contact_book_id" text NOT NULL DEFAULT '',
  "secondary_contact_book_id" text NOT NULL DEFAULT '',
  "primary_company_id" text NOT NULL,
  "secondary_company_id" text NOT NULL,
  "status" text NOT NULL DEFAULT 'waiting',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

COMMIT;
