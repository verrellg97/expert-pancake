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

-- name: InsertItem :one
INSERT INTO inventory.items(id, company_id, image_url,
code, name, brand_id, group_id, tag, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: InsertItemVariant :one
INSERT INTO inventory.item_variants(id, item_id, image_url,
name, price, stock, is_default)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetBrandById :one
SELECT id, company_id, name FROM inventory.brands
WHERE id = $1;

-- name: GetGroupById :one
SELECT id, company_id, name FROM inventory.groups
WHERE id = $1;