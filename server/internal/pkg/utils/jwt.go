package utils

import (
	"crypto/rand"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

var jwtKey []byte

func init() {
	jwtKey = make([]byte, 32) // 生成32字节的密钥
	if _, err := rand.Read(jwtKey); err != nil {
		panic(err)
	}
}

type JwtCustomClaims struct {
	ID string
	jwt.RegisteredClaims
}

func GenerateJwt(id string) (string, error) {
	now := time.Now().Local()
	claims := JwtCustomClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "codepzj",
			Subject:   "codepzj",
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseJwt(jwtStr string) (JwtCustomClaims, error) {
	claims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(jwtStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return JwtCustomClaims{}, err
	}
	if !token.Valid {
		return JwtCustomClaims{}, errors.New("无效的token")
	}
	return claims, nil
}
