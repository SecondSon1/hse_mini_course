-- name: GetAccount :one
SELECT * FROM accounts
WHERE name = $1 LIMIT 1;

-- name: CreateAccount :one
INSERT INTO accounts (name, balance)
VALUES ($1, $2) RETURNING *;

-- name: UpdateBalance :one
UPDATE accounts SET
  balance = balance + $2
WHERE name=$1 RETURNING *;

-- name: UpdateName :one
UPDATE accounts SET
  name = $2
WHERE name=$1 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE name=$1;
