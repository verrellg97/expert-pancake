-- Revert accounting:20221230_add_table_memorial_journals from pg

BEGIN;

DROP TABLE IF EXISTS "accounting"."memorial_journals";
DROP TABLE IF EXISTS "accounting"."memorial_journal_accounts";

COMMIT;
