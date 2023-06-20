-- name: CreateVerifyEmail :exec
INSERT INTO verify_emails (
    username,
    email,
    secret_code
) VALUES (
             ?, ?, ?
         );

-- name: UpdateVerifyEmail :exec
UPDATE verify_emails
SET
    is_used = TRUE
WHERE
        id = @id
  AND secret_code = @secret_code
  AND is_used = FALSE
  AND expired_at > now()
  ;