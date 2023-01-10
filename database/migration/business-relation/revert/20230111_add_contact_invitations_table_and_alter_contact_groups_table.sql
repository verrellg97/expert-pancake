-- Revert business-relation:20230111_add_contact_invitations_table_and_alter_contact_groups_table from pg

BEGIN;

ALTER TABLE "business_relation"."contact_groups" DROP COLUMN image_url;
ALTER TABLE "business_relation"."contact_groups" DROP COLUMN description;

DROP TABLE "business_relation"."contact_invitations";
DROP TYPE "business_relation"."contact_invitation_status";

COMMIT;
