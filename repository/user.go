package repository

import (
	"errors"
	"github/beomsun1234/go-restapi/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserById(user_id int) (*models.User, error)
	FindUserByName(user_name string) (*models.User, error)
	FindAllUsers() (*[]models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(di_db *gorm.DB) UserRepository {
	return &userRepository{
		db: di_db,
	}
}

func (userRepo *userRepository) FindUserById(user_id int) (*models.User, error) {
	find_user := &models.User{}
	result := userRepo.db.Find(find_user, "id=?", user_id)

	if result.RowsAffected < 1 {
		return nil, errors.New("not found user")
	}

	return find_user, nil
}

func (userRepo *userRepository) FindUserByName(user_name string) (*models.User, error) {
	find_user := &models.User{}
	result := userRepo.db.Find(find_user, "name=?", user_name)
	if result.RowsAffected < 1 {
		return nil, errors.New("not found user")
	}

	return find_user, nil
}

func (userRepo *userRepository) FindAllUsers() (*[]models.User, error) {
	var find_users []models.User
	result := userRepo.db.Find(&find_users)

	if result.RowsAffected < 1 {
		return nil, errors.New("not found user")
	}

	return &find_users, nil
}
func (userRepo *userRepository) CreateUser(user *models.User) (*models.User, error) {
	result := userRepo.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
