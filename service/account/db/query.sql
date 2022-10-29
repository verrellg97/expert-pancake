-- name: CreateUser :exec
INSERT INTO account.users (id, fullname, nickname, email, phone_number)
VALUES ($1, $2, $3, $4, $5);

-- name: UpsertUserInfo :exec
INSERT INTO account.user_infos(user_id, key, value)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, key)
DO UPDATE SET value = EXCLUDED.value;

-- name: UpsertUserPassword :exec
INSERT INTO account.user_passwords(user_id, password)
VALUES ($1, $2)
ON CONFLICT (user_id)
DO UPDATE SET password = EXCLUDED.password;