-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id,
                       to_account_id,
                       amount)
values ($1, $2, $3) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetAllTransfers :many
SELECT * FROM transfers
ORDER BY id;

-- name: ListTransfers :many
SELECT *
FROM transfers
WHERE from_account_id = $1 OR
      to_account_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;

-- name: UpdateTransfer :one
UPDATE transfers
set amount = $2
WHERE id = $1
    RETURNING *;