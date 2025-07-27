-- name: CreateMarket :exec
INSERT INTO markets (
  market_id,
  token_a_id,
  token_b_id,
  is_active,
  min_order_amount,
  started_at
) VALUES (
  @market_id, @token_a_id, @token_b_id, @is_active, @min_order_amount, @started_at
);

-- name: GetSupportedPairs :many
SELECT 
  m.market_id,
  m.token_a_id,
  ta.name AS token_a_name,
  m.token_b_id,
  tb.name AS token_b_name,
  m.min_order_amount,
  m.started_at,
  m.created_at,
  m.updated_at
FROM markets m
JOIN tokens ta ON m.token_a_id = ta.token_id
JOIN tokens tb ON m.token_b_id = tb.token_id
WHERE 
  m.deleted_at IS NULL
  AND m.is_active = TRUE
ORDER BY m.created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetMarketPair :one
SELECT *
FROM markets
WHERE 
  token_a_id = @token_a_id AND token_b_id = @token_b_id AND deleted_at IS NULL;

-- name: GetMarketByID :one
SELECT 
  *
FROM markets
WHERE market_id = @market_id AND deleted_at IS NULL;

