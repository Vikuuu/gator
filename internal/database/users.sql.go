// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id, created_at, updated_at, name
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const deleteUserData = `-- name: DeleteUserData :exec
DELETE FROM users
`

func (q *Queries) DeleteUserData(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteUserData)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, created_at, updated_at FROM users WHERE name = $1
`

type GetUserRow struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) GetUser(ctx context.Context, name string) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, name)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserID = `-- name: GetUserID :one
SELECT id FROM users WHERE name = $1
`

func (q *Queries) GetUserID(ctx context.Context, name string) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getUserID, name)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getUserNameFromID = `-- name: GetUserNameFromID :one
SELECT name FROM users WHERE id = $1
`

func (q *Queries) GetUserNameFromID(ctx context.Context, id uuid.UUID) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserNameFromID, id)
	var name string
	err := row.Scan(&name)
	return name, err
}

const getUsers = `-- name: GetUsers :many
SELECT name FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
