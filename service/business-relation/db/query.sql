-- name: InsertContactGroup :one
INSERT INTO business_relation.contact_groups(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateContactGroup :one
UPDATE business_relation.contact_groups
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetContactGroups :many
SELECT id, company_id, name
FROM business_relation.contact_groups
WHERE company_id = $1;

-- name: InsertContactBook :one
INSERT INTO business_relation.contact_books(id, primary_company_id, secondary_company_id,
contact_group_id, name, email, phone, mobile, web,
is_all_branches, is_customer, is_supplier)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,
$10, $11, $12)
RETURNING *;

-- name: InsertContactBookAdditionalInfo :exec
INSERT INTO business_relation.contact_book_additional_infos(contact_book_id,
nickname, tag, note)
VALUES ($1, $2, $3, $4);

-- name: InsertContactBookMailingAddress :exec
INSERT INTO business_relation.contact_book_mailing_addresses(contact_book_id,
province, regency, district, postal_code, full_address)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: InsertContactBookShippingAddress :exec
INSERT INTO business_relation.contact_book_shipping_addresses(contact_book_id,
province, regency, district, postal_code, full_address)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: InsertContactBookBranch :exec
INSERT INTO business_relation.contact_book_branches(contact_book_id, company_branch_id)
VALUES ($1, $2);