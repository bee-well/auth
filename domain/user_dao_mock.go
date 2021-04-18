package domain

type UserDaoMock struct {
	SaveMock         func(*User) error
	SaveCalls        int
	InsertMock       func(*User) error
	InsertCalls      int
	FindByIDMock     func(string) (User, error)
	FindByIDCalls    int
	FindByEmailMock  func(string) (User, error)
	FindByEmailCalls int
}

func MockUserDao(u *UserDaoMock) {
	mock = u
}

func (mock *UserDaoMock) Save(u *User) error {
	mock.SaveCalls++
	return mock.SaveMock(u)
}

func (mock *UserDaoMock) Insert(u *User) error {
	mock.InsertCalls++
	return mock.InsertMock(u)
}

func (mock *UserDaoMock) FindByID(id string) (User, error) {
	mock.FindByIDCalls++
	return mock.FindByIDMock(id)
}

func (mock *UserDaoMock) FindByEmail(email string) (User, error) {
	mock.FindByEmailCalls++
	return mock.FindByEmailMock(email)
}
