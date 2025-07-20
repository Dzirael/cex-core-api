package credentials

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Type string

const (
	PasswordType Type = "password"
	TotpType     Type = "totp"
	WebAuthType  Type = "webauthn"
	PasskeyType  Type = "passkey"
	PhoneOTPType Type = "phone_otp"
)

type UserCredentials struct {
	CredentialID uuid.UUID
	UserID       uuid.UUID
	Type         Type
	IsPrimary    bool
	IsVerified   bool
	Identifier   *string
	SecretData   json.RawMessage
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (c *UserCredentials) DecodeSecret() (any, error) {
	switch c.Type {
	case PasswordType:
		var secret PasswordSecret
		if err := json.Unmarshal(c.SecretData, &secret); err != nil {
			return nil, fmt.Errorf("decode password secret: %w", err)
		}
		return &secret, nil
	case TotpType:
		var secret TOTPSecret
		if err := json.Unmarshal(c.SecretData, &secret); err != nil {
			return nil, fmt.Errorf("decode totp secret: %w", err)
		}
		return &secret, nil
	default:
		return nil, fmt.Errorf("unsupported credential type: %s", c.Type)
	}
}

func (c *UserCredentials) CanAuthorize(req any) error {
	var authorizers = map[Type]CredentialAuthorizer{
		"password": PasswordAuthorizer{},
		"totp":     TOTPAuthorizer{},
	}

	authorizer, ok := authorizers[c.Type]
	if !ok {
		return fmt.Errorf("unsupported credential type: %s", c.Type)
	}

	secret, err := c.DecodeSecret()
	if err != nil {
		return err
	}

	return authorizer.Validate(req, secret)
}

type CredentialAuthorizer interface {
	Validate(req any, secret any) error
}
