package services

import (
	"errors"
	"time"

	"github.com/bee-well/auth/domain"
	"golang.org/x/crypto/bcrypt"
)

// SignIn ensures that the user credentials are correct and that
// an authorization token is generated and sent back to the caller
func SignIn(email, password string) (string, error) {
	dao, err := domain.NewUserDao()
	if err != nil {
		return "", err
	}

	user, err := dao.FindByEmail(email)
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("Invalid username or password")
	}

	t := domain.Token{
		ID:     user.ID,
		Issued: time.Now(),
	}

	return CreateJwt(&t, "some_secret")
}
