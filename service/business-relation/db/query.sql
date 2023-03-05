-- name: InsertContactGroup :one
INSERT INTO business_relation.contact_groups(id, company_id, image_url, name, description)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateContactGroup :one
UPDATE business_relation.contact_groups
SET 
    image_url = $2,
    name = $3,
    description = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetContactGroups :many
SELECT a.id, a.company_id, a.image_url, a.name, a.description, COUNT(b.id) AS member
FROM business_relation.contact_groups a 
LEFT JOIN business_relation.contact_books b ON a.id = b.contact_group_id
WHERE a.company_id = $1
GROUP BY a.id;

-- name: InsertContactBook :one
INSERT INTO business_relation.contact_books(id, konekin_id, primary_company_id, secondary_company_id,
contact_group_id, name, email, phone, mobile, web,
is_all_branches, is_customer, is_supplier, is_default)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,
$10, $11, $12, $13, $14)
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

-- name: UpdateContactBookTaxInfo :exec
UPDATE business_relation.contact_books
SET 
    is_tax = $2,
    tax_id = $3,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateContactBookGroupId :exec
UPDATE business_relation.contact_books
SET 
    contact_group_id = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateContactBookGroupIdByGroupId :exec
UPDATE business_relation.contact_books
SET 
    contact_group_id = @new_contact_group_id::text,
    updated_at = NOW()
WHERE contact_group_id = $1;

-- name: UpdateContactBookRelation :exec
UPDATE business_relation.contact_books
SET 
    konekin_id = $2,
    secondary_company_id = $3,
    updated_at = NOW()
WHERE id = $1;

-- name: GetContactBooks :many
SELECT a.id, a.konekin_id, a.primary_company_id, a.secondary_company_id, 
a.contact_group_id, COALESCE(e.name, '') AS contact_group_name,
a.name, a.email, a.phone, a.mobile, a.web,
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
LEFT JOIN business_relation.contact_groups e ON a.contact_group_id = e.id
WHERE a.primary_company_id = $1
AND a.konekin_id = ''
AND a.is_default = FALSE
AND CASE WHEN @is_filter_group_id::bool
THEN a.contact_group_id = $2 ELSE TRUE END
AND CASE WHEN @is_customer_applicant::bool
THEN a.is_customer = FALSE
WHEN @is_supplier_applicant::bool
THEN a.is_supplier = FALSE ELSE TRUE END
UNION ALL
SELECT a.id, a.konekin_id, a.primary_company_id, a.secondary_company_id, 
a.contact_group_id, COALESCE(f.name, '') AS contact_group_name,
b.name, b.email, b.phone, b.mobile, b.web,
a.is_all_branches, a.is_customer, a.is_supplier,
c.nickname, c.tag, c.note,
d.province AS mailing_province, d.regency AS mailing_regency,
d.district AS mailing_district, d.postal_code AS mailing_postal_code,
d.full_address AS mailing_full_address,
e.province AS shipping_province, e.regency AS shipping_regency,
e.district AS shipping_district, e.postal_code AS shipping_postal_code,
e.full_address AS shipping_full_address
FROM business_relation.contact_books a
JOIN business_relation.contact_books b ON a.konekin_id = b.konekin_id AND b.is_default = TRUE
JOIN business_relation.contact_book_additional_infos c ON b.id = c.contact_book_id
JOIN business_relation.contact_book_mailing_addresses d ON b.id = d.contact_book_id
JOIN business_relation.contact_book_shipping_addresses e ON b.id = e.contact_book_id
LEFT JOIN business_relation.contact_groups f ON a.contact_group_id = f.id
WHERE a.primary_company_id = $1
AND a.konekin_id <> ''
AND a.is_default = FALSE
AND CASE WHEN @is_filter_group_id::bool
THEN a.contact_group_id = $2 ELSE TRUE END
AND CASE WHEN @is_customer_applicant::bool
THEN a.is_customer = FALSE
WHEN @is_supplier_applicant::bool
THEN a.is_supplier = FALSE ELSE TRUE END;

-- name: GetContactBookById :one
SELECT a.id, a.konekin_id, a.primary_company_id, a.secondary_company_id,
a.contact_group_id, COALESCE(b.name, '') AS contact_group_name, a.name, a.email,
a.phone, a.mobile, a.web, a.is_all_branches, a.is_customer, a.is_supplier,
a.is_tax, a.tax_id, a.is_deleted
FROM business_relation.contact_books a
LEFT JOIN business_relation.contact_groups b ON a.contact_group_id = b.id
WHERE a.id = $1;

-- name: GetMyContactBook :one
SELECT a.id, a.konekin_id, a.primary_company_id,
a.name, a.email, a.phone, a.mobile, a.web
FROM business_relation.contact_books a
WHERE a.primary_company_id = $1
AND a.is_customer = FALSE
AND a.is_supplier = FALSE
AND a.is_default = TRUE;

-- name: GetContactBookAdditionalInfo :one
SELECT a.nickname, a.tag, a.note
FROM business_relation.contact_book_additional_infos a
WHERE a.contact_book_id = $1;

-- name: GetContactBookMailingAddress :one
SELECT a.province, a.regency, a.district, a.postal_code, a.full_address
FROM business_relation.contact_book_mailing_addresses a
WHERE a.contact_book_id = $1;

-- name: GetContactBookShippingAddress :one
SELECT a.province, a.regency, a.district, a.postal_code, a.full_address
FROM business_relation.contact_book_shipping_addresses a
WHERE a.contact_book_id = $1;

-- name: GetCountKonekinId :one
SELECT COUNT(a.id)
FROM business_relation.contact_books a
WHERE a.konekin_id LIKE $1;

-- name: GetCustomers :many
SELECT a.id, a.konekin_id, a.primary_company_id, a.secondary_company_id, a.contact_group_id,
COALESCE(b.name, '') AS contact_group_name, a.name, a.email, a.phone,
a.mobile, a.web, a.is_all_branches, a.is_customer, a.is_supplier,
a.is_tax, a.tax_id, a.is_default, a.is_deleted, COALESCE(c.pic, '') AS pic,
COALESCE(c.credit_limit, 0) AS credit_limit, COALESCE(c.payment_term, 0) AS payment_term
FROM business_relation.contact_books a
LEFT JOIN business_relation.contact_groups b ON a.contact_group_id = b.id
LEFT JOIN business_relation.contact_book_customer_infos c ON a.id = c.contact_book_id
WHERE a.primary_company_id = $1 AND a.is_customer;

-- name: GetSuppliers :many
SELECT a.id, a.konekin_id, a.primary_company_id, a.secondary_company_id, a.contact_group_id,
COALESCE(b.name, '') AS contact_group_name, a.name, a.email, a.phone,
a.mobile, a.web, a.is_all_branches, a.is_customer, a.is_supplier,
a.is_tax, a.tax_id, a.is_default, a.is_deleted, COALESCE(c.pic, '') AS pic,
COALESCE(c.credit_limit, 0) AS credit_limit, COALESCE(c.payment_term, 0) AS payment_term
FROM business_relation.contact_books a
LEFT JOIN business_relation.contact_groups b ON a.contact_group_id = b.id
LEFT JOIN business_relation.contact_book_supplier_infos c ON a.id = c.contact_book_id
WHERE a.primary_company_id = $1 AND a.is_supplier;

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

-- name: UpsertCustomerInfo :exec
INSERT INTO business_relation.contact_book_customer_infos(contact_book_id, pic, credit_limit, payment_term)
VALUES ($1, $2, $3, $4)
ON CONFLICT (contact_book_id)
DO UPDATE SET
    pic = EXCLUDED.pic,
    credit_limit = EXCLUDED.credit_limit,
    payment_term = EXCLUDED.payment_term,
    updated_at = NOW();

-- name: UpsertSupplierInfo :exec
INSERT INTO business_relation.contact_book_supplier_infos(contact_book_id, pic, credit_limit, payment_term)
VALUES ($1, $2, $3, $4)
ON CONFLICT (contact_book_id)
DO UPDATE SET
    pic = EXCLUDED.pic,
    credit_limit = EXCLUDED.credit_limit,
    payment_term = EXCLUDED.payment_term,
    updated_at = NOW();

-- name: InsertContactInvitation :one
INSERT INTO business_relation.contact_invitations(id, primary_contact_book_id,
primary_company_id, secondary_company_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateContactInvitation :one
UPDATE business_relation.contact_invitations
SET 
    secondary_contact_book_id = $2,
    status = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetContactInvitations :many
SELECT a.id, a.konekin_id, a.primary_company_id,
a.name, a.email, a.phone, a.mobile, a.web,
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
LEFT JOIN business_relation.contact_invitations e
ON ((a.primary_company_id = e.primary_company_id AND e.secondary_company_id = @company_id::text)
OR (a.primary_company_id = e.secondary_company_id AND e.primary_company_id = @company_id::text))
AND (e.status = 'waiting' OR e.status = 'accepted') 
WHERE a.primary_company_id <> @company_id::text
AND a.is_default = TRUE
AND a.is_customer = FALSE
AND a.is_supplier = FALSE
AND e.id IS NULL;

-- name: GetRequestInvitations :many
SELECT a.id, a.status,
b.id AS contact_book_id, b.konekin_id, b.primary_company_id,
b.name, b.email, b.phone, b.mobile, b.web,
c.nickname, c.tag, c.note,
d.province AS mailing_province, d.regency AS mailing_regency,
d.district AS mailing_district, d.postal_code AS mailing_postal_code,
d.full_address AS mailing_full_address,
e.province AS shipping_province, e.regency AS shipping_regency,
e.district AS shipping_district, e.postal_code AS shipping_postal_code,
e.full_address AS shipping_full_address
FROM business_relation.contact_invitations a
JOIN business_relation.contact_books b ON a.secondary_company_id = b.primary_company_id
AND b.is_default = TRUE AND b.is_customer = FALSE AND b.is_supplier = FALSE
JOIN business_relation.contact_book_additional_infos c ON b.id = c.contact_book_id
JOIN business_relation.contact_book_mailing_addresses d ON b.id = d.contact_book_id
JOIN business_relation.contact_book_shipping_addresses e ON b.id = e.contact_book_id
WHERE a.primary_company_id = $1
AND a.status <> 'cancelled';

-- name: GetReceiveInvitations :many
SELECT a.id, a.status,
b.id AS contact_book_id, b.konekin_id, b.primary_company_id,
b.name, b.email, b.phone, b.mobile, b.web,
c.nickname, c.tag, c.note,
d.province AS mailing_province, d.regency AS mailing_regency,
d.district AS mailing_district, d.postal_code AS mailing_postal_code,
d.full_address AS mailing_full_address,
e.province AS shipping_province, e.regency AS shipping_regency,
e.district AS shipping_district, e.postal_code AS shipping_postal_code,
e.full_address AS shipping_full_address
FROM business_relation.contact_invitations a
JOIN business_relation.contact_books b ON a.primary_company_id = b.primary_company_id
AND b.is_default = TRUE AND b.is_customer = FALSE AND b.is_supplier = FALSE
JOIN business_relation.contact_book_additional_infos c ON b.id = c.contact_book_id
JOIN business_relation.contact_book_mailing_addresses d ON b.id = d.contact_book_id
JOIN business_relation.contact_book_shipping_addresses e ON b.id = e.contact_book_id
WHERE a.secondary_company_id = $1
AND a.status <> 'cancelled';

-- name: AddCustomer :exec
UPDATE business_relation.contact_books
SET 
    is_customer = TRUE,
    updated_at = NOW()
WHERE id = ANY(@contact_book_ids::text[]);

-- name: AddSupplier :exec
UPDATE business_relation.contact_books
SET 
    is_supplier = TRUE,
    updated_at = NOW()
WHERE id = ANY(@contact_book_ids::text[]);