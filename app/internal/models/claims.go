package models

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	V      string    `json:"v"`

	data *ClaimsData `json:"-"`
	jwt.RegisteredClaims
}

func (c *Claims) SetRegisteredClaims(claims jwt.RegisteredClaims) {
	c.RegisteredClaims = claims
}

func (c *Claims) GetRegisteredClaims() *jwt.RegisteredClaims {
	return &c.RegisteredClaims
}

type ClaimsData struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

func NewClaimsData(name, surname, email string) *ClaimsData {
	return &ClaimsData{
		Name:    name,
		Surname: surname,
		Email:   email,
	}
}

func NewClaims(userID uuid.UUID, data *ClaimsData, secretKey [32]byte) (*Claims, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	encrypted, err := encryptAES(jsonData, secretKey)
	if err != nil {
		return nil, err
	}

	claims := &Claims{
		UserID: userID,
		V:      encrypted,
		data:   data,
	}

	return claims, nil
}

func encryptAES(plaintext []byte, key [32]byte) (string, error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()

	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	paddedText := append(plaintext, padtext...)

	ciphertext := make([]byte, blockSize+len(paddedText))
	iv := ciphertext[:blockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], paddedText)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
