-- name: CreateUser :one
INSERT INTO users(id, login_id, password, status, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users LIMIT $1 OFFSET $2;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $1, updated_at = $2
WHERE id = $3;

-- name: UpdateUserStatus :exec
UPDATE users
SET status = $1, updated_at = $2
WHERE id = $3;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;