package repository

import (
	"errors"
	"github/beomsun1234/go-restapi/models"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindUserById(user_id int) (*models.User, error)
	FindUserByName(user_name string) (*models.User, error)
	FindAllUsers() ([]*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(di_db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		db: di_db,
	}
}

func (userRepo *UserRepository) FindUserById(user_id int) (*models.User, error) {
	find_user := &models.User{}
	result := userRepo.db.Find(find_user, "id=?", user_id)

	if result.RowsAffected < 1 {
		return nil, errors.New("not found user")
	}

	return find_user, nil
}

func (userRepo *UserRepository) FindUserByName(user_name string) (*models.User, error) {
	find_user := &models.User{}
	result := userRepo.db.Find(find_user, "name=?", user_name)
	if result.RowsAffected < 1 {
		return nil, errors.New("not found user")
	}

	return find_user, nil
}

func (userRepo *UserRepository) FindAllUsers() ([]*models.User, error) {
	find_users := &[]*models.User{}
	result := userRepo.db.Find(find_users)
	if result.RowsAffected < 1 {
		return nil, errors.New("not found user")
	}

	return *find_users, nil
}
func (userRepo *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	result := userRepo.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
