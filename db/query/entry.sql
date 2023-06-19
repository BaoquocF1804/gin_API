-- name: CreateEntry :exec
INSERT INTO entries (
    accounts_id,
    amount
) VALUES (
    ?, ?
) ;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = ? LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE accounts_id = ?
ORDER BY id
    LIMIT ?
OFFSET ?;