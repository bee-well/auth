package domain

type UserDaoMock struct {
	SaveMock          func(*User) error
	SaveCalls         int
	InsertMock        func(*User) error
	InsertCalls       int
	FindByIDMock      func(int64) (User, error)
	FindByIDCalls     int
	FindByEmailMock   func(string) (User, error)
	FindByEmailCalls  int
	GetUserCountCalls int
	GetUserCountMock  func() (int, error)
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

func (mock *UserDaoMock) FindByID(id int64) (User, error) {
	mock.FindByIDCalls++
	return mock.FindByIDMock(id)
}

func (mock *UserDaoMock) FindByEmail(email string) (User, error) {
	mock.FindByEmailCalls++
	return mock.FindByEmailMock(email)
}

func (mock *UserDaoMock) GetUserCount() (int, error) {
	mock.GetUserCountCalls++
	return mock.GetUserCountMock()
}
