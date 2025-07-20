package credentials

import "fmt"

type TOTPAuthorizer struct{}

type TOTPLoginRequest struct {
	Code string
}

func (a TOTPAuthorizer) Validate(req any, secret any) error {
	r, ok := req.(TOTPLoginRequest)
	if !ok {
		return fmt.Errorf("invalid request type for TOTP auth")
	}
	s, ok := secret.(TOTPSecret)
	if !ok {
		return fmt.Errorf("invalid secret type for TOTP auth")
	}

	if r.Code != s.Secret {
		return fmt.Errorf("invalid totp code")
	}
	return nil
}
