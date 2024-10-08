// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: permissions.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPermission = `-- name: CreatePermission :one
INSERT INTO "Permissions" ( "Name", "Description", "CreatedAt" )
VALUES ( $1, $2, $3 )
    RETURNING "Id", "Name", "Description", "CreatedAt"
`

type CreatePermissionParams struct {
	Name        string         `json:"Name"`
	Description sql.NullString `json:"Description"`
	CreatedAt   time.Time      `json:"CreatedAt"`
}

func (q *Queries) CreatePermission(ctx context.Context, arg CreatePermissionParams) (Permission, error) {
	row := q.db.QueryRowContext(ctx, createPermission, arg.Name, arg.Description, arg.CreatedAt)
	var i Permission
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const deletePermission = `-- name: DeletePermission :exec
DELETE
FROM "Permissions"
WHERE "Id" = $1
`

func (q *Queries) DeletePermission(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePermission, id)
	return err
}

const getPermission = `-- name: GetPermission :one
SELECT "Id", "Name", "Description", "CreatedAt"
FROM "Permissions"
WHERE "Id" = $1
    LIMIT 1
`

func (q *Queries) GetPermission(ctx context.Context, id uuid.UUID) (Permission, error) {
	row := q.db.QueryRowContext(ctx, getPermission, id)
	var i Permission
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getPermissionForUpdate = `-- name: GetPermissionForUpdate :one
SELECT "Id", "Name", "Description", "CreatedAt"
FROM "Permissions"
WHERE "Id" = $1
    LIMIT 1
FOR UPDATE
`

func (q *Queries) GetPermissionForUpdate(ctx context.Context, id uuid.UUID) (Permission, error) {
	row := q.db.QueryRowContext(ctx, getPermissionForUpdate, id)
	var i Permission
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const listPermissions = `-- name: ListPermissions :many
SELECT "Id", "Name", "Description", "CreatedAt"
FROM "Permissions"
ORDER BY "CreatedAt" DESC
    LIMIT $1 OFFSET $2
`

type ListPermissionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPermissions(ctx context.Context, arg ListPermissionsParams) ([]Permission, error) {
	rows, err := q.db.QueryContext(ctx, listPermissions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Permission{}
	for rows.Next() {
		var i Permission
		if err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
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

const updatePermission = `-- name: UpdatePermission :one
UPDATE "Permissions"
SET "Name" = $2, "Description" = $3, "CreatedAt" = $4
WHERE "Id" = $1
    RETURNING "Id", "Name", "Description", "CreatedAt"
`

type UpdatePermissionParams struct {
	Id          uuid.UUID      `json:"Id"`
	Name        string         `json:"Name"`
	Description sql.NullString `json:"Description"`
	CreatedAt   time.Time      `json:"CreatedAt"`
}

func (q *Queries) UpdatePermission(ctx context.Context, arg UpdatePermissionParams) (Permission, error) {
	row := q.db.QueryRowContext(ctx, updatePermission,
		arg.Id,
		arg.Name,
		arg.Description,
		arg.CreatedAt,
	)
	var i Permission
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}
