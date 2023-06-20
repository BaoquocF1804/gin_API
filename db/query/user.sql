-- name: CreateUser :exec
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    ?, ?, ?, ?
    ) ;

-- name: GetUser :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
SET
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    email = COALESCE(sqlc.narg(email), email),
    is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
    username = sqlc.arg(username)
;