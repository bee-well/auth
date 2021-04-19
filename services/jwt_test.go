package services

import (
	"testing"
	"time"

	"github.com/bee-well/auth/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateJwt(t *testing.T) {
	time, _ := time.Parse("2006-01-02", "2006-01-02")
	token := domain.Token{
		ID:     1234,
		Issued: time,
	}

	jwt, err := CreateJwt(&token, "secret")
	assert.Nil(t, err)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIzNCwiaXNzdWVkIjoiMjAwNi0wMS0wMiJ9.k2ny-xP3dG3eQX_rFOufmdte3piSVYrOebKpzcP8IzQ", jwt)
}
