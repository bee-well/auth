package services

import (
	"github.com/bee-well/auth/domain"
	"github.com/dgrijalva/jwt-go"
)

func CreateJwt(t *domain.Token, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = t.ID
	claims["issued"] = t.Issued
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return at.SignedString([]byte(secret))
}
