package market

import (
	"cex-core-api/app/internal/storages/postgres/sqlc"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MarketRepository struct {
	repo *sqlc.Queries
	pool *pgxpool.Pool
}

func NewMarketRepository(repo *sqlc.Queries, pool *pgxpool.Pool) *MarketRepository {
	return &MarketRepository{repo: repo, pool: pool}
}

func (r *MarketRepository) CreateMarket(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}

func (r *MarketRepository) GetSupportedPairs(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}

func (r *MarketRepository) GetMarketPair(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}
