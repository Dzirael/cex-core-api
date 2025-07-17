-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chains (
    chain_id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tokens (
    token_id UUID PRIMARY KEY NOT NULL,
    is_native BOOLEAN NOT NULL DEFAULT FALSE,
    name VARCHAR(100) NOT NULL,
    symbol VARCHAR(20) NOT NULL,
    decimals INTEGER NOT NULL CHECK (decimals >= 0),
    logo_path VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS token_chains (
    token_id UUID NOT NULL REFERENCES tokens(token_id) ON DELETE CASCADE,
    chain_id UUID NOT NULL REFERENCES chains(chain_id) ON DELETE CASCADE,
    PRIMARY KEY (token_id, chain_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS token_chains;
DROP TABLE IF EXISTS tokens;
DROP TABLE IF EXISTS chains;
-- +goose StatementEnd
