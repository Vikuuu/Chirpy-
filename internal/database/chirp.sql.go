// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: chirp.sql

package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

const createChirpForUser = `-- name: CreateChirpForUser :one
INSERT INTO chirp (id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(), NOW(), NOW(), $1, $2
)
RETURNING id, created_at, updated_at, body, user_id
`

type CreateChirpForUserParams struct {
	Body   string
	UserID uuid.UUID
}

func (q *Queries) CreateChirpForUser(ctx context.Context, arg CreateChirpForUserParams) (Chirp, error) {
	row := q.db.QueryRowContext(ctx, createChirpForUser, arg.Body, arg.UserID)
	var i Chirp
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.UserID,
	)
	return i, err
}

const deleteChirp = `-- name: DeleteChirp :exec
DELETE FROM chirp
WHERE user_id = $1 AND id = $2
`

type DeleteChirpParams struct {
	UserID uuid.UUID
	ID     uuid.UUID
}

func (q *Queries) DeleteChirp(ctx context.Context, arg DeleteChirpParams) error {
	_, err := q.db.ExecContext(ctx, deleteChirp, arg.UserID, arg.ID)
	return err
}

const getChirp = `-- name: GetChirp :one
SELECT id, created_at, updated_at, body, user_id
FROM chirp
WHERE id = $1
`

func (q *Queries) GetChirp(ctx context.Context, id uuid.UUID) (Chirp, error) {
	row := q.db.QueryRowContext(ctx, getChirp, id)
	var i Chirp
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.UserID,
	)
	return i, err
}

const getChirps = `-- name: GetChirps :many
SELECT id, created_at, updated_at, body, user_id
FROM chirp
ORDER BY created_at ASC
`

func (q *Queries) GetChirps(ctx context.Context) ([]Chirp, error) {
	rows, err := q.db.QueryContext(ctx, getChirps)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chirp
	for rows.Next() {
		var i Chirp
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
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

const getChirpsForAuthor = `-- name: GetChirpsForAuthor :many
SELECT id, created_at, updated_at, body, user_id
FROM chirp
WHERE user_id = $1
ORDER BY created_at ASC
`

func (q *Queries) GetChirpsForAuthor(ctx context.Context, userID uuid.UUID) ([]Chirp, error) {
	rows, err := q.db.QueryContext(ctx, getChirpsForAuthor, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chirp
	for rows.Next() {
		var i Chirp
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
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

const getSortedChirps = `-- name: GetSortedChirps :many
SELECT id, created_at, updated_at, body, user_id
FROM chirp
ORDER BY created_at %s
`

func (q *Queries) GetSortedChirps(ctx context.Context, sort string) ([]Chirp, error) {
	sortDirection := "ASC"
	if sort == "desc" {
		sortDirection = "DESC"
	}

	query := fmt.Sprintf(getSortedChirps, sortDirection)

	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chirp
	for rows.Next() {
		var i Chirp
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
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

const getSortedChirpsForAuthor = `-- name: GetSortedChirpsForAuthor :many
SELECT id, created_at, updated_at, body, user_id
FROM chirp
WHERE user_id = $1
ORDER BY created_at %s
`

type GetSortedChirpsForAuthorParams struct {
	UserID uuid.UUID
	Sort   string
}

func (q *Queries) GetSortedChirpsForAuthor(ctx context.Context, arg GetSortedChirpsForAuthorParams) ([]Chirp, error) {
	sortDirection := "ASC"
	if arg.Sort == "desc" {
		sortDirection = "DESC"
	}

	query := fmt.Sprintf(getSortedChirpsForAuthor, sortDirection)
	rows, err := q.db.QueryContext(ctx, query, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chirp
	for rows.Next() {
		var i Chirp
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
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
