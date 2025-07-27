-- name: GetChains :many
SELECT * FROM chains
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetChainByID :one
SELECT * FROM chains
WHERE chain_id = @chain_id AND deleted_at IS NULL;

-- name: CreateChain :exec
INSERT INTO chains (chain_id, name, created_at, updated_at) 
VALUES (@chain_id, @name, NOW(), NOW());

-- name: GetSupportedTokens :many
SELECT DISTINCT ON (t.token_id)
  t.*,
  (
    SELECT ARRAY_AGG(tc2.chain_id)::uuid[]
    FROM token_chains tc2
    WHERE tc2.token_id = t.token_id
  ) AS chain_ids
FROM tokens t
JOIN token_chains tc ON t.token_id = tc.token_id
WHERE 
  (
    tc.chain_id = sqlc.narg('chain_id')
    OR sqlc.narg('chain_id') IS NULL
  ) 
  AND (
    t.is_native = sqlc.narg('is_native')
    OR sqlc.narg('is_native') IS NULL
  )
  AND t.deleted_at IS NULL
ORDER BY t.token_id, t.created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetTokenByID :one
SELECT * FROM tokens
WHERE token_id = @token_id AND deleted_at IS NULL;

-- name: GetTokenBySymbol :one
SELECT * FROM tokens
WHERE symbol = @symbol AND deleted_at IS NULL;

-- name: CreateToken :exec
INSERT INTO tokens (token_id, is_native, name, symbol, decimals, logo_path, created_at, updated_at) 
VALUES (@token_id, @is_native, @name, @symbol, @decimals, @logo_path, NOW(), NOW());

-- name: CreateTokenChains :exec
INSERT INTO token_chains (token_id, chain_id)
SELECT @token_id, UNNEST(@chain_ids::uuid[]);