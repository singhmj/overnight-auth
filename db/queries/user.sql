-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users LIMIT $1 OFFSET $2;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $1
WHERE id = $2;

-- name: UpdateUserStatus :exec
UPDATE users
SET status = $1
WHERE id = $2;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;