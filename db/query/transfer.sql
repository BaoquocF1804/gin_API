-- name: CreateTransfer :exec
INSERT INTO transfer (
    from_account_id,
    to_account_id,
    amount
) VALUES (
    ?, ?, ?
) ;

-- name: GetTransfer :one
SELECT * FROM transfer
WHERE id = ? LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfer
WHERE
    from_account_id = ? OR
    to_account_id = ?
ORDER BY id
    LIMIT ?
OFFSET ?;