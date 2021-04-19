package services

import (
	"errors"

	"github.com/bee-well/auth/domain"
	"github.com/bee-well/auth/mq"
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
		return errors.New("A user with that email is already registered.")
	}

	if err := mq.NewMq().Publish("users", mq.NewUserCreatedEvent(user).JsonBytes()); err != nil {
		// TODO: send email to support
		return errors.New("Could not complete your request, please contact customer support.")
	}

	return nil
}
