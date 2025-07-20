-- name: CreateCredential :exec
INSERT INTO credentials (
    credential_id,
    user_id,
    type,
    is_primary,
    is_verified,
    identifier,
    secret_data
)
VALUES (@credential_id, @user_id, @type, @is_primary, @is_verified, @identifier, @secret_data);

-- name: GetCredentialByID :one
SELECT * FROM credentials
WHERE credential_id = @credential_id;

-- name: GetUserCredentials :many
SELECT * FROM credentials
WHERE user_id = @user_id
ORDER BY is_primary DESC, created_at DESC;

-- name: GetUserCredentialByType :one
SELECT * FROM credentials
WHERE type = @type AND user_id = @user_id;

-- name: UpdateCredentialSecret :one
UPDATE credentials
SET secret_data = @secret_data,
    updated_at = NOW()
WHERE type = @type
RETURNING *;

-- name: VerifyCredential :exec
UPDATE credentials
SET is_verified = true,
    updated_at = NOW()
WHERE credential_id = @credential_id;

-- name: DeleteCredential :exec
DELETE FROM credentials
WHERE credential_id = @credential_id;
