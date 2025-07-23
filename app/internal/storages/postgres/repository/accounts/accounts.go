package accounts_repo

import (
	"context"
	"fmt"

	"cex-core-api/app/internal/models"
	"cex-core-api/app/internal/storages/postgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepository struct {
	repo *sqlc.Queries
	pool *pgxpool.Pool
}

func NewAccountRepository(repo *sqlc.Queries, pool *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{repo: repo, pool: pool}
}

func (r *AccountRepository) CreateAccount(ctx context.Context, params sqlc.CreateAccountParams) (*models.Account, error) {
	err := r.repo.CreateAccount(ctx, params)
	if err != nil {
		return nil, err
	}

	return r.GetAccountByID(ctx, params.AccountID)
}

func (r *AccountRepository) GetAccountsByUserID(ctx context.Context, userID uuid.UUID) ([]*models.Account, error) {
	accounts, err := r.repo.GetAccountsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	out := make([]*models.Account, len(accounts))
	for i, account := range accounts {
		out[i] = accountToModel(account)
	}
	return out, nil
}

func (r *AccountRepository) GetAccountByID(ctx context.Context, accountID uuid.UUID) (*models.Account, error) {
	account, err := r.repo.GetAccountByID(ctx, accountID)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetAccountByID: %w", err)
	}

	return accountToModel(account), nil
}
