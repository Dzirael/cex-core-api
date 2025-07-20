package accounts_repo_test

import (
	"context"
	"testing"

	"cex-core-api/app/internal/models"
	accounts_repo "cex-core-api/app/internal/storages/postgres/repository/accounts"
	"cex-core-api/app/internal/storages/postgres/sqlc"
	"cex-core-api/app/test/helpers"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func prepareUserAccountWithBalance(t *testing.T, userID, accountID, tokenID uuid.UUID) {
	ctx := t.Context()
	_, _, repo := helpers.GetDatabaseContainer(t)

	// Setup prerequisites
	err := repo.CreateUser(ctx, sqlc.CreateUserParams{
		UserID:  userID,
		Email:   "test@example.com",
		Name:    "Test",
		Surname: "User",
	})
	require.NoError(t, err)

	err = repo.CreateAccount(ctx, sqlc.CreateAccountParams{
		AccountID: accountID,
		UserID:    userID,
		Type:      "spot",
	})
	require.NoError(t, err)

	err = repo.CreateToken(ctx, sqlc.CreateTokenParams{
		TokenID:  tokenID,
		IsNative: true,
		Name:     "USDT",
		Symbol:   "USDT",
		Decimals: 6,
		LogoPath: helpers.ToPointer("usdt.png"),
	})
	require.NoError(t, err)

	// Create initial balance
	balanceID, err := repo.IncreaseAccountBalance(ctx, sqlc.IncreaseAccountBalanceParams{
		BalanceID:      uuid.New(),
		AccountID:      accountID,
		TokenID:        tokenID,
		IncreaseAmount: decimal.NewFromInt(100),
	})
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, balanceID)
}

func TestUpdateAccountBalance(t *testing.T) {
	ctx := t.Context()
	_, pool, repo := helpers.GetDatabaseContainer(t)
	accountRepo := accounts_repo.NewAccountRepository(repo, pool)

	// Create user, account, token, and balance before testing
	userID := uuid.New()
	accountID := uuid.New()
	tokenID := uuid.New()

	helpers.TruncateAllTables(t, pool)
	prepareUserAccountWithBalance(t, userID, accountID, tokenID)

	type testCase struct {
		name      string
		params    accounts_repo.ChangeAccountBalanceParams
		expectErr bool
		checkFn   func(t *testing.T)
	}

	testCases := []testCase{
		{
			name: "Increase balance",
			params: accounts_repo.ChangeAccountBalanceParams{
				AccountID:    accountID,
				TokenID:      tokenID,
				Type:         models.ChangeTypeIncrease,
				Action:       models.ChangeActionDeposit,
				Status:       models.ChangeStatusPending,
				ChangeAmount: decimal.NewFromInt(50),
				Sender:       userID.String(),
				Recipient:    "some recipient",
			},
			expectErr: false,
			checkFn: func(t *testing.T) {
				balances, err := accountRepo.GetAccountBalances(ctx, sqlc.GetTokenBalanceByAccountIDParams{
					AccountID: accountID,
					Offset:    0,
					Limit:     10,
				})

				require.NoError(t, err)
				require.Len(t, balances, 1)
				require.Equal(t, decimal.NewFromInt(150), balances[0].Amount)
				require.Equal(t, decimal.NewFromInt(50), balances[0].LockedAmount)
			},
		},
		{
			name: "Reduce balance",
			params: accounts_repo.ChangeAccountBalanceParams{
				AccountID:    accountID,
				TokenID:      tokenID,
				Type:         models.ChangeTypeReduce,
				Action:       models.ChangeActionWithdraw,
				Status:       models.ChangeStatusPending,
				ChangeAmount: decimal.NewFromInt(30),
				Sender:       userID.String(),
				Recipient:    "user",
			},
			expectErr: false,
			checkFn: func(t *testing.T) {
				balances, err := accountRepo.GetAccountBalances(ctx, sqlc.GetTokenBalanceByAccountIDParams{
					AccountID: accountID,
					Offset:    0,
					Limit:     10,
				})
				require.NoError(t, err)
				require.Len(t, balances, 1)
				require.Equal(t, decimal.NewFromInt(120), balances[0].Amount)
				require.Equal(t, decimal.NewFromInt(50), balances[0].LockedAmount)
			},
		},
		{
			name: "Reduce balance with locked funds",
			params: accounts_repo.ChangeAccountBalanceParams{
				AccountID:    accountID,
				TokenID:      tokenID,
				Type:         models.ChangeTypeReduce,
				Action:       models.ChangeActionWithdraw,
				Status:       models.ChangeStatusPending,
				ChangeAmount: decimal.NewFromInt(100),
				Sender:       "exchange",
				Recipient:    "user",
			},
			expectErr: true,
			checkFn:   nil,
		},
		{
			name: "Reduce balance with insufficient funds",
			params: accounts_repo.ChangeAccountBalanceParams{
				AccountID:    accountID,
				TokenID:      tokenID,
				Type:         models.ChangeTypeReduce,
				Action:       models.ChangeActionWithdraw,
				Status:       models.ChangeStatusPending,
				ChangeAmount: decimal.NewFromInt(9999),
				Sender:       "exchange",
				Recipient:    "user",
			},
			expectErr: true,
			checkFn:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			update, err := accountRepo.UpdateAccountBalance(ctx, tc.params)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotEqual(t, uuid.Nil, update.BalanceID)
				require.NotEqual(t, uuid.Nil, update.ChangeID)
				if tc.checkFn != nil {
					tc.checkFn(t)
				}
			}
		})
	}
}

func TestGetBalanceTransfers(t *testing.T) {
	ctx := context.Background()
	_, pool, repo := helpers.GetDatabaseContainer(t)
	accountRepo := accounts_repo.NewAccountRepository(repo, pool)

	// Create user, account, token, and balance before testing
	userID := uuid.New()
	accountID := uuid.New()
	tokenID := uuid.New()

	helpers.TruncateAllTables(t, pool)
	prepareUserAccountWithBalance(t, userID, accountID, tokenID)

	// Create several transfers
	transfersToCreate := []sqlc.CreateBalanceTransferParams{
		{
			ChangeID:  uuid.New(),
			AccountID: accountID,
			TokenID:   tokenID,
			Type:      models.ChangeTypeIncrease,
			Action:    models.ChangeActionDeposit,
			Status:    models.ChangeStatusCompleted,
			Amount:    decimal.NewFromInt(100),
			Sender:    "external",
			Recipient: "user",
		},
		{
			ChangeID:  uuid.New(),
			AccountID: accountID,
			TokenID:   tokenID,
			Type:      models.ChangeTypeReduce,
			Action:    models.ChangeActionWithdraw,
			Status:    models.ChangeStatusCompleted,
			Amount:    decimal.NewFromInt(50),
			Sender:    "user",
			Recipient: "external",
		},
	}

	for _, transfer := range transfersToCreate {
		err := repo.CreateBalanceTransfer(ctx, transfer)
		require.NoError(t, err)
	}

	testCases := []struct {
		name      string
		params    sqlc.GetBalanceChangesParams
		expectErr bool
		checkFn   func(t *testing.T, results []*models.AccountBalanceChange)
	}{
		{
			name: "Get all transfers for account and token",
			params: sqlc.GetBalanceChangesParams{
				AccountID: accountID,
				Limit:     10,
				Offset:    0,
			},
			expectErr: false,
			checkFn: func(t *testing.T, results []*models.AccountBalanceChange) {
				require.Len(t, results, len(transfersToCreate))
				for i, r := range results {
					expected := transfersToCreate[len(transfersToCreate)-1-i]
					require.Equal(t, expected.AccountID, r.AccountID)
					require.Equal(t, expected.TokenID, r.TokenID)
					require.Equal(t, expected.Amount, r.Amount)
				}
			},
		},
		{
			name: "No transfers for unknown account",
			params: sqlc.GetBalanceChangesParams{
				AccountID: uuid.New(),
				Limit:     10,
				Offset:    0,
			},
			expectErr: false,
			checkFn: func(t *testing.T, results []*models.AccountBalanceChange) {
				require.Empty(t, results)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			results, err := accountRepo.GetBalanceChanges(ctx, tc.params)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tc.checkFn != nil {
					tc.checkFn(t, results)
				}
			}
		})
	}
}
