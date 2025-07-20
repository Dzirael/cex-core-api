package accounts_repo

import (
	"cex-core-api/app/internal/models"
	"cex-core-api/app/internal/storages/postgres/sqlc"
)

func accountToModel(data sqlc.Account) *models.Account {
	return &models.Account{
		AccountID: data.AccountID,
		UserID:    data.UserID,
		Type:      data.Type,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}

// func accountBalanceToModel(data sqlc.AccountBalance) *models.AccountBalance {
// 	return &models.AccountBalance{
// 		BalanceID:    data.BalanceID,
// 		AccountID:    data.AccountID,
// 		TokenID:      data.TokenID,
// 		Amount:       data.Amount,
// 		LockedAmount: data.LockedAmount,
// 		CreatedAt:    data.CreatedAt,
// 		UpdatedAt:    data.UpdatedAt,
// 		DeletedAt:    data.DeletedAt,
// 	}
// }

func tokenBalanceToModel(data sqlc.GetTokenBalanceByAccountIDRow) *models.AccountBalanceResult {
	return &models.AccountBalanceResult{
		BalanceID:    data.BalanceID,
		Amount:       data.Amount,
		LockedAmount: data.LockedAmount,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,

		TokenID:  data.TokenID,
		IsNative: data.IsNative,
		Name:     data.Name,
		Symbol:   data.Symbol,
		Decimals: uint(data.Decimals),
		LogoPath: data.LogoPath,
	}
}

func accountBalanceChangeToModel(data sqlc.AccountBalanceChange) *models.AccountBalanceChange {
	return &models.AccountBalanceChange{
		ChangeID:  data.ChangeID,
		AccountID: data.AccountID,
		TokenID:   data.TokenID,
		Type:      data.Type,
		Amount:    data.Amount,
		Action:    data.Action,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}

func accountChainTransferToModel(data sqlc.AccountBalanceChange) *models.AccountBalanceChange {
	return &models.AccountBalanceChange{
		ChangeID:  data.ChangeID,
		AccountID: data.AccountID,
		TokenID:   data.TokenID,
		Type:      data.Type,
		Amount:    data.Amount,
		Action:    data.Action,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}
