package sqlc

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (name, dob)
VALUES (?, ?)
`

type CreateUserParams struct {
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.Name, arg.Dob)
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, dob FROM users WHERE id = ?
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Dob)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, dob FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Dob); err != nil {
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

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET name = ?, dob = ?
WHERE id = ?
`

type UpdateUserParams struct {
	Name string `json:"name"`
	Dob  string `json:"dob"`
	ID   int32  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.Name, arg.Dob, arg.ID)
	return err
}
