-- name: InsertContactGroup :one
INSERT INTO business_relation.contact_groups(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateContactGroup :one
UPDATE business_relation.contact_groups
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;