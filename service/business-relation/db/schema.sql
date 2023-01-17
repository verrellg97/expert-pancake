CREATE SCHEMA IF NOT EXISTS business_relation;

CREATE TABLE business_relation.contact_groups (
  "id" text NOT NULL,
  "company_id" text NOT NULL,
  "image_url" text NOT NULL DEFAULT '',
  "name" text NOT NULL,
  "description" text NOT NULL DEFAULT '',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE business_relation.contact_books (
  "id" text NOT NULL,
  "konekin_id" text NOT NULL DEFAULT '',
  "primary_company_id" text NOT NULL,
  "secondary_company_id" text NOT NULL DEFAULT '',
  "contact_group_id" text NOT NULL DEFAULT '',
  "name" text NOT NULL,
  "email" text NOT NULL DEFAULT '',
  "phone" text NOT NULL DEFAULT '',
  "mobile" text NOT NULL DEFAULT '',
  "web" text NOT NULL DEFAULT '',
  "is_all_branches" bool NOT NULL DEFAULT false,
  "is_customer" bool NOT NULL DEFAULT false,
  "is_supplier" bool NOT NULL DEFAULT false,
  "is_tax" bool NOT NULL DEFAULT false,
  "tax_id" text NOT NULL DEFAULT '',
  "is_default" bool NOT NULL DEFAULT false,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE business_relation.contact_book_additional_infos (
  "contact_book_id" text NOT NULL,
  "nickname" text NOT NULL DEFAULT '',
  "tag" text NOT NULL DEFAULT '',
  "note" text NOT NULL DEFAULT '',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE business_relation.contact_book_mailing_addresses (
  "contact_book_id" text NOT NULL,
  "province" text NOT NULL DEFAULT '',
  "regency" text NOT NULL DEFAULT '',
  "district" text NOT NULL DEFAULT '',
  "postal_code" text NOT NULL DEFAULT '',
  "full_address" text NOT NULL DEFAULT '',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE business_relation.contact_book_shipping_addresses (
  "contact_book_id" text NOT NULL,
  "province" text NOT NULL DEFAULT '',
  "regency" text NOT NULL DEFAULT '',
  "district" text NOT NULL DEFAULT '',
  "postal_code" text NOT NULL DEFAULT '',
  "full_address" text NOT NULL DEFAULT '',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE business_relation.contact_book_branches (
  "contact_book_id" text NOT NULL,
  "company_branch_id" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE business_relation.contact_book_customer_infos (
  "contact_book_id" text NOT NULL,
  "pic" text NOT NULL DEFAULT '',
  "credit_limit" bigint NOT NULL DEFAULT 0,
  "payment_term" int NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("contact_book_id")
);

CREATE TABLE business_relation.contact_book_supplier_infos (
  "contact_book_id" text NOT NULL,
  "pic" text NOT NULL DEFAULT '',
  "credit_limit" bigint NOT NULL DEFAULT 0,
  "payment_term" int NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("contact_book_id")
);

CREATE TABLE business_relation.contact_invitations (
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