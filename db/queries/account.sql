-- name: GetAccountByID :one
SELECT id, name FROM accounts
WHERE id = $1 LIMIT 1;