-- name: CreateOrder :exec
INSERT INTO orders (
  order_id,
  account_id,
  market_id,
  side,
  type,
  method,
  amount,
  amount_filled,
  status,
  price,
  expires_at
) VALUES (
  @order_id, @account_id, @market_id, @side, @type, @method, @amount, @amount_filled, @status, @price, @expires_at
);

-- name: UpdateOrder :exec
UPDATE orders
SET
  amount = COALESCE(sqlc.narg('amount'), amount),
  price = COALESCE(sqlc.narg('price'), price),
  updated_at = NOW()
WHERE
  order_id = sqlc.arg('order_id')
  AND deleted_at IS NULL;

-- name: GetUserOrderHistory :many
SELECT 
  *
FROM orders
WHERE 
  (
    status = sqlc.narg('status')
    OR sqlc.narg('status') IS NULL
  ) 
  AND account_id = @account_id
  AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

