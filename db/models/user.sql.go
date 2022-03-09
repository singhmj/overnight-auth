// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package models

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(id, login_id, password, status, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)
RETURNING id, login_id, password, status, created_at, updated_at
`

type CreateUserParams struct {
	ID        string      `json:"id"`
	LoginID   string      `json:"login_id"`
	Password  string      `json:"password"`
	Status    string      `json:"status"`
	CreatedAt interface{} `json:"created_at"`
	UpdatedAt interface{} `json:"updated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.ID,
		arg.LoginID,
		arg.Password,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.LoginID,
		&i.Password,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, login_id, password, status, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.LoginID,
		&i.Password,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, login_id, password, status, created_at, updated_at FROM users LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.LoginID,
			&i.Password,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users
SET password = $1, updated_at = $2
WHERE id = $3
`

type UpdateUserPasswordParams struct {
	Password  string      `json:"password"`
	UpdatedAt interface{} `json:"updated_at"`
	ID        string      `json:"id"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.exec(ctx, q.updateUserPasswordStmt, updateUserPassword, arg.Password, arg.UpdatedAt, arg.ID)
	return err
}

const updateUserStatus = `-- name: UpdateUserStatus :exec
UPDATE users
SET status = $1, updated_at = $2
WHERE id = $3
`

type UpdateUserStatusParams struct {
	Status    string      `json:"status"`
	UpdatedAt interface{} `json:"updated_at"`
	ID        string      `json:"id"`
}

func (q *Queries) UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) error {
	_, err := q.exec(ctx, q.updateUserStatusStmt, updateUserStatus, arg.Status, arg.UpdatedAt, arg.ID)
	return err
}
