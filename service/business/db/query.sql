-- name: InsertCompany :one
INSERT INTO business.companies(id, user_id, name, initial_name, type, responsible_person, is_deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateCompany :one
UPDATE business.companies
SET name = $2, 
initial_name = $3, 
type = $4, 
responsible_person = $5, 
updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCompany :exec
UPDATE business.companies
SET is_deleted = true, 
updated_at = NOW()
WHERE id = $1;

-- name: GetUserCompanies :many
SELECT id, user_id, name, initial_name, type, responsible_person FROM business.companies
WHERE user_id = $1 AND is_deleted = false;

-- name: GetUserCompaniesFilteredByName :many
SELECT id, user_id, name, initial_name, type, responsible_person FROM business.companies
WHERE user_id = $1 AND is_deleted = false AND name LIKE $2;

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
WHERE user_id = $1 AND company_id = $2 AND is_deleted = false AND name LIKE $3;