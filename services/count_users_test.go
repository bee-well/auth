package services

import (
	"errors"
	"testing"

	"github.com/bee-well/auth/domain"
	"github.com/bee-well/auth/mq"
	"github.com/stretchr/testify/assert"
)

func TestCountUsers(t *testing.T) {
	daoMock := domain.UserDaoMock{
		GetUserCountMock: func() (int, error) {
			return 10, nil
		},
	}

	mqMock := mq.MqMock{
		PublishMock: func(string, []byte) error {
			return nil
		},
	}

	mq.MockMq(&mqMock)
	domain.MockUserDao(&daoMock)

	count, err := GetUserCount()

	assert.Equal(t, 1, daoMock.GetUserCountCalls, "expected exactly one FindByEmail call")
	assert.Nil(t, err)
	assert.Equal(t, 10, count)

	daoMock = domain.UserDaoMock{
		GetUserCountMock: func() (int, error) {
			return -1, errors.New("")
		},
	}

	count, err = GetUserCount()

	assert.Equal(t, 1, daoMock.GetUserCountCalls, "expected exactly one FindByEmail call")
	assert.NotNil(t, err)
	assert.Equal(t, -1, count)
}
