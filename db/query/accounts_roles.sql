-- name: CreateAccountRole :one
INSERT INTO "AccountsRoles" ("AccountId", "RoleId", "CreatedAt")
VALUES ($1, $2, $3)
    RETURNING "Id", "AccountId", "RoleId", "CreatedAt";

-- name: GetAccountRole :one
SELECT "Id", "AccountId", "RoleId", "CreatedAt"
FROM "AccountsRoles"
WHERE "Id" = $1
    LIMIT 1;

-- name: GetAccountRoleForUpdate :one
SELECT "Id", "AccountId", "RoleId", "CreatedAt"
FROM "AccountsRoles"
WHERE "Id" = $1
    LIMIT 1
FOR UPDATE;

-- name: ListAccountRoles :many
SELECT "Id", "AccountId", "RoleId", "CreatedAt"
FROM "AccountsRoles"
WHERE ("AccountId" = $1 OR $1 IS NULL)
ORDER BY "CreatedAt" DESC
    LIMIT $2 OFFSET $3;

-- name: DeleteAccountRole :exec
DELETE
FROM "AccountsRoles"
WHERE "Id" = $1;
