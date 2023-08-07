-- name: InsertCompany :one
INSERT INTO business.companies(id, user_id, name, initial_name, type, responsible_person, image_url, is_deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateCompany :one
UPDATE business.companies
SET name = $2, 
initial_name = $3, 
type = $4, 
responsible_person = $5, 
image_url = $6, 
updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCompany :exec
UPDATE business.companies
SET is_deleted = true, 
updated_at = NOW()
WHERE id = $1;

-- name: GetUserCompanies :many
SELECT id, user_id, name, initial_name, type, responsible_person, image_url FROM business.companies
WHERE user_id = $1 AND is_deleted = false;

-- name: GetUserCompaniesFilteredByName :many
SELECT id, user_id, name, initial_name, type, responsible_person, image_url
FROM business.companies
WHERE user_id = @user_id::text AND is_deleted = false AND name LIKE @keyword::text
UNION
SELECT a.id, a.user_id, a.name, a.initial_name, a.type, a.responsible_person, a.image_url
FROM business.companies a
JOIN business.company_members b ON a.id = b.company_id AND b.is_deleted = false
WHERE b.user_id = @user_id::text AND a.name LIKE @keyword::text;

-- name: InsertCompanyBranch :one
INSERT INTO business.company_branches(id, user_id, company_id, name, address, phone_number, is_central, is_deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateCompanyBranch :one
UPDATE business.company_branches
SET name = $2, 
address = $3, 
phone_number = $4, 
updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCompanyBranch :exec
UPDATE business.company_branches
SET is_deleted = true, 
updated_at = NOW()
WHERE id = $1;

-- name: DeleteCompanyBranchesByCompanyId :exec
UPDATE business.company_branches
SET is_deleted = true, 
updated_at = NOW()
WHERE company_id = $1;

-- name: GetUserCompanyBranches :many
SELECT id, user_id, company_id, name, address, phone_number, is_central 
FROM business.company_branches
WHERE user_id = $1 AND company_id = $2 AND is_deleted = false;

-- name: GetUserCompanyBranchesFilteredByName :many
SELECT id, user_id, company_id, name, address, phone_number, is_central 
FROM business.company_branches
WHERE user_id = @user_id::text AND company_id = @company_id::text AND is_deleted = false AND name LIKE @keyword::text
UNION
SELECT a.id, a.user_id, a.company_id, a.name, a.address, a.phone_number, a.is_central 
FROM business.company_branches a
JOIN business.company_members b ON a.company_id = b.company_id AND b.is_deleted = false
WHERE b.user_id = @user_id::text AND a.company_id = @company_id::text AND a.is_deleted = false AND a.name LIKE @keyword::text;

-- name: GetCompanyByName :one
SELECT id FROM business.companies
WHERE name = $1;

-- name: GetCompanyById :one
SELECT id, user_id, name, initial_name, type, responsible_person
FROM business.companies
WHERE id = $1;

-- name: GetCompanyBranchesByCompany :many
SELECT id, user_id, company_id, name, address, phone_number, is_central 
FROM business.company_branches
WHERE company_id = $1;

-- name: GetCompanyBranches :many
SELECT id, user_id, company_id, name, address, phone_number, is_central 
FROM business.company_branches
WHERE company_id = $1
ORDER BY is_central DESC;

-- name: InsertMemberRequest :exec
INSERT INTO business.company_member_requests(
    id, user_id, company_id
)
VALUES ($1, $2, $3);

-- name: GetReceiveMemberRequests :many
SELECT id, user_id, company_id, status
FROM business.company_member_requests
WHERE company_id = $1;

-- name: InsertCompanyMember :exec
INSERT INTO business.company_members(
    id, user_id, company_id
)
VALUES ($1, $2, $3);

-- name: UpdateMemberRequest :one
UPDATE business.company_member_requests
SET status = $2,
updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetCompaniesFilteredByName :many
SELECT id, user_id, name, initial_name, type, responsible_person, image_url
FROM business.companies
WHERE is_deleted = false AND name LIKE $1;