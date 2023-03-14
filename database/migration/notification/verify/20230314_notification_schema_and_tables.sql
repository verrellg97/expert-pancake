-- Verify notification:20230314_notification_schema_and_tables on pg

BEGIN;

DO $$ << if_notification_schema_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM information_schema.schemata
    WHERE schema_name = 'notification'
) THEN RAISE EXCEPTION 'schema notification not found';
END IF;
END if_notification_schema_exist_test $$;

DO $$ << if_notifications_table_exist_test >> BEGIN IF NOT EXISTS(
    SELECT 1
    FROM pg_tables
    WHERE schemaname = 'notification'
        AND tablename = 'notifications'
) THEN RAISE EXCEPTION 'table notification.notifications not found';
END IF;
END if_notifications_table_exist_test $$;

ROLLBACK;
