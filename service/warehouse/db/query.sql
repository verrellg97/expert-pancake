-- name: UpsertRack :one
INSERT INTO warehouse.racks(id, branch_id, name)
VALUES ($1, $2, $3)
ON CONFLICT (id)
DO UPDATE SET
    name = EXCLUDED.name,
    branch_id = EXCLUDED.branch_id,
    updated_at = NOW()
RETURNING *;
