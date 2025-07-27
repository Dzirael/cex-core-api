package crypto_repo_test

import (
	"testing"

	crypto_repo "cex-core-api/app/internal/storages/postgres/repository/crypto"
	"cex-core-api/app/internal/storages/postgres/sqlc"
	"cex-core-api/app/test/helpers"

	"github.com/google/uuid"
	"github.com/test-go/testify/require"
)

func TestCreateToken(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	cryptoRepo := crypto_repo.NewCryptoRepository(repo, pool)

	helpers.TruncateAllTables(t, pool)

	type testCase struct {
		name      string
		params    sqlc.CreateTokenParams
		expectErr bool
	}

	testCases := []testCase{
		{
			name: "Create valid token",
			params: sqlc.CreateTokenParams{
				TokenID:  uuid.New(),
				IsNative: false,
				Name:     "TestToken",
				Symbol:   "TTK",
				Decimals: 18,
				LogoPath: nil,
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := cryptoRepo.CreateToken(t.Context(), tc.params)

			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				token, err := cryptoRepo.GetTokenByID(t.Context(), tc.params.TokenID)
				require.NoError(t, err)
				require.Equal(t, tc.params.Name, token.Name)
				require.Equal(t, uint(tc.params.Decimals), token.Decimals)
				require.Equal(t, tc.params.IsNative, token.IsNative)
				require.Equal(t, tc.params.LogoPath, token.LogoPath)
				require.Equal(t, tc.params.TokenID, token.TokenID)
			}
		})
	}
}

func TestCreateChain(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	cryptoRepo := crypto_repo.NewCryptoRepository(repo, pool)

	helpers.TruncateAllTables(t, pool)

	type testCase struct {
		name      string
		chainID   uuid.UUID
		chainName string
		expectErr bool
	}

	testCases := []testCase{
		{
			name:      "Create valid chain",
			chainID:   uuid.New(),
			chainName: "Ethereum",
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := cryptoRepo.CreateChain(t.Context(), tc.chainID, tc.chainName)

			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				chain, err := cryptoRepo.GetChainByID(t.Context(), tc.chainID)
				require.NoError(t, err)
				require.Equal(t, tc.chainID, chain.ChainID)
				require.Equal(t, tc.chainName, chain.Name)
			}
		})
	}
}

func TestAddChainToToken(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	cryptoRepo := crypto_repo.NewCryptoRepository(repo, pool)

	helpers.TruncateAllTables(t, pool)

	validTokenID := uuid.New()
	require.NoError(t, cryptoRepo.CreateToken(t.Context(), sqlc.CreateTokenParams{
		TokenID:  validTokenID,
		IsNative: false,
		Name:     "ValidToken",
		Symbol:   "VAL",
		Decimals: 18,
		LogoPath: nil,
	}))

	chainID1 := uuid.New()
	chainID2 := uuid.New()
	require.NoError(t, cryptoRepo.CreateChain(t.Context(), chainID1, "ChainOne"))
	require.NoError(t, cryptoRepo.CreateChain(t.Context(), chainID2, "ChainTwo"))

	type testCase struct {
		name      string
		tokenID   uuid.UUID
		chainIDs  []uuid.UUID
		expectErr bool
	}

	testCases := []testCase{
		{
			name:      "Valid token and two valid chains",
			tokenID:   validTokenID,
			chainIDs:  []uuid.UUID{chainID1, chainID2},
			expectErr: false,
		},
		{
			name:      "Invalid token ID",
			tokenID:   uuid.New(), // not created
			chainIDs:  []uuid.UUID{chainID1},
			expectErr: true,
		},
		{
			name:      "Invalid chain ID",
			tokenID:   validTokenID,
			chainIDs:  []uuid.UUID{uuid.New()}, // non-existent chain
			expectErr: true,
		},
		{
			name:      "Empty chain list",
			tokenID:   validTokenID,
			chainIDs:  []uuid.UUID{},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := cryptoRepo.AddChainToToken(t.Context(), tc.tokenID, tc.chainIDs)

			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGetSupportedChains(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	cryptoRepo := crypto_repo.NewCryptoRepository(repo, pool)

	helpers.TruncateAllTables(t, pool)

	chainID1 := uuid.New()
	chainID2 := uuid.New()
	require.NoError(t, cryptoRepo.CreateChain(t.Context(), chainID1, "ChainOne"))
	require.NoError(t, cryptoRepo.CreateChain(t.Context(), chainID2, "ChainTwo"))

	type testCase struct {
		name      string
		chainIDs  []uuid.UUID
		expectErr bool
	}

	testCases := []testCase{
		{
			name:      "Get supported chains",
			chainIDs:  []uuid.UUID{chainID1, chainID2},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chains, err := cryptoRepo.GetSupportedChains(t.Context(), 100, 0)

			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Len(t, chains, 2)
			}
		})
	}
}

func TestGetSupportedTokens(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	cryptoRepo := crypto_repo.NewCryptoRepository(repo, pool)

	helpers.TruncateAllTables(t, pool)

	validTokenID := uuid.New()
	require.NoError(t, cryptoRepo.CreateToken(t.Context(), sqlc.CreateTokenParams{
		TokenID:  validTokenID,
		IsNative: false,
		Name:     "ValidToken",
		Symbol:   "VAL",
		Decimals: 18,
		LogoPath: nil,
	}))

	chainID1 := uuid.New()
	chainID2 := uuid.New()
	chainID3 := uuid.New()
	require.NoError(t, cryptoRepo.CreateChain(t.Context(), chainID1, "ChainOne"))
	require.NoError(t, cryptoRepo.CreateChain(t.Context(), chainID2, "ChainTwo"))

	require.NoError(t, cryptoRepo.AddChainToToken(t.Context(), validTokenID, []uuid.UUID{chainID1, chainID2}))

	type testCase struct {
		name             string
		chainIDs         *uuid.UUID
		expectedChainLen int
		expectedTokenLen int
		expectErr        bool
	}

	testCases := []testCase{
		{
			name:             "Get supported tokens, without filter",
			chainIDs:         nil,
			expectedTokenLen: 1,
			expectedChainLen: 2,
			expectErr:        false,
		},
		{
			name:             "Get supported tokens, with valid filter",
			chainIDs:         &chainID1,
			expectedTokenLen: 1,
			expectedChainLen: 2,
			expectErr:        false,
		},
		{
			name:             "Get supported tokens, with invalid filter",
			chainIDs:         &chainID3,
			expectedTokenLen: 0,
			expectedChainLen: 0,
			expectErr:        false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := cryptoRepo.GetSupportedTokens(t.Context(), sqlc.GetSupportedTokensParams{
				ChainID: tc.chainIDs,
				Limit:   100,
				Offset:  0,
			})

			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Len(t, tokens, tc.expectedTokenLen)
				if tc.expectedChainLen != 0 {
					require.Len(t, tokens[0].ChainIDs, tc.expectedChainLen)

				}
			}
		})
	}
}

func TestGetTokenBySymbol(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	cryptoRepo := crypto_repo.NewCryptoRepository(repo, pool)

	helpers.TruncateAllTables(t, pool)

	type testCase struct {
		name      string
		params    sqlc.CreateTokenParams
		expectErr bool
	}

	testCases := []testCase{
		{
			name: "Create valid token",
			params: sqlc.CreateTokenParams{
				TokenID:  uuid.New(),
				IsNative: false,
				Name:     "TestToken",
				Symbol:   "TTK",
				Decimals: 18,
				LogoPath: nil,
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := cryptoRepo.CreateToken(t.Context(), tc.params)

			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				token, err := cryptoRepo.GetTokenBySymbol(t.Context(), tc.params.Symbol)
				require.NoError(t, err)
				require.Equal(t, tc.params.Name, token.Name)
				require.Equal(t, uint(tc.params.Decimals), token.Decimals)
				require.Equal(t, tc.params.IsNative, token.IsNative)
				require.Equal(t, tc.params.LogoPath, token.LogoPath)
				require.Equal(t, tc.params.TokenID, token.TokenID)
			}
		})
	}
}
