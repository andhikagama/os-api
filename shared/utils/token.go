package utils

import (
	"github.com/dgrijalva/jwt-go"
)

type (
	Token struct {
		UserID    uint64
		UserPhone string
		jwt.StandardClaims
	}
)

func GenerateToken(cc Token) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	tokenstring, err := token.SignedString([]byte(_CipherKey))
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func ClaimToken(tokenString string) (*Token, error) {
	var userToken *Token

	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(_CipherKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		userToken = claims
	}

	return userToken, nil
}
