package accounts_repo

import (
	"context"
	"fmt"

	"cex-core-api/app/internal/models"
	"cex-core-api/app/internal/storages/postgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

type ChangeTransferParams struct {
	AccountID uuid.UUID
	TokenID   uuid.UUID
	Amount    decimal.Decimal
	Status    models.ChangeStatus
	Recipient uuid.UUID
}

func (r *AccountRepository) CreateIntenalTransfer(ctx context.Context, params ChangeTransferParams) (*models.InternalTransferResult, error) {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := r.repo.WithTx(tx)
	fromBalanceID := uuid.New()
	toBalanceID := uuid.New()
	fromChangeID := uuid.New()
	toChangeID := uuid.New()

	result, err := qtx.DecreaseAccountBalance(ctx, sqlc.DecreaseAccountBalanceParams{
		AccountID:      params.AccountID,
		TokenID:        params.TokenID,
		DecreaseAmount: params.Amount,
	})
	if err != nil {
		return nil, fmt.Errorf("sqlc: DecreaseAccountBalance: %w", err)
	}

	if !result.Success {
		return nil, fmt.Errorf("sqlc: infussient funds for transfer")
	}
	fromBalanceID = result.BalanceID

	toBalanceID, err = qtx.IncreaseAccountBalance(ctx, sqlc.IncreaseAccountBalanceParams{
		BalanceID:      toBalanceID,
		AccountID:      params.Recipient,
		TokenID:        params.TokenID,
		IncreaseAmount: params.Amount,
	})
	if err != nil {
		return nil, fmt.Errorf("sqlc: IncreaseAccountBalance: %w", err)
	}

	err = qtx.CreateBalanceTransfer(ctx, sqlc.CreateBalanceTransferParams{
		ChangeID:  fromChangeID,
		AccountID: params.AccountID,
		TokenID:   params.TokenID,
		Type:      models.ChangeTypeReduce,
		Action:    models.ChangeActionTransfer,
		Status:    params.Status,
		Amount:    params.Amount,
		Sender:    params.AccountID.String(),
		Recipient: params.Recipient.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("sqlc: CreateBalanceTransfer: %w", err)
	}

	err = qtx.CreateBalanceTransfer(ctx, sqlc.CreateBalanceTransferParams{
		ChangeID:  toChangeID,
		AccountID: params.Recipient,
		TokenID:   params.TokenID,
		Type:      models.ChangeTypeIncrease,
		Action:    models.ChangeActionTransfer,
		Amount:    params.Amount,
		Status:    params.Status,
		Sender:    params.AccountID.String(),
		Recipient: params.Recipient.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("sqlc: CreateBalanceTransfer: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("sqlc: commit transaction: %w", err)
	}

	return &models.InternalTransferResult{
		FromBalanceID: fromBalanceID,
		FromChangeID:  fromChangeID,
		ToBalanceID:   toBalanceID,
		ToChangeID:    toChangeID,
	}, nil
}
