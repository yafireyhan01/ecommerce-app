package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Guid  string `json:"guid"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(guid, email, role string) (string, error) {
	now := time.Now()
	expTime := now.Add(time.Hour * 12)

	claims := &Claims{
		Guid:  guid,
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "synapsis-test",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
