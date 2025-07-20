package users_repo

import (
	"cex-core-api/app/internal/models"
	"cex-core-api/app/internal/models/credentials"
	"cex-core-api/app/internal/storages/postgres/sqlc"
)

func userToModel(data sqlc.User) *models.User {
	return &models.User{
		UserID:    data.UserID,
		Email:     data.Email,
		Name:      data.Name,
		Surname:   data.Surname,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}

func credentialToModel(data sqlc.Credential) *credentials.UserCredentials {
	return &credentials.UserCredentials{
		CredentialID: data.CredentialID,
		UserID:       data.UserID,
		Type:         data.Type,
		IsPrimary:    data.IsPrimary,
		IsVerified:   data.IsVerified,
		Identifier:   data.Identifier,
		SecretData:   data.SecretData,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}
