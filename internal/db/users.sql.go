// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkUserExists = `-- name: CheckUserExists :one
SELECT EXISTS (SELECT 1 FROM users WHERE email = $1) AS exists
`

func (q *Queries) CheckUserExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRow(ctx, checkUserExists, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password_hash, email, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id
`

type CreateUserParams struct {
	Username     string             `json:"username"`
	PasswordHash string             `json:"password_hash"`
	Email        string             `json:"email"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.PasswordHash,
		arg.Email,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email, password_hash, created_at, updated_at, active FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, password_hash, created_at, updated_at, active FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
	)
	return i, err
}