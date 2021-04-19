package services

import (
	"testing"

	"github.com/bee-well/auth/domain"
	"github.com/bee-well/auth/mq"
	"github.com/stretchr/testify/assert"
)

func TestSignInBadPassword(t *testing.T) {
	daoMock := domain.UserDaoMock{
		FindByEmailMock: func(string) (domain.User, error) {
			return domain.User{
				Email: "test@test.com",
				// password decoded is 'password'
				Password: "$2y$14$E41z9Wx8/dUXi.YRfn0zgeG3HSwJuI8El58eA6LypX8aomn0mGkhe",
			}, nil
		},
	}

	mqMock := mq.MqMock{
		PublishMock: func(string, []byte) error {
			return nil
		},
	}

	mq.MockMq(&mqMock)
	domain.MockUserDao(&daoMock)

	token, err := SignIn("test@test.com", "test")

	assert.Equal(t, 1, daoMock.FindByEmailCalls, "expected exactly one FindByEmail call")
	assert.NotNil(t, err, "expected error to be returned")
	assert.Equal(t, "", token, "expected empty token to be returned")
}

func TestSignInOK(t *testing.T) {
	daoMock := domain.UserDaoMock{
		FindByEmailMock: func(string) (domain.User, error) {
			return domain.User{
				Email: "test@test.com",
				// password decoded is 'password'
				Password: "$2y$14$E41z9Wx8/dUXi.YRfn0zgeG3HSwJuI8El58eA6LypX8aomn0mGkhe",
			}, nil
		},
	}

	mqMock := mq.MqMock{
		PublishMock: func(string, []byte) error {
			return nil
		},
	}

	mq.MockMq(&mqMock)
	domain.MockUserDao(&daoMock)

	token, err := SignIn("test@test.com", "password")

	assert.Equal(t, 1, daoMock.FindByEmailCalls, "expected exactly one FindByEmail call")
	assert.Nil(t, err, "no error should be returned")
	assert.NotEqual(t, "", token, "expected token")
}
