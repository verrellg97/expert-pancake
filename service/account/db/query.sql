-- name: CreateUser :exec
INSERT INTO account.users (id, fullname, nickname, email, phone_number)
VALUES ($1, $2, $3, $4, $5);

-- name: UpsertUserInfo :exec
INSERT INTO account.user_infos(user_id, key, value)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, key)
DO UPDATE SET
    value = EXCLUDED.value,
    updated_at = NOW();

-- name: UpsertUserPassword :exec
INSERT INTO account.user_passwords(user_id, password)
VALUES ($1, $2)
ON CONFLICT (user_id)
DO UPDATE SET
    password = EXCLUDED.password,
    updated_at = NOW();

-- name: GetUserPassword :one
SELECT * FROM account.user_passwords
WHERE user_id = $1;

-- name: GetUser :one
SELECT * FROM account.users
WHERE id = $1;

-- name: GetUserByPhoneNumber :one
SELECT id FROM account.users
WHERE phone_number = $1;

-- name: UpsertUserAddresses :exec
INSERT INTO account.user_addresses(user_id, country, province, regency, district, full_address)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (user_id)
DO UPDATE SET
    country = EXCLUDED.country,
    province = EXCLUDED.province,
    regency = EXCLUDED.regency,
    district = EXCLUDED.district,
    full_address = EXCLUDED.full_address,
    updated_at = NOW();

-- name: UpdateUser :exec
UPDATE account.users
SET
    fullname = $2,
    nickname = $3,
    email = $4,
    phone_number = $5,
    updated_at = NOW()
WHERE
    id = $1;
