-- name: InsertNotification :one
INSERT INTO notification.notifications(id, company_id, branch_id, title, content, type)
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: GetNotifications :many
SELECT id, company_id, branch_id, title, content, type, created_at
FROM notification.notifications
WHERE 
    company_id = $1
    AND branch_id LIKE $2
    AND CASE WHEN @is_read_filter::bool
    THEN is_read = $3 ELSE TRUE END
    AND is_deleted = false
ORDER BY created_at DESC;

-- name: ReadNotification :exec
UPDATE notification.notifications
SET is_read = true
WHERE id = $1;

-- name: DeleteNotification :exec
UPDATE notification.notifications
SET is_deleted = true
WHERE id = $1;
