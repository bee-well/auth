package services

import (
	"errors"
	"fmt"

	"github.com/bee-well/auth/domain"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(
	email,
	password,
	firstName,
	lastName string,
) error {
	dao, err := domain.NewUserDao()
	if err != nil {
		return err
	}

	user := domain.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return errors.New("Could not hash password.")
	}

	user.Password = string(hash)

	if err := dao.Insert(&user); err != nil {
		fmt.Println(err.Error())
		return errors.New("A user with that email is already registered.")
	}

	return nil
}
