package vault

import (
	"cex-core-api/app/config"
	"cex-core-api/app/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/vault/api"
)

const (
	TransitPasswordKey = "passwords-key"
)

type Engine struct {
	client *api.Client

	path      string
	rootToken string

	roleID   string
	secretID string

	ttl            time.Duration
	tokenValidTill time.Time
}

type Client struct {
	engines map[string]*Engine
}

func New(cfg *config.Config) (*Client, error) {
	apiCfg := api.DefaultConfig()
	apiCfg.Address = cfg.VaultConfig.Endpoint

	engines := make(map[string]*Engine)
	for name, engineCfg := range cfg.VaultConfig.Engines {

		vaultClient, err := api.NewClient(apiCfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create vault client: %w", err)
		}

		engine := &Engine{
			client:    vaultClient,
			path:      engineCfg.Path,
			roleID:    engineCfg.RoleID,
			secretID:  engineCfg.SecretID,
			ttl:       engineCfg.TTL,
			rootToken: cfg.VaultConfig.RootToken,
		}
		if err := engine.authorize(); err != nil {
			return nil, fmt.Errorf("failed to authorize engine '%s': %w", name, err)
		}
		engines[name] = engine
	}

	return &Client{
		engines: engines,
	}, nil
}

func (v *Engine) authorize() error {
	if v.rootToken != "" {
		v.client.SetToken(v.rootToken)
		return nil
	}

	if v.tokenValidTill.After(time.Now()) {
		return nil
	}

	resp, err := v.client.Logical().Write("auth/approle/login", map[string]interface{}{
		"role_id":   v.roleID,
		"secret_id": v.secretID,
	})

	if err != nil {
		return fmt.Errorf("approle login failed: %w", err)
	}
	if resp == nil || resp.Auth == nil || resp.Auth.ClientToken == "" {
		return fmt.Errorf("invalid approle login response")
	}

	v.client.SetToken(resp.Auth.ClientToken)
	v.tokenValidTill = time.Now().Add(v.ttl)
	return nil
}

func (c *Client) GetEngine(name string) (*Engine, error) {
	engine, ok := c.engines[name]
	if !ok {
		return nil, fmt.Errorf("engine '%s' not registered", name)
	}
	return engine, nil
}

type EncodeRequest struct {
	InputBase64 string  `json:"input"`
	Algorithm   *string `json:"algorithm,omitempty"`
	KeyVersion  *int    `json:"key_version,omitempty"`
}
type HmacResult struct {
	HMAC string `json:"hmac"`
}

func (v *Client) EncodeHmac(ctx context.Context, params EncodeRequest) (*HmacResult, error) {
	engine, err := v.GetEngine("transit")
	if err != nil {
		return nil, err
	}

	if err := engine.authorize(); err != nil {
		return nil, fmt.Errorf("failed to check authorize: %w", err)
	}

	logicClient := engine.client.Logical()
	path := fmt.Sprintf("%s/%s/%s", engine.path, "hmac", TransitPasswordKey)

	if params.Algorithm != nil {
		path += "/" + *params.Algorithm
	}

	data, _ := json.Marshal(&params)

	raw, err := logicClient.WriteRawWithContext(ctx, path, data)
	if err != nil {
		return nil, fmt.Errorf("hmac encode at path %q: %w", path, err)
	}

	out, err := utils.UnmarhalRawResponse[VaultResponse[HmacResult]](raw.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	return &out.Data, nil
}
