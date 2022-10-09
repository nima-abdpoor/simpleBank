-- name: CreateEntry :one
INSERT INTO entries (account_id,
                      amount)
values ($1, $2) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: GetAllEntries :many
SELECT * FROM entries
ORDER BY id;

-- name: ListEntries :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY id
    LIMIT $2
OFFSET $3;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;

-- name: UpdateEntry :one
UPDATE entries
set amount = $3
WHERE id = $1 and account_id = $2
    RETURNING *;