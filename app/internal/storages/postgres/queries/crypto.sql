-- name: GetChains :many
SELECT * FROM chains
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CreateChain :exec
INSERT INTO chains (chain_id, name, created_at, updated_at) 
VALUES (@chain_id, @name, NOW(), NOW());

-- name: GetSupportedTokens :many
SELECT t.*
FROM tokens t
JOIN token_chains tc ON t.token_id = tc.token_id
WHERE 
  (@chain_id::uuid IS NULL OR tc.chain_id = @chain_id) AND
  (@is_native::boolean IS NULL OR t.is_native = @is_native) AND
  t.deleted_at IS NULL
ORDER BY t.created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetTokenByID :one
SELECT * FROM tokens
WHERE token_id = @token_id AND deleted_at IS NULL;

-- name: CreateToken :exec
INSERT INTO tokens (token_id, is_native, name, symbol, decimals, logo_path, created_at, updated_at) 
VALUES (@token_id, @is_native, @name, @symbol, @decimals, @logo_path, NOW(), NOW());

-- name: CreateTokenChains :exec
INSERT INTO token_chains (token_id, chain_id) 
VALUES (@token_id, @chain_id);