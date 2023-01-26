-- name: InsertItemBrand :one
INSERT INTO inventory.item_brands(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateItemBrand :one
UPDATE inventory.item_brands
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetItemBrands :many
SELECT id, company_id, name FROM inventory.item_brands
WHERE company_id = $1 AND name LIKE $2;

-- name: InsertItemGroup :one
INSERT INTO inventory.item_groups(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateItemGroup :one
UPDATE inventory.item_groups
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
