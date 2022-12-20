package mocks

import (
	"github/beomsun1234/go-restapi/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (mu *MockUserRepository) FindUserById(id int) (*models.User, error) {
	ret := mu.Mock.Called(id)

	var mock_user *models.User
	if ret.Get(0) != nil {
		mock_user = ret.Get(0).(*models.User)
	}

	var err error
	if ret.Get(1) != nil {
		err = ret.Get(1).(error)
	}
	return mock_user, err
}
func (mu *MockUserRepository) FindUserByName(name string) (*models.User, error) {
	ret := mu.Mock.Called(name)

	var mock_user *models.User
	if ret.Get(0) != nil {
		mock_user = ret.Get(0).(*models.User)
	}

	var err error
	if ret.Get(1) != nil {
		err = ret.Get(1).(error)
	}
	return mock_user, err

}
func (mu *MockUserRepository) CreateUser(u *models.User) (*models.User, error) {
	ret := mu.Mock.Called(u)

	var mock_user *models.User
	if ret.Get(0) != nil {
		mock_user = ret.Get(0).(*models.User)
	}

	var err error
	if ret.Get(1) != nil {
		err = ret.Get(1).(error)
	}
	return mock_user, err
}

func (mu *MockUserRepository) FindAllUsers() ([]*models.User, error) {

	ret := mu.Mock.Called()

	var mock_user []*models.User
	if ret.Get(0) != nil {
		mock_user = ret.Get(0).([]*models.User)
	}

	var err error
	if ret.Get(1) != nil {
		err = ret.Get(1).(error)
	}
	return mock_user, err

}
