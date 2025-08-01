// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"database/sql/driver"
	"fmt"
	"time"

	"cex-core-api/app/internal/models"
	"cex-core-api/app/internal/models/credentials"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AccountType string

const (
	AccountTypeSpot    AccountType = "spot"
	AccountTypeMargin  AccountType = "margin"
	AccountTypeFutures AccountType = "futures"
)

func (e *AccountType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccountType(s)
	case string:
		*e = AccountType(s)
	default:
		return fmt.Errorf("unsupported scan type for AccountType: %T", src)
	}
	return nil
}

type NullAccountType struct {
	AccountType AccountType
	Valid       bool // Valid is true if AccountType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccountType) Scan(value interface{}) error {
	if value == nil {
		ns.AccountType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AccountType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccountType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AccountType), nil
}

func (e AccountType) Valid() bool {
	switch e {
	case AccountTypeSpot,
		AccountTypeMargin,
		AccountTypeFutures:
		return true
	}
	return false
}

func AllAccountTypeValues() []AccountType {
	return []AccountType{
		AccountTypeSpot,
		AccountTypeMargin,
		AccountTypeFutures,
	}
}

type ChangeAction string

const (
	ChangeActionOrder    ChangeAction = "order"
	ChangeActionTransfer ChangeAction = "transfer"
	ChangeActionWithdraw ChangeAction = "withdraw"
	ChangeActionDeposit  ChangeAction = "deposit"
)

func (e *ChangeAction) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ChangeAction(s)
	case string:
		*e = ChangeAction(s)
	default:
		return fmt.Errorf("unsupported scan type for ChangeAction: %T", src)
	}
	return nil
}

type NullChangeAction struct {
	ChangeAction ChangeAction
	Valid        bool // Valid is true if ChangeAction is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullChangeAction) Scan(value interface{}) error {
	if value == nil {
		ns.ChangeAction, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ChangeAction.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullChangeAction) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ChangeAction), nil
}

func (e ChangeAction) Valid() bool {
	switch e {
	case ChangeActionOrder,
		ChangeActionTransfer,
		ChangeActionWithdraw,
		ChangeActionDeposit:
		return true
	}
	return false
}

func AllChangeActionValues() []ChangeAction {
	return []ChangeAction{
		ChangeActionOrder,
		ChangeActionTransfer,
		ChangeActionWithdraw,
		ChangeActionDeposit,
	}
}

type ChangeStatus string

const (
	ChangeStatusCreated   ChangeStatus = "created"
	ChangeStatusPending   ChangeStatus = "pending"
	ChangeStatusCancelled ChangeStatus = "cancelled"
	ChangeStatusCompleted ChangeStatus = "completed"
	ChangeStatusFailed    ChangeStatus = "failed"
)

func (e *ChangeStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ChangeStatus(s)
	case string:
		*e = ChangeStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ChangeStatus: %T", src)
	}
	return nil
}

type NullChangeStatus struct {
	ChangeStatus ChangeStatus
	Valid        bool // Valid is true if ChangeStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullChangeStatus) Scan(value interface{}) error {
	if value == nil {
		ns.ChangeStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ChangeStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullChangeStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ChangeStatus), nil
}

func (e ChangeStatus) Valid() bool {
	switch e {
	case ChangeStatusCreated,
		ChangeStatusPending,
		ChangeStatusCancelled,
		ChangeStatusCompleted,
		ChangeStatusFailed:
		return true
	}
	return false
}

func AllChangeStatusValues() []ChangeStatus {
	return []ChangeStatus{
		ChangeStatusCreated,
		ChangeStatusPending,
		ChangeStatusCancelled,
		ChangeStatusCompleted,
		ChangeStatusFailed,
	}
}

type ChangeType string

const (
	ChangeTypeReduce   ChangeType = "reduce"
	ChangeTypeIncrease ChangeType = "increase"
)

func (e *ChangeType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ChangeType(s)
	case string:
		*e = ChangeType(s)
	default:
		return fmt.Errorf("unsupported scan type for ChangeType: %T", src)
	}
	return nil
}

type NullChangeType struct {
	ChangeType ChangeType
	Valid      bool // Valid is true if ChangeType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullChangeType) Scan(value interface{}) error {
	if value == nil {
		ns.ChangeType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ChangeType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullChangeType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ChangeType), nil
}

func (e ChangeType) Valid() bool {
	switch e {
	case ChangeTypeReduce,
		ChangeTypeIncrease:
		return true
	}
	return false
}

func AllChangeTypeValues() []ChangeType {
	return []ChangeType{
		ChangeTypeReduce,
		ChangeTypeIncrease,
	}
}

type CredentialsType string

const (
	CredentialsTypePassword CredentialsType = "password"
	CredentialsTypeTotp     CredentialsType = "totp"
	CredentialsTypeWebauthn CredentialsType = "webauthn"
	CredentialsTypePasskey  CredentialsType = "passkey"
	CredentialsTypePhoneOtp CredentialsType = "phone_otp"
)

func (e *CredentialsType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CredentialsType(s)
	case string:
		*e = CredentialsType(s)
	default:
		return fmt.Errorf("unsupported scan type for CredentialsType: %T", src)
	}
	return nil
}

type NullCredentialsType struct {
	CredentialsType CredentialsType
	Valid           bool // Valid is true if CredentialsType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCredentialsType) Scan(value interface{}) error {
	if value == nil {
		ns.CredentialsType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CredentialsType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCredentialsType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CredentialsType), nil
}

func (e CredentialsType) Valid() bool {
	switch e {
	case CredentialsTypePassword,
		CredentialsTypeTotp,
		CredentialsTypeWebauthn,
		CredentialsTypePasskey,
		CredentialsTypePhoneOtp:
		return true
	}
	return false
}

func AllCredentialsTypeValues() []CredentialsType {
	return []CredentialsType{
		CredentialsTypePassword,
		CredentialsTypeTotp,
		CredentialsTypeWebauthn,
		CredentialsTypePasskey,
		CredentialsTypePhoneOtp,
	}
}

type OrderMethod string

const (
	OrderMethodFOK OrderMethod = "FOK"
	OrderMethodIOK OrderMethod = "IOK"
	OrderMethodGTC OrderMethod = "GTC"
)

func (e *OrderMethod) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrderMethod(s)
	case string:
		*e = OrderMethod(s)
	default:
		return fmt.Errorf("unsupported scan type for OrderMethod: %T", src)
	}
	return nil
}

type NullOrderMethod struct {
	OrderMethod OrderMethod
	Valid       bool // Valid is true if OrderMethod is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderMethod) Scan(value interface{}) error {
	if value == nil {
		ns.OrderMethod, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrderMethod.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderMethod) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrderMethod), nil
}

func (e OrderMethod) Valid() bool {
	switch e {
	case OrderMethodFOK,
		OrderMethodIOK,
		OrderMethodGTC:
		return true
	}
	return false
}

func AllOrderMethodValues() []OrderMethod {
	return []OrderMethod{
		OrderMethodFOK,
		OrderMethodIOK,
		OrderMethodGTC,
	}
}

type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

func (e *OrderSide) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrderSide(s)
	case string:
		*e = OrderSide(s)
	default:
		return fmt.Errorf("unsupported scan type for OrderSide: %T", src)
	}
	return nil
}

type NullOrderSide struct {
	OrderSide OrderSide
	Valid     bool // Valid is true if OrderSide is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderSide) Scan(value interface{}) error {
	if value == nil {
		ns.OrderSide, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrderSide.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderSide) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrderSide), nil
}

func (e OrderSide) Valid() bool {
	switch e {
	case OrderSideBuy,
		OrderSideSell:
		return true
	}
	return false
}

func AllOrderSideValues() []OrderSide {
	return []OrderSide{
		OrderSideBuy,
		OrderSideSell,
	}
}

type OrderStatus string

const (
	OrderStatusCreated         OrderStatus = "created"
	OrderStatusPartiallyFilled OrderStatus = "partially_filled"
	OrderStatusFilled          OrderStatus = "filled"
	OrderStatusCancelled       OrderStatus = "cancelled"
)

func (e *OrderStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrderStatus(s)
	case string:
		*e = OrderStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for OrderStatus: %T", src)
	}
	return nil
}

type NullOrderStatus struct {
	OrderStatus OrderStatus
	Valid       bool // Valid is true if OrderStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderStatus) Scan(value interface{}) error {
	if value == nil {
		ns.OrderStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrderStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrderStatus), nil
}

func (e OrderStatus) Valid() bool {
	switch e {
	case OrderStatusCreated,
		OrderStatusPartiallyFilled,
		OrderStatusFilled,
		OrderStatusCancelled:
		return true
	}
	return false
}

func AllOrderStatusValues() []OrderStatus {
	return []OrderStatus{
		OrderStatusCreated,
		OrderStatusPartiallyFilled,
		OrderStatusFilled,
		OrderStatusCancelled,
	}
}

type OrderType string

const (
	OrderTypeMarket OrderType = "market"
	OrderTypeLimit  OrderType = "limit"
)

func (e *OrderType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrderType(s)
	case string:
		*e = OrderType(s)
	default:
		return fmt.Errorf("unsupported scan type for OrderType: %T", src)
	}
	return nil
}

type NullOrderType struct {
	OrderType OrderType
	Valid     bool // Valid is true if OrderType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderType) Scan(value interface{}) error {
	if value == nil {
		ns.OrderType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrderType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrderType), nil
}

func (e OrderType) Valid() bool {
	switch e {
	case OrderTypeMarket,
		OrderTypeLimit:
		return true
	}
	return false
}

func AllOrderTypeValues() []OrderType {
	return []OrderType{
		OrderTypeMarket,
		OrderTypeLimit,
	}
}

type Account struct {
	AccountID uuid.UUID          `db:"account_id"`
	UserID    uuid.UUID          `db:"user_id"`
	Type      models.AccountType `db:"type"`
	CreatedAt time.Time          `db:"created_at"`
	UpdatedAt time.Time          `db:"updated_at"`
	DeletedAt *time.Time         `db:"deleted_at"`
}

type AccountBalance struct {
	BalanceID    uuid.UUID       `db:"balance_id"`
	AccountID    uuid.UUID       `db:"account_id"`
	TokenID      uuid.UUID       `db:"token_id"`
	Amount       decimal.Decimal `db:"amount"`
	LockedAmount decimal.Decimal `db:"locked_amount"`
	CreatedAt    time.Time       `db:"created_at"`
	UpdatedAt    time.Time       `db:"updated_at"`
	DeletedAt    *time.Time      `db:"deleted_at"`
}

type AccountBalanceChange struct {
	ChangeID  uuid.UUID           `db:"change_id"`
	AccountID uuid.UUID           `db:"account_id"`
	TokenID   uuid.UUID           `db:"token_id"`
	Type      models.ChangeType   `db:"type"`
	Action    models.ChangeAction `db:"action"`
	Status    models.ChangeStatus `db:"status"`
	Amount    decimal.Decimal     `db:"amount"`
	Sender    string              `db:"sender"`
	Recipient string              `db:"recipient"`
	CreatedAt time.Time           `db:"created_at"`
	UpdatedAt time.Time           `db:"updated_at"`
	DeletedAt *time.Time          `db:"deleted_at"`
}

type Chain struct {
	ChainID   uuid.UUID  `db:"chain_id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type Credential struct {
	CredentialID uuid.UUID        `db:"credential_id"`
	UserID       uuid.UUID        `db:"user_id"`
	Type         credentials.Type `db:"type"`
	IsPrimary    bool             `db:"is_primary"`
	IsVerified   bool             `db:"is_verified"`
	Identifier   *string          `db:"identifier"`
	SecretData   []byte           `db:"secret_data"`
	CreatedAt    time.Time        `db:"created_at"`
	UpdatedAt    time.Time        `db:"updated_at"`
}

type Market struct {
	MarketID       uuid.UUID       `db:"market_id"`
	TokenAID       uuid.UUID       `db:"token_a_id"`
	TokenBID       uuid.UUID       `db:"token_b_id"`
	IsActive       bool            `db:"is_active"`
	MinOrderAmount decimal.Decimal `db:"min_order_amount"`
	UpdatedAt      time.Time       `db:"updated_at"`
	CreatedAt      time.Time       `db:"created_at"`
	StartedAt      time.Time       `db:"started_at"`
	DeletedAt      *time.Time      `db:"deleted_at"`
}

type Order struct {
	OrderID      uuid.UUID       `db:"order_id"`
	AccountID    uuid.UUID       `db:"account_id"`
	MarketID     uuid.UUID       `db:"market_id"`
	Side         OrderSide       `db:"side"`
	Type         OrderType       `db:"type"`
	Method       OrderMethod     `db:"method"`
	Amount       decimal.Decimal `db:"amount"`
	AmountFilled decimal.Decimal `db:"amount_filled"`
	Status       OrderStatus     `db:"status"`
	Price        decimal.Decimal `db:"price"`
	ExpiresAt    time.Time       `db:"expires_at"`
	CompletedAt  *time.Time      `db:"completed_at"`
	CreatedAt    time.Time       `db:"created_at"`
	UpdatedAt    time.Time       `db:"updated_at"`
	DeletedAt    *time.Time      `db:"deleted_at"`
}

type OrderTrade struct {
	TradeID      uuid.UUID       `db:"trade_id"`
	OrderID      uuid.UUID       `db:"order_id"`
	AccountAID   uuid.UUID       `db:"account_a_id"`
	AccountBID   uuid.UUID       `db:"account_b_id"`
	AmountFilled decimal.Decimal `db:"amount_filled"`
	CreatedAt    time.Time       `db:"created_at"`
	UpdatedAt    time.Time       `db:"updated_at"`
	DeletedAt    *time.Time      `db:"deleted_at"`
}

type Token struct {
	TokenID   uuid.UUID  `db:"token_id"`
	IsNative  bool       `db:"is_native"`
	Name      string     `db:"name"`
	Symbol    string     `db:"symbol"`
	Decimals  int32      `db:"decimals"`
	LogoPath  *string    `db:"logo_path"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type TokenChain struct {
	TokenID uuid.UUID `db:"token_id"`
	ChainID uuid.UUID `db:"chain_id"`
}

type User struct {
	UserID    uuid.UUID  `db:"user_id"`
	Email     string     `db:"email"`
	Name      string     `db:"name"`
	Surname   string     `db:"surname"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
