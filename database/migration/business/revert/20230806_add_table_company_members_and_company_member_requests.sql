-- Revert business:20230806_add_table_company_members_and_company_member_requests from pg

BEGIN;

DROP TABLE IF EXISTS "business"."company_members";

DROP TABLE IF EXISTS "business"."company_member_requests";

COMMIT;
