package helpers

import (
	"cex-core-api/app/config"
	"time"
)

func ToPointer[T any](v T) *T {
	return &v
}

var (
	vaultAddr = ""
)

func TestConfig() *config.Config {
	return &config.Config{
		VaultConfig: config.VaultConfig{
			Endpoint:    vaultAddr,
			RootToken:   "root",
			EngineNames: []string{"totp", "transit"},
			Engines: map[string]config.VaultEngineConfig{
				"totp": {
					Path:     "totp",
					RoleID:   "",
					SecretID: "",
					TTL:      time.Hour,
				},
				"transit": {
					Path:     "transit",
					RoleID:   "",
					SecretID: "",
					TTL:      time.Hour,
				},
			},
		},
	}
}
