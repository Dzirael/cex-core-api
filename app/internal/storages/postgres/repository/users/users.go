package users_repo

import (
	"context"
	"errors"
	"fmt"

	"cex-core-api/app/internal/apperrors"
	"cex-core-api/app/internal/models"
	"cex-core-api/app/internal/storages/postgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	repo *sqlc.Queries
	pool *pgxpool.Pool
}

func NewUsersRepository(repo *sqlc.Queries, pool *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{repo: repo, pool: pool}
}

func (r *UsersRepository) CreateUser(ctx context.Context, params sqlc.CreateUserParams) (*models.User, error) {
	err := r.repo.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("sqlc: CreateUser: %w", err)
	}

	return r.GetUserByID(ctx, params.UserID)
}

func (r *UsersRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	user, err := r.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetUserByID: %w", err)
	}

	return userToModel(user), nil
}

func (r *UsersRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := r.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Join(err, apperrors.ErrNotFound)
		}

		return nil, fmt.Errorf("sqlc: GetUserByID: %w", err)
	}

	return userToModel(user), nil
}
