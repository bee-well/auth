package mq

import (
	"encoding/json"

	"github.com/bee-well/auth/domain"
)

const (
	typeCreated = "created"
)

type Event struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type UserPayload struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func NewUserCreatedEvent(u domain.User) Event {
	return Event{
		Type: typeCreated,
		Payload: UserPayload{
			ID:        u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		},
	}
}

func (e Event) JsonBytes() []byte {
	b, _ := json.Marshal(e)
	return b
}
