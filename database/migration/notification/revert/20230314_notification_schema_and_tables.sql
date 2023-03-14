-- Revert notification:20230314_notification_schema_and_tables from pg

BEGIN;

DROP TABLE IF EXISTS "notification"."notifications";

DROP SCHEMA IF EXISTS notification;

COMMIT;
