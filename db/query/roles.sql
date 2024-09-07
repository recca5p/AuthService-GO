-- name: CreateRole :one
INSERT INTO "Roles" ( "Name", "CreatedAt" )
VALUES ( $1, $2 )
    RETURNING *;

-- name: GetRole :one
SELECT *
FROM "Roles"
WHERE "Id" = $1
    LIMIT 1;

-- name: GetRoleForUpdate :one
SELECT *
FROM "Roles"
WHERE "Id" = $1
    LIMIT 1
FOR UPDATE;

-- name: ListRoles :many
SELECT *
FROM "Roles"
ORDER BY "CreatedAt" DESC
    LIMIT $1 OFFSET $2;

-- name: UpdateRole :one
UPDATE "Roles"
SET "Name" = $2, "CreatedAt" = $3
WHERE "Id" = $1
    RETURNING *;

-- name: DeleteRole :exec
DELETE
FROM "Roles"
WHERE "Id" = $1;
