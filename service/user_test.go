package service

import (
	"errors"
	"github/beomsun1234/go-restapi/mocks"
	"github/beomsun1234/go-restapi/models"

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserService_FindUserById(t *testing.T) {
	t.Run("성공", func(t *testing.T) {
		find_id := 1
		mockUser := models.NewUser().BuildName("park").BuildId(1)

		mockUserRepo := new(mocks.MockUserRepository)
		us := NewUserService(mockUserRepo)

		mockUserRepo.On("FindUserById", find_id).Return(mockUser, nil)

		find_user, err := us.FindUserById(find_id)

		assert.NoError(t, err)
		assert.Equal(t, find_user.Name, mockUser.Name)
		assert.Equal(t, find_user.ID, find_id)
	})

	t.Run("실패", func(t *testing.T) {
		find_id := 10

		mockUserRepo := new(mocks.MockUserRepository)
		us := NewUserService(mockUserRepo)

		mockUserRepo.On("FindUserById", find_id).Return(nil, errors.New("not found"))

		find_user, err := us.FindUserById(find_id)

		assert.Error(t, err)
		assert.Nil(t, find_user)
	})
}

func Test_UserService_FindByName(t *testing.T) {
	t.Run("성공", func(t *testing.T) {
		name := "park"
		mockUser := models.NewUser().BuildName("park").BuildId(1)

		mockUserRepo := new(mocks.MockUserRepository)
		us := NewUserService(mockUserRepo)

		mockUserRepo.On("FindUserByName", name).Return(mockUser, nil)

		find_user, err := us.FindUserByName(name)

		assert.NoError(t, err)
		assert.Equal(t, find_user.Name, mockUser.Name)
		assert.Equal(t, find_user.ID, 1)
	})
	t.Run("실패", func(t *testing.T) {
		name := "park"

		mockUserRepo := new(mocks.MockUserRepository)
		us := NewUserService(mockUserRepo)

		mockUserRepo.On("FindUserByName", name).Return(nil, errors.New("not found"))

		find_user, err := us.FindUserByName(name)

		assert.Error(t, err)
		assert.Nil(t, find_user)
	})
}

func Test_UserService_FindAllUsers(t *testing.T) {
	t.Run("성공", func(t *testing.T) {

		mockUsers := []*models.User{
			models.NewUser().BuildId(1).BuildName("park"),
			models.NewUser().BuildId(2).BuildName("kim"),
			models.NewUser().BuildId(3).BuildName("cho"),
		}

		mockUserRepo := new(mocks.MockUserRepository)
		us := NewUserService(mockUserRepo)

		mockUserRepo.On("FindAllUsers").Return(mockUsers, nil)

		find_users, err := us.FindUsers()

		assert.NoError(t, err)
		assert.Equal(t, find_users[0].ID, mockUsers[0].ID)
		assert.Equal(t, find_users[0].Name, mockUsers[0].Name)

	})
	t.Run("실패", func(t *testing.T) {

		mockUserRepo := new(mocks.MockUserRepository)
		us := NewUserService(mockUserRepo)

		mockUserRepo.On("FindAllUsers").Return(nil, errors.New("not found"))

		find_users, err := us.FindUsers()

		assert.Error(t, err)
		assert.Nil(t, find_users)
	})
}

func Test_UserService_CreateUser(t *testing.T) {
	t.Run("성공", func(t *testing.T) {

		mockUser := models.NewUser().BuildName("park").BuildId(1)

		mockUserRepo := new(mocks.MockUserRepository)
		us := NewUserService(mockUserRepo)

		mockUserRepo.On("CreateUser", mockUser).Return(mockUser, nil)

		find_user, err := us.CreateUser(mockUser)

		assert.NoError(t, err)
		assert.Equal(t, find_user.ID, mockUser.ID)
		assert.Equal(t, find_user.Name, mockUser.Name)

	})
	t.Run("실패", func(t *testing.T) {

		mockUser := models.NewUser().BuildName("park").BuildId(1)
		mockUserRepo := new(mocks.MockUserRepository)
		us := NewUserService(mockUserRepo)

		mockUserRepo.On("CreateUser", mockUser).Return(nil, errors.New("db error"))

		find_users, err := us.CreateUser(mockUser)

		assert.Error(t, err)
		assert.Nil(t, find_users)
	})
}
