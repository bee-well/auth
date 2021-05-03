package services

import "github.com/bee-well/auth/domain"

var NullUserPayload = UserPayload{}

type UserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func GetUser(id int64) (UserPayload, error) {
	dao, err := domain.NewUserDao()
	if err != nil {
		return NullUserPayload, err
	}

	user, err := dao.FindByID(id)
	if err != nil {
		return NullUserPayload, err
	}

	return UserPayload{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}
