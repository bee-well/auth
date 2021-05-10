package services

import "github.com/bee-well/auth/domain"

func GetUserCount() (int, error) {
	dao, err := domain.NewUserDao()
	if err != nil {
		return -1, err
	}
	return dao.GetUserCount()
}
