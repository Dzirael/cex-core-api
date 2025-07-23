package users_repo_test

import (
	"context"
	"testing"

	"cex-core-api/app/internal/models/credentials"
	users_repo "cex-core-api/app/internal/storages/postgres/repository/users"
	"cex-core-api/app/internal/storages/postgres/sqlc"
	"cex-core-api/app/test/helpers"

	"github.com/google/uuid"
	"github.com/test-go/testify/require"
)

func TestCreateCredential(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	usersRepo := users_repo.NewUsersRepository(repo, pool)

	helpers.TruncateAllTables(t, pool)

	ctx := context.Background()
	userID := uuid.New()
	_, err := usersRepo.CreateUser(ctx, sqlc.CreateUserParams{
		UserID:  userID,
		Email:   "creduser@example.com",
		Name:    "Test",
		Surname: "Cred",
	})
	require.NoError(t, err)

	credentialID := uuid.New()
	params := sqlc.CreateCredentialParams{
		CredentialID: credentialID,
		UserID:       userID,
		Type:         credentials.PasswordType,
		IsPrimary:    true,
		IsVerified:   true,
		SecretData:   []byte(`"some-encrypted-secret"`),
	}

	cred, err := usersRepo.CreateCredential(ctx, params)
	require.NoError(t, err)
	require.NotNil(t, cred)
	require.Equal(t, credentialID, cred.CredentialID)
	require.Equal(t, userID, cred.UserID)
	require.Equal(t, credentials.PasswordType, cred.Type)
}

func TestGetCredentialByID(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	usersRepo := users_repo.NewUsersRepository(repo, pool)

	helpers.TruncateAllTables(t, pool)

	ctx := context.Background()
	userID := uuid.New()
	_, err := usersRepo.CreateUser(ctx, sqlc.CreateUserParams{
		UserID:  userID,
		Email:   "getcred@example.com",
		Name:    "Get",
		Surname: "Cred",
	})
	require.NoError(t, err)

	credentialID := uuid.New()
	params := sqlc.CreateCredentialParams{
		CredentialID: credentialID,
		UserID:       userID,
		Type:         credentials.PasswordType,
		IsPrimary:    true,
		IsVerified:   true,
		SecretData:   []byte(`"some-encrypted-secret"`),
	}

	_, err = usersRepo.CreateCredential(ctx, params)
	require.NoError(t, err)

	cred, err := usersRepo.GetCredentialByID(ctx, credentialID)
	require.NoError(t, err)
	require.NotNil(t, cred)
	require.Equal(t, credentialID, cred.CredentialID)
}

func TestGetUserCredentials(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	usersRepo := users_repo.NewUsersRepository(repo, pool)

	helpers.TruncateAllTables(t, pool)

	ctx := context.Background()
	userID := uuid.New()
	_, err := usersRepo.CreateUser(ctx, sqlc.CreateUserParams{
		UserID:  userID,
		Email:   "multi@example.com",
		Name:    "Multi",
		Surname: "Creds",
	})
	require.NoError(t, err)

	creds := []sqlc.CreateCredentialParams{
		{
			CredentialID: uuid.New(),
			UserID:       userID,
			Type:         credentials.PasswordType,
			IsPrimary:    true,
			IsVerified:   true,
			SecretData:   []byte(`"secret1"`),
		},
		{
			CredentialID: uuid.New(),
			UserID:       userID,
			Type:         credentials.TotpType,
			IsPrimary:    true,
			IsVerified:   true,
			SecretData:   []byte(`"secret2"`),
		},
	}

	for _, c := range creds {
		_, err := usersRepo.CreateCredential(ctx, c)
		require.NoError(t, err)
	}

	results, err := usersRepo.GetUserCredentials(ctx, userID)
	require.NoError(t, err)
	require.Len(t, results, len(creds))
}

func TestGetUserCredentialByType(t *testing.T) {
	_, pool, repo := helpers.GetDatabaseContainer(t)
	usersRepo := users_repo.NewUsersRepository(repo, pool)

	helpers.TruncateAllTables(t, pool)

	ctx := context.Background()
	userID := uuid.New()
	_, err := usersRepo.CreateUser(ctx, sqlc.CreateUserParams{
		UserID:  userID,
		Email:   "typed@example.com",
		Name:    "Typed",
		Surname: "Cred",
	})
	require.NoError(t, err)

	params := sqlc.CreateCredentialParams{
		CredentialID: uuid.New(),
		UserID:       userID,
		Type:         credentials.TotpType,
		IsPrimary:    true,
		IsVerified:   true,
		SecretData:   []byte(`"typed-secret"`),
	}

	_, err = usersRepo.CreateCredential(ctx, params)
	require.NoError(t, err)

	cred, err := usersRepo.GetUserCredentialByType(ctx, userID, credentials.TotpType)
	require.NoError(t, err)
	require.NotNil(t, cred)
	require.Equal(t, credentials.TotpType, cred.Type)
}
