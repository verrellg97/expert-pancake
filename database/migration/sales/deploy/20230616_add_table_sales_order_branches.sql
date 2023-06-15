-- Deploy sales:20230616_add_table_sales_order_branches to pg

BEGIN;

ALTER TABLE "sales"."sales_orders" ADD COLUMN is_all_branches BOOL NOT NULL DEFAULT FALSE;

CREATE TABLE "sales"."sales_order_branches" (
  "sales_order_id" text NOT NULL,
  "company_branch_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

COMMIT;
