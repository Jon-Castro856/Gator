-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, user_name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_name = $1;


-- name: ResetTable :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT user_name FROM users;

-- name: GetUserID :one
SELECT id FROM users
WHERE user_name = $1;

-- name: GetUserName :one
SELECT user_name FROM users
WHERE id = $1;

-- name: GetUserInfo :one
SELECT * FROM users
WHERE user_name = $1;