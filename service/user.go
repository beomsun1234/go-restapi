package service

import (
	"github/beomsun1234/go-restapi/models"
	"github/beomsun1234/go-restapi/repository"
)

type UserServiceInterface interface {
	FindUserById(user_id int) (*models.User, error)
	FindUserByName(user_name string) (*models.User, error)
	FindUsers() ([]*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserService(di_userRepo repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepo: di_userRepo,
	}
}

func (u *UserService) FindUserById(user_id int) (*models.User, error) {
	find_user, err := u.userRepo.FindUserById(user_id)
	if err != nil {
		return nil, err
	}
	return find_user, nil
}

func (u *UserService) FindUserByName(user_name string) (*models.User, error) {
	find_user, err := u.userRepo.FindUserByName(user_name)
	if err != nil {
		return nil, err
	}
	return find_user, nil
}

func (u *UserService) FindUsers() ([]*models.User, error) {
	find_users, err := u.userRepo.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return find_users, nil
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	new_user, err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return new_user, nil
}
