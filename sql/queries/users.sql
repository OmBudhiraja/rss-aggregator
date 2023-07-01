-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;