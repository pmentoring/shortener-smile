package service

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTService struct {
	secretKey string
	publicKey string
}

type JWTClaims struct {
	sub string
	exp string
	iat string
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{secretKey: secret}
}

func (s JWTService) CreateJwtForId(id int) (string, error) {
	currTime := time.Now

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"sub": id,
		"exp": currTime().Add(3 * time.Hour).Format(time.RFC3339),
		"iat": currTime().Format(time.RFC3339),
	})

	block, _ := pem.Decode([]byte(s.secretKey))

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	return token.SignedString(key)
}

//func (s JWTService) IsJwtValid(tokenString string) (bool, error) {
//	jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
//
//	})
//}
