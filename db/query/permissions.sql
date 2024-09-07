-- name: CreatePermission :one
INSERT INTO "Permissions" ( "Name", "Description", "CreatedAt" )
VALUES ( $1, $2, $3 )
    RETURNING *;

-- name: GetPermission :one
SELECT *
FROM "Permissions"
WHERE "Id" = $1
    LIMIT 1;

-- name: GetPermissionForUpdate :one
SELECT *
FROM "Permissions"
WHERE "Id" = $1
    LIMIT 1
FOR UPDATE;

-- name: ListPermissions :many
SELECT *
FROM "Permissions"
ORDER BY "CreatedAt" DESC
    LIMIT $1 OFFSET $2;

-- name: UpdatePermission :one
UPDATE "Permissions"
SET "Name" = $2, "Description" = $3, "CreatedAt" = $4
WHERE "Id" = $1
    RETURNING *;

-- name: DeletePermission :exec
DELETE
FROM "Permissions"
WHERE "Id" = $1;
