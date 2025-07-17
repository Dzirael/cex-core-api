-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TYPE account_type AS ENUM ('spot', 'margin', 'futures');

CREATE TABLE IF NOT EXISTS accounts (
    account_id UUID PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    type account_type NOT NULL DEFAULT 'spot',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS account_balances (
    balance_id UUID PRIMARY KEY NOT NULL,
    account_id UUID NOT NULL REFERENCES accounts(account_id) ON DELETE CASCADE,
    token_id UUID NOT NULL REFERENCES tokens(token_id) ON DELETE CASCADE,
    amount DECIMAL NOT NULL CHECK (amount >= 0),
    locked_amount DECIMAL NOT NULL CHECK (locked_amount >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

ALTER TABLE account_balances
ADD CONSTRAINT unique_account_token UNIQUE (account_id, token_id);

CREATE TYPE change_type AS ENUM ('reduce', 'increase');
CREATE TYPE change_action AS ENUM ('order', 'transfer');

CREATE TABLE IF NOT EXISTS account_balance_changes (
    change_id UUID PRIMARY KEY NOT NULL,
    account_id UUID NOT NULL REFERENCES accounts(account_id) ON DELETE CASCADE,
    token_id UUID NOT NULL REFERENCES tokens(token_id) ON DELETE CASCADE,
    type change_type NOT NULL,
    action change_action NOT NULL,
    amount DECIMAL NOT NULL CHECK (amount >= 0),
    sender VARCHAR(50) NOT NULL,
    recipient VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- CREATE TYPE transfer_type AS ENUM ('deposit', 'withdraw');


-- CREATE TABLE IF NOT EXISTS account_chain_transfers (
--     transfer_id UUID PRIMARY KEY NOT NULL,
--     account_id UUID NOT NULL REFERENCES accounts(account_id) ON DELETE CASCADE,
--     token_id UUID NOT NULL REFERENCES tokens(token_id) ON DELETE CASCADE,
--     chain_id UUID NOT NULL REFERENCES chains(chain_id) ON DELETE CASCADE,
--     type transfer_type NOT NULL,
--     amount DECIMAL NOT NULL CHECK (amount >= 0),
--     recipient VARCHAR(50) NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     deleted_at TIMESTAMP
-- );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account_chain_transfers;
DROP TABLE IF EXISTS account_balance_changes;
DROP TABLE IF EXISTS account_balances;
DROP TABLE IF EXISTS accounts;
DROP TYPE IF EXISTS transfer_type;
DROP TYPE IF EXISTS change_type;
DROP TYPE IF EXISTS account_type;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
