package credentials

type PasswordSecret struct {
	HMAC string `json:"hmac"`
	Salt string `json:"salt"`
}

type TOTPSecret struct {
	Secret string `json:"secret"`
	Issuer string `json:"issuer"`
}
