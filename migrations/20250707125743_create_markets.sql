-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS markets (
    market_id UUID PRIMARY KEY NOT NULL,
    token_a_id UUID NOT NULL REFERENCES tokens(token_id) ON DELETE CASCADE,
    token_b_id UUID NOT NULL REFERENCES tokens(token_id) ON DELETE CASCADE,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    min_order_amount DECIMAL NOT NULL CHECK (min_order_amount >= 0),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    started_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TYPE order_side AS ENUM ('buy', 'sell');
CREATE TYPE order_type AS ENUM ('market', 'limit');
CREATE TYPE order_method AS ENUM ('FOK', 'IOK', 'GTC');
CREATE TYPE order_status AS ENUM ('created', 'partially_filled', 'filled', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
    order_id UUID PRIMARY KEY NOT NULL,
    account_id UUID NOT NULL REFERENCES accounts(account_id) ON DELETE CASCADE,
    market_id UUID NOT NULL REFERENCES markets(market_id) ON DELETE CASCADE,
    side order_side NOT NULL,
    type order_type NOT NULL,
    method order_method NOT NULL,
    amount DECIMAL NOT NULL CHECK (amount >= 0),
    amount_filled DECIMAL NOT NULL CHECK (amount_filled >= 0),
    status order_status NOT NULL,
    price DECIMAL CHECK (price >= 0), -- for market orders price can be null
    expires_at TIMESTAMP NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_trade (
    trade_id UUID PRIMARY KEY NOT NULL,
    order_id UUID NOT NULL REFERENCES orders(order_id) ON DELETE CASCADE,
    account_a_id UUID NOT NULL REFERENCES accounts(account_id) ON DELETE CASCADE,
    account_b_id UUID NOT NULL REFERENCES accounts(account_id) ON DELETE CASCADE,
    amount_filled DECIMAL NOT NULL CHECK (amount_filled >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_trade;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS markets;

DROP TYPE IF EXISTS order_status;
DROP TYPE IF EXISTS order_method;
DROP TYPE IF EXISTS order_type;
DROP TYPE IF EXISTS order_side;
-- +goose StatementEnd

