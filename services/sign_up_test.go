package services

import (
	"errors"
	"testing"

	"github.com/bee-well/auth/domain"
	"github.com/stretchr/testify/assert"
)

func TestSignUpDatabaseError(t *testing.T) {
	daoMock := domain.UserDaoMock{
		InsertMock: func(d *domain.User) error {
			return errors.New("")
		},
	}

	domain.MockUserDao(&daoMock)

	err := SignUp(
		"email",
		"password",
		"first name",
		"last name",
	)

	assert.NotNil(t, err, "expected error to be returned")
	assert.Equal(t, 1, daoMock.InsertCalls)
}

func TestSignUpOK(t *testing.T) {
	var u *domain.User
	daoMock := domain.UserDaoMock{
		InsertMock: func(d *domain.User) error {
			u = d
			return nil
		},
	}

	domain.MockUserDao(&daoMock)

	err := SignUp(
		"email",
		"password",
		"first name",
		"last name",
	)

	assert.Nil(t, err)
	assert.Equal(t, 1, daoMock.InsertCalls)
	assert.NotEqual(t, "password", u.Password, "expected password to be hashed")
}
