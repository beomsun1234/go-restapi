package mocks

import (
	"errors"
	"github/beomsun1234/go-restapi/models"
	"github/beomsun1234/go-restapi/repository"
	"log"
	"strconv"
)

type MockUserCacheRepository struct {
}

func NewMockUserCacheRepository() repository.UserCacheRepositoryInterface {
	return &MockUserCacheRepository{}
}

func (r *MockUserCacheRepository) GetData(user_id string) (*models.User, error) {
	if user_id == "" {
		return nil, errors.New("error")
	}
	u_id, _ := strconv.Atoi(user_id)

	return models.NewUser().BuildId(u_id), nil
}

func (r *MockUserCacheRepository) SetData(u *models.User) error {
	if u == nil {
		return errors.New("error")
	}
	return nil
}

func (r *MockUserCacheRepository) GetDatas() ([]*models.User, error) {
	log.Println("get cache data")
	return nil, nil
}

func (r *MockUserCacheRepository) SetDatas(u []*models.User) error {
	log.Println("set cache datas")
	return nil
}

func (r *MockUserCacheRepository) DelDatasByKey(key string) error {
	log.Println("delete cache datas by key")
	return nil
}
