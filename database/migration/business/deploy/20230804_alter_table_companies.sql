-- Deploy business:20230804_alter_table_companies to pg

BEGIN;

ALTER TABLE "business"."companies" ADD COLUMN image_url TEXT NOT NULL DEFAULT '';

COMMIT;
