package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type VaultEngineConfig struct {
	Path     string        `env:"PATH,required"`
	RoleID   string        `env:"ROLE_ID,required"`
	SecretID string        `env:"SECRET_ID,required"`
	TTL      time.Duration `env:"TTL,required"`
}

type VaultConfig struct {
	Endpoint    string   `env:"VAULT_ENDPOINT,required"`
	RootToken   string   `env:"VAULT_ROOT"`
	EngineNames []string `env:"VAULT_ENGINES" envSeparator:","`
	Engines     map[string]VaultEngineConfig
}

type Config struct {
	VaultConfig
}

func New() (*Config, error) {
	_ = godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	cfg.Engines = make(map[string]VaultEngineConfig)
	for _, name := range cfg.EngineNames {
		engineCfg := VaultEngineConfig{}
		prefix := fmt.Sprintf("ENGINE_%s_", toEnvKey(name))

		if err := env.ParseWithOptions(&engineCfg, env.Options{
			Prefix: prefix,
		}); err != nil {
			return nil, fmt.Errorf("parse engine config (%s): %w", name, err)
		}

		cfg.Engines[name] = engineCfg
	}

	return &cfg, nil
}

func toEnvKey(name string) string {
	return strings.ToUpper(strings.ReplaceAll(name, "-", "_"))
}
