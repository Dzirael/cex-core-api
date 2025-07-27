package crypto_repo

import (
	"context"
	"fmt"

	"cex-core-api/app/internal/models"
	"cex-core-api/app/internal/storages/postgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CryptoRepository struct {
	repo *sqlc.Queries
	pool *pgxpool.Pool
}

func NewCryptoRepository(repo *sqlc.Queries, pool *pgxpool.Pool) *CryptoRepository {
	return &CryptoRepository{repo: repo, pool: pool}
}

func (r *CryptoRepository) CreateToken(ctx context.Context, params sqlc.CreateTokenParams) error {
	err := r.repo.CreateToken(ctx, params)
	if err != nil {
		return fmt.Errorf("sqlc: CreateToken: %w", err)
	}

	return nil
}

func (r *CryptoRepository) AddChainToToken(ctx context.Context, tokenID uuid.UUID, chainIDs []uuid.UUID) error {
	err := r.repo.CreateTokenChains(ctx, tokenID, chainIDs)
	if err != nil {
		return fmt.Errorf("sqlc: AddChainToToken: %w", err)
	}

	return nil
}

func (r *CryptoRepository) CreateChain(ctx context.Context, chainID uuid.UUID, name string) error {
	err := r.repo.CreateChain(ctx, chainID, name)
	if err != nil {
		return fmt.Errorf("sqlc: CreateChain: %w", err)
	}

	return nil
}

func (r *CryptoRepository) GetSupportedChains(ctx context.Context, limit, ofset int32) ([]*models.Chain, error) {
	chains, err := r.repo.GetChains(ctx, ofset, limit)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetChains: %w", err)
	}
	out := make([]*models.Chain, len(chains))
	for i, chain := range chains {
		out[i] = chainToModel(chain)
	}
	return out, nil
}

func (r *CryptoRepository) GetSupportedTokens(cxt context.Context, params sqlc.GetSupportedTokensParams) ([]*models.SupportedTokensResult, error) {
	rows, err := r.repo.GetSupportedTokens(cxt, params)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetChains: %w", err)
	}
	out := make([]*models.SupportedTokensResult, len(rows))
	for i, row := range rows {
		out[i] = supportedTokenToModel(row)
	}

	return out, nil
}

func (r *CryptoRepository) GetChainByID(cxt context.Context, chainID uuid.UUID) (*models.Chain, error) {
	chain, err := r.repo.GetChainByID(cxt, chainID)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetChainByID: %w", err)
	}

	return chainToModel(chain), nil
}

func (r *CryptoRepository) GetTokenByID(cxt context.Context, chainID uuid.UUID) (*models.Token, error) {
	token, err := r.repo.GetTokenByID(cxt, chainID)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetChainByID: %w", err)
	}

	return tokenToModel(token), nil
}

func (r *CryptoRepository) GetTokenBySymbol(cxt context.Context, symbol string) (*models.Token, error) {
	token, err := r.repo.GetTokenBySymbol(cxt, symbol)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetChainByID: %w", err)
	}

	return tokenToModel(token), nil
}
