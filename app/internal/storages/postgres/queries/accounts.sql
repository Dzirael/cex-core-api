-- name: CreateAccount :exec
INSERT INTO accounts (account_id, user_id, type) VALUES (@account_id, @user_id, @type);

-- name: GetAccountByID :one
SELECT * FROM accounts
WHERE account_id = @account_id;

-- name: GetAccountsByUserID :many
SELECT * FROM accounts
WHERE user_id = @user_id;

-- name: IncreaseAccountBalance :one
INSERT INTO account_balances (balance_id, account_id, token_id, amount, locked_amount)
VALUES (@balance_id, @account_id, @token_id, @increase_amount, @increase_locked_amount)
ON CONFLICT (account_id, token_id) DO UPDATE
SET amount = account_balances.amount + EXCLUDED.amount,
    locked_amount = account_balances.locked_amount + EXCLUDED.locked_amount,
    updated_at = NOW()
RETURNING balance_id;

-- name: DecreaseAccountBalance :one
WITH updated AS (
    UPDATE account_balances
    SET amount = amount - @decrease_amount,
        updated_at = NOW()
    WHERE account_id = @account_id
      AND token_id = @token_id
      AND (amount - locked_amount) >= @decrease_amount
    RETURNING *
)
SELECT EXISTS(SELECT 1 FROM updated) AS success,
       (SELECT balance_id FROM updated LIMIT 1) AS balance_id;


-- name: DecreaseAccountLockedBalance :exec
UPDATE account_balances
SET locked_amount = locked_amount - @amount
WHERE balance_id = @balance_id;

-- name: CreateBalanceTransfer :exec
INSERT INTO account_balance_changes (change_id, account_id, token_id, type,action, status, amount, sender, recipient)
VALUES (@change_id, @account_id, @token_id, @type, @action, @status, @amount, @sender, @recipient);

-- name: UpdateBalanceTransferStatus :exec 
UPDATE account_balance_changes
SET status = @status
WHERE change_id = @change_id;

-- name: GetBalanceChanges :many
SELECT * FROM account_balance_changes
WHERE account_id = @account_id
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetBalanceChangeByID :one
SELECT * FROM account_balance_changes
WHERE change_id = @change_id;

-- name: GetTokenBalanceByAccountID :many
SELECT 
    ab.balance_id,
    ab.amount,
    ab.locked_amount,
    ab.created_at,
    ab.updated_at,
    
    t.token_id,
    t.is_native,
    t.name,
    t.symbol,
    t.decimals,
    t.logo_path
FROM account_balances ab
JOIN tokens t ON ab.token_id = t.token_id
WHERE 
    ab.account_id = @account_id AND
    ab.deleted_at IS NULL AND
    t.deleted_at IS NULL
ORDER BY t.name ASC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');