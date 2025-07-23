package vault_test

import (
	"crypto/rand"
	"encoding/base64"
	"testing"

	"cex-core-api/app/internal/storages/vault"
	"cex-core-api/app/pkg/utils"
	"cex-core-api/app/test/helpers"

	"github.com/stretchr/testify/require"
)

func TestCreateHmac(t *testing.T) {
	_ = helpers.GetVaultContainer(t)

	testCfg := helpers.TestConfig()

	vaultClient, err := vault.New(testCfg)
	require.NoError(t, err)

	password := "some_plain_password"
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	require.NoError(t, err)

	combined := append([]byte(password+":"), salt...)

	firstResult, err := vaultClient.EncodeHmac(t.Context(), vault.EncodeRequest{
		InputBase64: base64.StdEncoding.EncodeToString(combined),
		Algorithm:   utils.ToPointer("sha2-256"),
		KeyVersion:  utils.ToPointer(1),
	})
	require.NoError(t, err)

	secondResult, err := vaultClient.EncodeHmac(t.Context(), vault.EncodeRequest{
		InputBase64: base64.StdEncoding.EncodeToString(combined),
		Algorithm:   utils.ToPointer("sha2-256"),
		KeyVersion:  utils.ToPointer(1),
	})
	require.NoError(t, err)

	require.Equal(t, firstResult.HMAC, secondResult.HMAC)
}
