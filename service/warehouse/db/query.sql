-- name: UpsertWarehouse :one
INSERT INTO warehouse.warehouses(id, branch_id, code, name, address, type)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (id)
DO UPDATE SET
    branch_id = EXCLUDED.branch_id,
    name = EXCLUDED.name,
    address = EXCLUDED.address,
    type = EXCLUDED.type,
    updated_at = NOW()
RETURNING *;

-- name: GetWarehouses :many
SELECT id, branch_id, code, name, address, type
FROM warehouse.warehouses
WHERE branch_id = $1 AND name LIKE $2
AND is_deleted = false;