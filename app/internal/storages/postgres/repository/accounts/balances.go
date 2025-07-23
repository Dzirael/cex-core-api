package accounts_repo

import (
	"context"
	"fmt"
	"slices"

	"cex-core-api/app/internal/models"
	"cex-core-api/app/internal/storages/postgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

func (r *AccountRepository) GetAccountBalances(ctx context.Context, params sqlc.GetTokenBalanceByAccountIDParams) ([]*models.AccountBalanceResult, error) {
	balances, err := r.repo.GetTokenBalanceByAccountID(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetAccountBalances: %w", err)
	}

	out := make([]*models.AccountBalanceResult, len(balances))
	for i, balance := range balances {
		out[i] = tokenBalanceToModel(balance)
	}
	return out, nil
}

type ChangeAccountBalanceParams struct {
	AccountID    uuid.UUID
	TokenID      uuid.UUID
	Type         models.ChangeType
	Action       models.ChangeAction
	Status       models.ChangeStatus
	ChangeAmount decimal.Decimal
	Sender       string
	Recipient    string
}

func (r *AccountRepository) UpdateAccountBalance(ctx context.Context, params ChangeAccountBalanceParams) (*models.UpdateAccountBalanceResult, error) {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := r.repo.WithTx(tx)
	balanceID := uuid.New()
	changeID := uuid.New()

	lockAmount := decimal.Zero
	if slices.Contains([]models.ChangeStatus{models.ChangeStatusCreated, models.ChangeStatusPending}, params.Status) {
		lockAmount = lockAmount.Add(params.ChangeAmount)
	}

	switch params.Type {

	case models.ChangeTypeIncrease:
		balanceID, err = qtx.IncreaseAccountBalance(ctx, sqlc.IncreaseAccountBalanceParams{
			BalanceID:            balanceID,
			AccountID:            params.AccountID,
			TokenID:              params.TokenID,
			IncreaseAmount:       params.ChangeAmount,
			IncreaseLockedAmount: lockAmount,
		})
		if err != nil {
			return nil, fmt.Errorf("sqlc: IncreaseAccountBalance: %w", err)
		}

	case models.ChangeTypeReduce:
		result, err := qtx.DecreaseAccountBalance(ctx, sqlc.DecreaseAccountBalanceParams{
			AccountID:      params.AccountID,
			TokenID:        params.TokenID,
			DecreaseAmount: params.ChangeAmount,
		})
		if err != nil {
			return nil, fmt.Errorf("sqlc: DecreaseAccountBalance: %w", err)
		}

		if !result.Success {
			return nil, fmt.Errorf("sqlc: infussient funds for transfer")
		}

		balanceID = result.BalanceID
	}

	err = qtx.CreateBalanceTransfer(ctx, sqlc.CreateBalanceTransferParams{
		ChangeID:  changeID,
		AccountID: params.AccountID,
		TokenID:   params.TokenID,
		Type:      models.ChangeTypeIncrease,
		Action:    params.Action,
		Status:    params.Status,
		Amount:    params.ChangeAmount,
		Sender:    params.Sender,
		Recipient: params.Recipient,
	})
	if err != nil {
		return nil, fmt.Errorf("sqlc: CreateBalanceTransfer: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("sqlc: commit transaction: %w", err)
	}

	return &models.UpdateAccountBalanceResult{
		BalanceID: balanceID,
		ChangeID:  changeID,
	}, nil
}

func (r *AccountRepository) GetBalanceChanges(ctx context.Context, params sqlc.GetBalanceChangesParams) ([]*models.AccountBalanceChange, error) {
	transfers, err := r.repo.GetBalanceChanges(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetBalanceChanges: %w", err)
	}

	out := make([]*models.AccountBalanceChange, len(transfers))
	for i, transfer := range transfers {
		out[i] = accountBalanceChangeToModel(transfer)
	}
	return out, nil
}

func (r *AccountRepository) GetBalanceChangeByID(ctx context.Context, changeID uuid.UUID) (*models.AccountBalanceChange, error) {
	transfer, err := r.repo.GetBalanceChangeByID(ctx, changeID)
	if err != nil {
		return nil, fmt.Errorf("sqlc: GetBalanceChangeByID: %w", err)
	}

	return accountBalanceChangeToModel(transfer), nil
}

func (r *AccountRepository) DecreaseAccountLockedBalance(ctx context.Context, balanceID uuid.UUID, amount decimal.Decimal) error {
	err := r.repo.DecreaseAccountLockedBalance(ctx, amount, balanceID)
	if err != nil {
		return fmt.Errorf("sqlc: DecreaseAccountLockedBalance: %w", err)
	}

	return nil
}
