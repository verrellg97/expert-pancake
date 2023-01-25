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
