-- name: CreateAccount :one
INSERT INTO "Accounts" ("Username", "PasswordHash", "CreatedAt", "UpdatedAt")
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetAccount :one
SELECT *
FROM "Accounts"
WHERE "Id" = $1 LIMIT 1;

-- name: UpdateAccount :one
UPDATE "Accounts"
SET "PasswordHash" = $2, "UpdatedAt" = $3
WHERE "Id" = $1 RETURNING *;

-- name: DeleteAccount :exec
DELETE
FROM "Accounts"
WHERE "Id" = $1;

-- name: ListAccounts :many
SELECT *
FROM "Accounts"
WHERE ("Username" = $1 OR $1 IS NULL OR $1 = '')
ORDER BY "UpdatedAt" DESC LIMIT $2 OFFSET $3;
