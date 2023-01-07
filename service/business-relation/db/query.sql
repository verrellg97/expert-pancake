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

-- name: UpdateContactBook :one
UPDATE business_relation.contact_books
SET 
    contact_group_id = $2,
    name = $3,
    email = $4,
    phone = $5,
    mobile = $6,
    web = $7,
    is_all_branches = $8,
    is_customer = $9,
    is_supplier = $10,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetContactBooks :many
SELECT a.id, a.primary_company_id, a.secondary_company_id, 
a.contact_group_id, a.name, a.email, a.phone, a.mobile, a.web,
a.is_all_branches, a.is_customer, a.is_supplier,
b.nickname, b.tag, b.note,
c.province AS mailing_province, c.regency AS mailing_regency,
c.district AS mailing_district, c.postal_code AS mailing_postal_code,
c.full_address AS mailing_full_address,
d.province AS shipping_province, d.regency AS shipping_regency,
d.district AS shipping_district, d.postal_code AS shipping_postal_code,
d.full_address AS shipping_full_address
FROM business_relation.contact_books a
JOIN business_relation.contact_book_additional_infos b ON a.id = b.contact_book_id
JOIN business_relation.contact_book_mailing_addresses c ON a.id = c.contact_book_id
JOIN business_relation.contact_book_shipping_addresses d ON a.id = d.contact_book_id
WHERE a.primary_company_id = $1;

-- name: InsertContactBookAdditionalInfo :exec
INSERT INTO business_relation.contact_book_additional_infos(contact_book_id,
nickname, tag, note)
VALUES ($1, $2, $3, $4);

-- name: UpdateContactBookAdditionalInfo :exec
UPDATE business_relation.contact_book_additional_infos
SET 
    nickname = $2,
    tag = $3,
    note = $4,
    updated_at = NOW()
WHERE contact_book_id = $1;

-- name: InsertContactBookMailingAddress :exec
INSERT INTO business_relation.contact_book_mailing_addresses(contact_book_id,
province, regency, district, postal_code, full_address)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateContactBookMailingAddress :exec
UPDATE business_relation.contact_book_mailing_addresses
SET 
    province = $2,
    regency = $3,
    district = $4,
    postal_code = $5,
    full_address = $6,
    updated_at = NOW()
WHERE contact_book_id = $1;

-- name: InsertContactBookShippingAddress :exec
INSERT INTO business_relation.contact_book_shipping_addresses(contact_book_id,
province, regency, district, postal_code, full_address)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateContactBookShippingAddress :exec
UPDATE business_relation.contact_book_shipping_addresses
SET 
    province = $2,
    regency = $3,
    district = $4,
    postal_code = $5,
    full_address = $6,
    updated_at = NOW()
WHERE contact_book_id = $1;

-- name: InsertContactBookBranch :exec
INSERT INTO business_relation.contact_book_branches(contact_book_id, company_branch_id)
VALUES ($1, $2);

-- name: DeleteContactBookBranches :exec
DELETE FROM business_relation.contact_book_branches
WHERE contact_book_id = $1;

-- name: GetContactBookBranches :many
SELECT contact_book_id, company_branch_id FROM business_relation.contact_book_branches
WHERE contact_book_id = $1;