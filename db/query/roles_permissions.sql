-- name: CreateRolePermission :one
INSERT INTO "RolesPermissions" ("RoleId", "PermissionId", "CreatedAt")
VALUES ($1, $2, $3)
    RETURNING *;

-- name: GetRolePermission :one
SELECT *
FROM "RolesPermissions"
WHERE "Id" = $1
    LIMIT 1;

-- name: GetRolePermissionForUpdate :one
SELECT *
FROM "RolesPermissions"
WHERE "Id" = $1
    LIMIT 1
FOR UPDATE;

-- name: ListRolePermissions :many
SELECT *
FROM "RolesPermissions"
WHERE ("RoleId" = $1 OR $1 IS NULL)
ORDER BY "CreatedAt" DESC
    LIMIT $2 OFFSET $3;

-- name: DeleteRolePermission :exec
DELETE
FROM "RolesPermissions"
WHERE "Id" = $1;
