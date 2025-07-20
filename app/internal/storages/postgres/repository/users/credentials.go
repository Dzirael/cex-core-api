package users_repo

import (
	"context"
	"fmt"

	"cex-core-api/app/internal/models/credentials"
	"cex-core-api/app/internal/storages/postgres/sqlc"

	"github.com/google/uuid"
)

func (r *UsersRepository) CreateCredential(ctx context.Context, params sqlc.CreateCredentialParams) (*credentials.UserCredentials, error) {
	err := r.repo.CreateCredential(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("sqlc: CreateCredential: %w", err)
	}

	return r.GetCredentialByID(ctx, params.CredentialID)
}

func (r *UsersRepository) GetCredentialByID(ctx context.Context, credentialID uuid.UUID) (*credentials.UserCredentials, error) {
	data, err := r.repo.GetCredentialByID(ctx, credentialID)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetCredentialByID: %w", err)
	}

	return credentialToModel(data), nil
}

func (r *UsersRepository) GetUserCredentials(ctx context.Context, userID uuid.UUID) ([]*credentials.UserCredentials, error) {
	data, err := r.repo.GetUserCredentials(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetUserCredentials: %w", err)
	}

	out := make([]*credentials.UserCredentials, len(data))
	for i, credential := range data {
		out[i] = credentialToModel(credential)
	}
	return out, nil
}

func (r *UsersRepository) GetUserCredentialByType(ctx context.Context, userID uuid.UUID, t credentials.Type) (*credentials.UserCredentials, error) {
	data, err := r.repo.GetUserCredentialByType(ctx, t, userID)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetUserCredentialByType: %w", err)
	}

	return credentialToModel(data), nil
}
