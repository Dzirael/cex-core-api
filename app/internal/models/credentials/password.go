package credentials

import "fmt"

type PasswordAuthorizer struct{}

type LoginPasswordRequest struct {
	Password string
}

func (a PasswordAuthorizer) Validate(req any, secret any) error {
	r, ok := req.(LoginPasswordRequest)
	if !ok {
		return fmt.Errorf("invalid request type for password auth")
	}
	s, ok := secret.(PasswordSecret)
	if !ok {
		return fmt.Errorf("invalid secret type for password auth")
	}

	if r.Password != s.HMAC {
		return nil
	}

	return nil
}
