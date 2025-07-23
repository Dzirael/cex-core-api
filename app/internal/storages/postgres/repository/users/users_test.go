package users_repo_test

import (
	"context"
	"testing"

	"cex-core-api/app/internal/apperrors"
	users_repo "cex-core-api/app/internal/storages/postgres/repository/users"
	"cex-core-api/app/internal/storages/postgres/sqlc"
	"cex-core-api/app/test/helpers"

	"github.com/google/uuid"
	"github.com/test-go/testify/require"
)

func TestCreateUser(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	usersRepo := users_repo.NewUsersRepository(repo, pool)

	type testCase struct {
		name        string
		userID      uuid.UUID
		email       string
		userName    string
		userSurname string
		expectErr   bool
	}

	testCases := []testCase{
		{
			name:        "Create user with valid ID",
			userID:      uuid.New(),
			email:       "user1@example.com",
			userName:    "some name",
			userSurname: "some surname",
			expectErr:   false,
		},
		{
			name:        "Create user with empty UUID",
			userID:      uuid.Nil,
			email:       "invalid@example.com",
			userName:    "some name",
			userSurname: "some surname",
			expectErr:   false,
		},
		{
			name:        "Create user with already used emal",
			userID:      uuid.Nil,
			email:       "invalid@example.com",
			userName:    "some name",
			userSurname: "some surname",
			expectErr:   true,
		},
	}

	helpers.TruncateAllTables(t, pool)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			params := sqlc.CreateUserParams{
				UserID:  tc.userID,
				Email:   tc.email,
				Name:    "Test",
				Surname: "User",
			}

			_, err := usersRepo.CreateUser(ctx, params)

			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				user, err := usersRepo.GetUserByID(ctx, tc.userID)
				require.NoError(t, err)
				require.Equal(t, tc.userID, user.UserID)
				require.Equal(t, tc.email, user.Email)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	usersRepo := users_repo.NewUsersRepository(repo, pool)

	type testCase struct {
		name        string
		userID      uuid.UUID
		getSameID   bool
		email       string
		userName    string
		userSurname string
		expectErr   bool
	}

	testCases := []testCase{
		{
			name:        "Get user with valid ID",
			userID:      uuid.New(),
			getSameID:   true,
			email:       "user1@example.com",
			userName:    "some name",
			userSurname: "some surname",
			expectErr:   false,
		},
		{
			name:        "Get user with invalid UUID",
			userID:      uuid.New(),
			getSameID:   false,
			email:       "invalid@example.com",
			userName:    "some name",
			userSurname: "some surname",
			expectErr:   false,
		},
	}

	helpers.TruncateAllTables(t, pool)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			params := sqlc.CreateUserParams{
				UserID:  tc.userID,
				Email:   tc.email,
				Name:    "Test",
				Surname: "User",
			}

			_, err := usersRepo.CreateUser(ctx, params)

			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				if tc.getSameID {
					user, err := usersRepo.GetUserByID(ctx, tc.userID)
					require.NoError(t, err)
					require.Equal(t, tc.userID, user.UserID)
					require.Equal(t, tc.email, user.Email)
				} else {
					user, err := usersRepo.GetUserByID(ctx, uuid.New())
					require.Error(t, err)
					require.Nil(t, user)
				}
			}
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	usersRepo := users_repo.NewUsersRepository(repo, pool)

	type testCase struct {
		name      string
		email     string
		expectErr bool
		targetErr error
	}

	testCases := []testCase{
		{
			name:      "Get user with valid email",
			email:     "some@email.com",
			expectErr: false,
			targetErr: nil,
		},
		{
			name:      "Get user with non-existion email",
			email:     "invalid@example.com",
			expectErr: true,
			targetErr: apperrors.ErrNotFound,
		},
	}

	helpers.TruncateAllTables(t, pool)
	userID := uuid.Must(uuid.NewV7())
	params := sqlc.CreateUserParams{
		UserID:  userID,
		Email:   "some@email.com",
		Name:    "Test",
		Surname: "User",
	}

	_, err := usersRepo.CreateUser(t.Context(), params)
	require.NoError(t, err)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := usersRepo.GetUserByEmail(t.Context(), tc.email)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				require.NoError(t, err)
				require.Equal(t, tc.email, user.Email)
				require.Equal(t, userID, user.UserID)
			}
		})
	}
}
