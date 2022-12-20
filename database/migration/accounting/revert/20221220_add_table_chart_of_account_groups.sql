-- Revert accounting:20221220_add_table_chart_of_account_groups from pg

BEGIN;

DROP TABLE IF EXISTS "accounting"."chart_of_account_groups";

COMMIT;
