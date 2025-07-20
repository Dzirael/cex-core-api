package jwt_test

import (
	"testing"
	"time"

	"cex-core-api/app/internal/jwt"

	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwtlib.RegisteredClaims
}

func (c *CustomClaims) SetRegisteredClaims(claims jwtlib.RegisteredClaims) {
	c.RegisteredClaims = claims
}

func (c *CustomClaims) GetRegisteredClaims() *jwtlib.RegisteredClaims {
	return &c.RegisteredClaims
}

func TestCreateAndVerifyToken(t *testing.T) {
	secretKey := []byte("supersecret")

	claims := &CustomClaims{
		UserID: "user123",
	}

	tokenStr, err := jwt.CreateToken(claims, secretKey, time.Minute)
	assert.NoError(t, err, "expected valid token")
	assert.NotEmpty(t, tokenStr, "expected no zero length token")

	verifiedClaims, err := jwt.VerifyToken[*CustomClaims](tokenStr, secretKey)
	assert.NoError(t, err)
	assert.Equal(t, "user123", (*verifiedClaims).UserID)
}

func TestExpiredToken(t *testing.T) {
	secretKey := []byte("supersecret")

	claims := &CustomClaims{
		UserID: "user123",
	}

	tokenStr, err := jwt.CreateToken(claims, secretKey, -1*time.Minute)
	assert.NoError(t, err)

	_, err = jwt.VerifyToken[*CustomClaims](tokenStr, secretKey)
	assert.ErrorIs(t, err, jwt.ErrExpiredToken)
}

func TestInvalidSignature(t *testing.T) {
	claims := &CustomClaims{
		UserID: "user123",
	}

	tokenStr, err := jwt.CreateToken(claims, []byte("right-key"), time.Minute)
	require.NoError(t, err)

	_, err = jwt.VerifyToken[*CustomClaims](tokenStr, []byte("wrong-key"))
	require.ErrorIs(t, err, jwt.ErrInvalidToken)
}
