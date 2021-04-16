package domain

import (
	_ "github.com/lib/pq"
)

type UserDao interface {
	Save(*User) error
	Insert(*User) error
	FindByID(string) (User, error)
	FindByEmail(string) (User, error)
}

func NewUserDao() UserDao {
	return &userDao{}
}

type userDao struct{}

func (userDao) Save(u *User) error {
	panic("not implemented")
}

func (userDao) Insert(u *User) error {
	db, err := newSqlConnector().Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		`
		INSERT INTO users 
			(email, first_name, last_name, password)
		VALUES 
			($1, $2, $3, $4)
		`,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

func (userDao) FindByID(id string) (User, error) {
	db, err := newSqlConnector().Connect()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	var user User
	err = db.QueryRow("SELECT * FROM users WHERE id=$1 LIMIT 1", id).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
	)
	return user, err
}

func (userDao) FindByEmail(email string) (User, error) {
	db, err := newSqlConnector().Connect()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	var user User
	err = db.QueryRow("SELECT * FROM users WHERE email=$1 LIMIT 1", email).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
	)
	return user, err
}
