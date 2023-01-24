-- name: InsertItemBrand :one
INSERT INTO inventory.item_brands(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;
