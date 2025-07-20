package helpers

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	vaultOnce      sync.Once
	vaultContainer testcontainers.Container
)

func GetVaultContainer(t *testing.T) testcontainers.Container {
	t.Helper()

	vaultOnce.Do(func() {
		ctx := context.Background()

		const (
			image     = "hashicorp/vault:latest"
			port      = "8200/tcp"
			token     = "root"
			waitPath  = "/v1/sys/health"
			startupTO = 10 * time.Second
		)

		req := testcontainers.ContainerRequest{
			Image:        image,
			ExposedPorts: []string{port},
			Env: map[string]string{
				"VAULT_DEV_ROOT_TOKEN_ID":  token,
				"VAULT_DEV_LISTEN_ADDRESS": "0.0.0.0:8200",
			},
			WaitingFor: wait.ForHTTP(waitPath).
				WithPort(port).
				WithStartupTimeout(startupTO),
		}

		container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
		require.NoError(t, err, "failed to start container")

		vaultContainer = container

		host, err := container.Host(ctx)
		require.NoError(t, err, "failed to get container host")

		mappedPort, err := container.MappedPort(ctx, port)
		require.NoError(t, err, "failed to get container port")

		vaultAddr = fmt.Sprintf("http://%s:%s", host, mappedPort.Port())

		client, err := api.NewClient(&api.Config{Address: vaultAddr})
		require.NoError(t, err, "failed to create Vault client")
		client.SetToken(token)

		mounts := []struct {
			Path string
			Type string
		}{
			{"transit", "transit"},
			{"totp", "totp"},
		}
		for _, m := range mounts {
			if err := client.Sys().Mount(m.Path, &api.MountInput{Type: m.Type}); err != nil {
				require.NoError(t, err, "failed to mount engine")
			}
		}

		_, err = client.Logical().Write("transit/keys/passwords-key", map[string]interface{}{
			"type": "aes256-gcm96",
		})
		require.NoError(t, err, "failed to add passwords-key")
	})

	return vaultContainer
}
