package models

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	TokenID   uuid.UUID  `json:"token_id"`
	IsNative  bool       `json:"is_native"`
	Name      string     `json:"name"`
	Symbol    string     `json:"symbol"`
	Decimals  uint       `json:"decimals"`
	LogoPath  *string    `json:"logo_path"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Chain struct {
	ChainID   uuid.UUID  `json:"chain_id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type SupportedTokensResult struct {
	TokenID   uuid.UUID   `json:"token_id"`
	IsNative  bool        `json:"is_native"`
	Name      string      `json:"name"`
	Symbol    string      `json:"symbol"`
	Decimals  uint        `json:"decimals"`
	LogoPath  *string     `json:"logo_path"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at"`
	ChainIDs  []uuid.UUID `json:"chain_ids"`
}
