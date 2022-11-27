-- name: UpsertCompany :one
INSERT INTO business.companies(id, user_id, name, initial_name, type, responsible_person, is_deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (id)
DO UPDATE SET
    name = EXCLUDED.name,
    initial_name = EXCLUDED.initial_name,
    type = EXCLUDED.type,
    responsible_person = EXCLUDED.responsible_person, 
    is_deleted = EXCLUDED.is_deleted, 
    updated_at = NOW()
RETURNING *;

-- name: GetUserCompanies :many
SELECT id, user_id, name, initial_name, type, responsible_person FROM business.companies
WHERE user_id = $1 AND is_deleted = 0;

-- name: GetUserCompaniesFilteredByName :many
SELECT id, user_id, name, initial_name, type, responsible_person FROM business.companies
WHERE user_id = $1 AND is_deleted = 0 AND name LIKE $2;

-- name: InsertCompanyBranch :exec
INSERT INTO business.company_branches(id, user_id, company_id, name, address, phone_number, is_central, is_deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: UpsertCompanyBranch :one
INSERT INTO business.company_branches(id, user_id, company_id, name, address, phone_number, is_central, is_deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (id)
DO UPDATE SET
    name = EXCLUDED.name,
    address = EXCLUDED.address,
    phone_number = EXCLUDED.phone_number,
    is_deleted = EXCLUDED.is_deleted, 
    updated_at = NOW()
RETURNING *;

-- name: GetUserCompanyBranches :many
SELECT id, user_id, company_id, name, address, phone_number FROM business.company_branches
WHERE user_id = $1 AND company_id = $2 AND is_deleted = 0;

-- name: GetUserCompanyBranchesFilteredByName :many
SELECT id, user_id, company_id, name, address, phone_number FROM business.company_branches
WHERE user_id = $1 AND company_id = $2 AND is_deleted = 0 AND name LIKE $3;