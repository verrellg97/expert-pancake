-- name: InsertBrand :one
INSERT INTO inventory.brands(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateBrand :one
UPDATE inventory.brands
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetBrands :many
SELECT id, company_id, name FROM inventory.brands
WHERE company_id = $1 AND name LIKE $2;

-- name: InsertGroup :one
INSERT INTO inventory.groups(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateGroup :one
UPDATE inventory.groups
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetGroups :many
SELECT id, company_id, name FROM inventory.groups
WHERE company_id = $1 AND name LIKE $2;

-- name: InsertUnit :one
INSERT INTO inventory.units(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUnit :one
UPDATE inventory.units
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetUnits :many
SELECT id, company_id, name FROM inventory.units
WHERE company_id = $1 AND name LIKE $2;
