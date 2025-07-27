package crypto_repo

import (
	"cex-core-api/app/internal/models"
	"cex-core-api/app/internal/storages/postgres/sqlc"
)

func supportedTokenToModel(data sqlc.GetSupportedTokensRow) *models.SupportedTokensResult {
	return &models.SupportedTokensResult{
		TokenID:   data.TokenID,
		IsNative:  data.IsNative,
		Name:      data.Name,
		Symbol:    data.Symbol,
		Decimals:  uint(data.Decimals),
		LogoPath:  data.LogoPath,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		ChainIDs:  data.ChainIds,
	}
}

func tokenToModel(data sqlc.Token) *models.Token {
	return &models.Token{
		TokenID:   data.TokenID,
		IsNative:  data.IsNative,
		Name:      data.Name,
		Symbol:    data.Symbol,
		Decimals:  uint(data.Decimals),
		LogoPath:  data.LogoPath,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}

func chainToModel(data sqlc.Chain) *models.Chain {
	return &models.Chain{
		ChainID:   data.ChainID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}
