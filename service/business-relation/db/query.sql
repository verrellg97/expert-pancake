-- name: InsertContactGroup :one
INSERT INTO business_relation.contact_groups(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;