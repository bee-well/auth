package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/bee-well/auth/domain"
	"github.com/bee-well/auth/mq"
	"github.com/stretchr/testify/assert"
)

func TestSignUpDatabaseError(t *testing.T) {
	daoMock := domain.UserDaoMock{
		InsertMock: func(d *domain.User) error {
			return errors.New("")
		},
	}

	mqMock := mq.MqMock{
		PublishMock: func(q string, e []byte) error {
			return nil
		},
	}

	mq.MockMq(&mqMock)
	domain.MockUserDao(&daoMock)

	err := SignUp(
		"test@test.com",
		"password",
		"Test",
		"Testsson",
	)

	assert.Equal(t, 0, mqMock.PublishCalls)
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

	mqMock := mq.MqMock{
		PublishMock: func(q string, e []byte) error {
			var event mq.Event
			err := json.Unmarshal(e, &event)
			assert.Nil(t, err)
			assert.Equal(t, event.Type, "created")
			fmt.Printf("%v\n", event.Payload)
			payload, ok := event.Payload.(map[string]interface{})
			assert.True(t, ok)
			assert.Equal(t, "test@test.com", payload["email"].(string))
			assert.Equal(t, "Test", payload["firstName"].(string))
			assert.Equal(t, "Testsson", payload["lastName"].(string))
			return nil
		},
	}

	mq.MockMq(&mqMock)
	domain.MockUserDao(&daoMock)

	err := SignUp(
		"test@test.com",
		"password",
		"Test",
		"Testsson",
	)

	assert.Nil(t, err)
	assert.Equal(t, 1, mqMock.PublishCalls)
	assert.Equal(t, 1, daoMock.InsertCalls)
	assert.NotEqual(t, "password", u.Password, "expected password to be hashed")
}
