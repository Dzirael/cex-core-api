package service

import (
	"context"

	"cex-core-api/app/internal/models"
	"cex-core-api/app/internal/models/credentials"
	accounts_repo "cex-core-api/app/internal/storages/postgres/repository/accounts"
	"cex-core-api/app/internal/storages/postgres/sqlc"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, params sqlc.CreateAccountParams) (*models.Account, error)
	GetAccountsByUserID(ctx context.Context, userID uuid.UUID) ([]*models.Account, error)
	GetAccountByID(ctx context.Context, accountID uuid.UUID) (*models.Account, error)

	GetAccountBalances(ctx context.Context, params sqlc.GetTokenBalanceByAccountIDParams) ([]*models.AccountBalanceResult, error)
	UpdateAccountBalance(ctx context.Context, params accounts_repo.ChangeAccountBalanceParams) (*models.UpdateAccountBalanceResult, error)
	GetBalanceChanges(ctx context.Context, params sqlc.GetBalanceChangesParams) ([]*models.AccountBalanceChange, error)
	GetBalanceChangeByID(ctx context.Context, changeID uuid.UUID) (*models.AccountBalanceChange, error)
	DecreaseAccountLockedBalance(ctx context.Context, balanceID uuid.UUID, amount decimal.Decimal) error

	CreateIntenalTransfer(ctx context.Context, params accounts_repo.ChangeTransferParams) (*models.InternalTransferResult, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, params sqlc.CreateUserParams) (*models.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error)

	CreateCredential(ctx context.Context, params sqlc.CreateCredentialParams) (*credentials.UserCredentials, error)
	GetCredentialByID(ctx context.Context, credentialID uuid.UUID) (*credentials.UserCredentials, error)
	GetUserCredentials(ctx context.Context, userID uuid.UUID) ([]*credentials.UserCredentials, error)
	GetUserCredentialByType(ctx context.Context, userID uuid.UUID, t credentials.Type) (*credentials.UserCredentials, error)
}
