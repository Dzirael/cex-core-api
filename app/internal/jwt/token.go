package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token is expired")
)

type JWTClaims interface {
	jwt.Claims
	SetRegisteredClaims(claims jwt.RegisteredClaims)
}

func VerifyToken[T any](tokenString string, secretKey []byte) (*T, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method is HMAC (common for JWT)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	// Handle parsing errors
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenUsedBeforeIssued) {
			return nil, ErrExpiredToken
		}
		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return nil, fmt.Errorf("get expiration time: %v", err)
	}

	if time.Now().After(exp.Time) {
		return nil, ErrExpiredToken
	}

	// Marshal claims map to JSON bytes
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return nil, fmt.Errorf("marshal claims: %v", err)
	}

	var out T
	// Unmarshal JSON bytes into generic output struct
	if err := json.Unmarshal(claimsJSON, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func CreateToken[T JWTClaims](claims T, secretKey []byte, ttl time.Duration) (string, error) {
	now := time.Now()
	claims.SetRegisteredClaims(jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
	})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
