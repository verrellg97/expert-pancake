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
WHERE CASE WHEN @is_filter_id::bool THEN id = $1
ELSE branch_id = $2 AND name LIKE $3
AND is_deleted = false END;

-- name: UpsertWarehouseRack :one
INSERT INTO warehouse.warehouse_racks(id, warehouse_id, name)
VALUES ($1, $2, $3)
ON CONFLICT (id)
DO UPDATE SET
    warehouse_id = EXCLUDED.warehouse_id,
    name = EXCLUDED.name,
    updated_at = NOW()
RETURNING *;

-- name: GetWarehouseRacks :many
SELECT id, warehouse_id, name
FROM warehouse.warehouse_racks
WHERE CASE WHEN @is_filter_id::bool THEN id = $1
ELSE warehouse_id = $2 AND name LIKE $3 AND is_deleted = false END;

-- name: DeleteWarehouse :exec
UPDATE warehouse.warehouses
SET is_deleted = true
WHERE id = $1;

-- name: DeleteWarehouseRack :exec
UPDATE warehouse.warehouse_racks
SET is_deleted = true
WHERE id = $1;