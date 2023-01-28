-- name: UpsertRack :one
INSERT INTO warehouse.racks(id, branch_id, name)
VALUES ($1, $2, $3)
ON CONFLICT (id)
DO UPDATE SET
    name = EXCLUDED.name,
    branch_id = EXCLUDED.branch_id,
    updated_at = NOW()
RETURNING *;

-- name: GetRacks :many
SELECT a.id, a.branch_id, a.name
FROM warehouse.racks a
LEFT JOIN warehouse.warehouse_racks b ON a.id = b.rack_id
WHERE a.branch_id = $1 
    AND a.name LIKE $2 
    AND CASE WHEN
        @is_get_all::bool THEN TRUE ELSE b.rack_id IS NULL END;
