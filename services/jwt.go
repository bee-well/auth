package services

import (
	"fmt"

	"github.com/bee-well/auth/config"
	"github.com/bee-well/auth/domain"
	"github.com/dgrijalva/jwt-go"
)

var (
	NullToken = domain.Token{}
)

func CreateJwt(t *domain.Token, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = t.ID
	claims["issued"] = t.Issued.Format("2006-01-02T15:04:05Z07:00")
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return at.SignedString([]byte(secret))
}

func Verify(token string) (domain.Token, bool) {
	apiToken := domain.Token{}
	t, err := jwt.ParseWithClaims(token, &apiToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetString(config.JwtKey)), nil
	})

	if err != nil || !t.Valid {
		fmt.Println(err)
		fmt.Println(t.Valid)
		return NullToken, false
	}

	return apiToken, true
}
