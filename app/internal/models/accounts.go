package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountType string

const (
	AccountTypeSpot    AccountType = "spot"
	AccountTypeFutures AccountType = "futures"
)

type ChangeType string

const (
	ChangeTypeReduce   ChangeType = "reduce"
	ChangeTypeIncrease ChangeType = "increase"
)

type ChangeAction string

const (
	ChangeActionOrder    ChangeAction = "order"
	ChangeActionTransfer ChangeAction = "transfer"
	ChangeActionDeposit  ChangeAction = "deposit"
	ChangeActionWithdraw ChangeAction = "withdraw"
)

type ChangeStatus string

const (
	ChangeStatusCreated   ChangeStatus = "order"
	ChangeStatusPending   ChangeStatus = "pending"
	ChangeStatusCancelled ChangeStatus = "cancelled"
	ChangeStatusCompleted ChangeStatus = "completed"
	ChangeStatusFailed    ChangeStatus = "failed"
)

type Account struct {
	AccountID uuid.UUID   `json:"account_id"`
	UserID    uuid.UUID   `json:"user_id"`
	Type      AccountType `json:"type"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at,omitempty"`
}

type AccountBalance struct {
	BalanceID    uuid.UUID       `json:"balance_id"`
	AccountID    uuid.UUID       `json:"account_id"`
	TokenID      uuid.UUID       `json:"token_id"`
	Amount       decimal.Decimal `json:"amount"`
	LockedAmount decimal.Decimal `json:"locked_amount"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    *time.Time      `json:"deleted_at,omitempty"`
}

type AccountBalanceChange struct {
	ChangeID  uuid.UUID       `json:"change_id"`
	AccountID uuid.UUID       `json:"account_id"`
	TokenID   uuid.UUID       `json:"token_id"`
	Type      ChangeType      `json:"type"`
	Action    ChangeAction    `json:"action"`
	Status    ChangeStatus    `json:"status"`
	Amount    decimal.Decimal `json:"amount"`
	Sender    string          `json:"sender"`
	Recipient string          `json:"recipient"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *time.Time      `json:"deleted_at,omitempty"`
}

// Aggregation from tokens and account balances table
type AccountBalanceResult struct {
	// Account balance Info
	BalanceID    uuid.UUID       `json:"balance_id"`
	Amount       decimal.Decimal `json:"amount"`
	LockedAmount decimal.Decimal `json:"locked_amount"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`

	// Token Info
	TokenID  uuid.UUID `json:"token_id"`
	IsNative bool      `json:"is_native"`
	Name     string    `json:"name"`
	Symbol   string    `json:"symbol"`
	Decimals uint      `json:"decimals"`
	LogoPath *string   `json:"logo_path,omitempty"`
}

type UpdateAccountBalanceResult struct {
	BalanceID uuid.UUID
	ChangeID  uuid.UUID
}

type InternalTransferResult struct {
	FromBalanceID uuid.UUID
	ToBalanceID   uuid.UUID
	FromChangeID  uuid.UUID
	ToChangeID    uuid.UUID
}
