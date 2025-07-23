package accounts_repo_test

import (
	"context"
	"fmt"
	"testing"

	"cex-core-api/app/internal/models"
	accounts_repo "cex-core-api/app/internal/storages/postgres/repository/accounts"
	"cex-core-api/app/internal/storages/postgres/sqlc"
	"cex-core-api/app/test/helpers"

	"github.com/google/uuid"
	"github.com/test-go/testify/require"
)

func TestCreateAccount(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	accountRepo := accounts_repo.NewAccountRepository(repo, pool)

	type testCase struct {
		name        string
		userID      uuid.UUID
		accountID   uuid.UUID
		accountType models.AccountType
		wantErr     bool
	}

	testCases := []testCase{
		{
			name:        "Create spot account",
			userID:      uuid.New(),
			accountID:   uuid.New(),
			accountType: models.AccountTypeSpot,
			wantErr:     false,
		},
		{
			name:        "Create margin account",
			userID:      uuid.New(),
			accountID:   uuid.New(),
			accountType: models.AccountTypeFutures,
			wantErr:     false,
		},
		{
			name:        "Create account with missing user",
			userID:      uuid.Nil, // assuming FK will fail
			accountID:   uuid.New(),
			accountType: models.AccountTypeSpot,
			wantErr:     true,
		},
	}

	helpers.TruncateAllTables(t, pool)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			if tc.userID != uuid.Nil {
				err := repo.CreateUser(ctx, sqlc.CreateUserParams{
					UserID:  tc.userID,
					Email:   fmt.Sprintf("test_%s@example.com", tc.userID.String()[:8]),
					Name:    "some name",
					Surname: "some surname",
				})
				require.NoError(t, err)
			}

			accountID, err := accountRepo.CreateAccount(ctx, sqlc.CreateAccountParams{
				AccountID: tc.accountID,
				UserID:    tc.userID,
				Type:      tc.accountType,
			})
			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotEqual(t, uuid.Nil, accountID)

			account, err := accountRepo.GetAccountByID(ctx, tc.accountID)
			require.NoError(t, err)
			require.Equal(t, tc.userID, account.UserID)
			require.Equal(t, tc.accountType, account.Type)
		})
	}
}

func TestGetAccountsByUserID(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	accountRepo := accounts_repo.NewAccountRepository(repo, pool)

	type testCase struct {
		name         string
		userID       uuid.UUID
		accountTypes []models.AccountType
		wantErr      bool
	}

	testCases := []testCase{
		{
			name:         "User with spot and futures accounts",
			userID:       uuid.New(),
			accountTypes: []models.AccountType{models.AccountTypeSpot, models.AccountTypeFutures},
			wantErr:      false,
		},
		{
			name:         "User with no accounts",
			userID:       uuid.New(),
			accountTypes: []models.AccountType{},
			wantErr:      false,
		},
		{
			name:         "Non-existent user",
			userID:       uuid.New(),
			accountTypes: nil,
			wantErr:      false,
		},
	}

	helpers.TruncateAllTables(t, pool)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			if !tc.wantErr {
				err := repo.CreateUser(ctx, sqlc.CreateUserParams{
					UserID:  tc.userID,
					Email:   fmt.Sprintf("user_%s@example.com", tc.userID.String()[:8]),
					Name:    "Test",
					Surname: "User",
				})
				require.NoError(t, err)

				for _, acctType := range tc.accountTypes {
					_, err = accountRepo.CreateAccount(ctx, sqlc.CreateAccountParams{
						AccountID: uuid.New(),
						UserID:    tc.userID,
						Type:      acctType,
					})
					require.NoError(t, err)
				}
			}

			accounts, err := accountRepo.GetAccountsByUserID(ctx, tc.userID)

			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Len(t, accounts, len(tc.accountTypes))

			for i, acctType := range tc.accountTypes {
				require.Equal(t, tc.userID, accounts[i].UserID)
				require.Equal(t, acctType, accounts[i].Type)
			}
		})
	}
}

func TestGetAccountByID(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	accountRepo := accounts_repo.NewAccountRepository(repo, pool)

	type testCase struct {
		name             string
		userID           uuid.UUID
		accountID        uuid.UUID
		getSameAccountID bool
		accountType      models.AccountType
		wantErr          bool
	}

	testCases := []testCase{
		{
			name:             "Get spot account with valid id",
			userID:           uuid.New(),
			accountID:        uuid.New(),
			getSameAccountID: true,
			accountType:      models.AccountTypeSpot,
			wantErr:          false,
		},
		{
			name:             "Get futures account with valid id",
			userID:           uuid.New(),
			accountID:        uuid.New(),
			getSameAccountID: true,
			accountType:      models.AccountTypeFutures,
			wantErr:          false,
		},
		{
			name:             "Get spot account with invalid id",
			userID:           uuid.New(),
			accountID:        uuid.New(),
			getSameAccountID: false,
			accountType:      models.AccountTypeSpot,
			wantErr:          false,
		},
		{
			name:             "Get futures account with invalid id",
			userID:           uuid.New(),
			accountID:        uuid.New(),
			getSameAccountID: false,
			accountType:      models.AccountTypeFutures,
			wantErr:          false,
		},
	}

	helpers.TruncateAllTables(t, pool)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			if tc.userID != uuid.Nil {
				err := repo.CreateUser(ctx, sqlc.CreateUserParams{
					UserID:  tc.userID,
					Email:   fmt.Sprintf("test_%s@example.com", tc.userID.String()[:8]),
					Name:    "some name",
					Surname: "some surname",
				})
				require.NoError(t, err)
			}

			accountID, err := accountRepo.CreateAccount(ctx, sqlc.CreateAccountParams{
				AccountID: tc.accountID,
				UserID:    tc.userID,
				Type:      tc.accountType,
			})
			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotEqual(t, uuid.Nil, accountID)

			if tc.getSameAccountID {
				account, err := accountRepo.GetAccountByID(ctx, tc.accountID)
				require.NoError(t, err)
				require.Equal(t, tc.userID, account.UserID)
				require.Equal(t, tc.accountType, account.Type)
			} else {
				account, err := accountRepo.GetAccountByID(ctx, uuid.New())
				require.Error(t, err)
				require.Nil(t, account)
			}
		})
	}
}
