package repository

import (
	"context"
	"encoding/json"
	"github/beomsun1234/go-restapi/models"
	"strconv"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type UserCacheRepositoryInterface interface {
	GetData(user_id string) (*models.User, error)
	SetData(u *models.User) error
	SetDatas(u []*models.User) error
	GetDatas() ([]*models.User, error)
	DelDatasByKey(key string) error
}

type UserCacheRepository struct {
	Rdb *redis.Client
}

func NewUserCacheRepository(di_rdb *redis.Client) UserCacheRepositoryInterface {
	return &UserCacheRepository{
		Rdb: di_rdb,
	}
}

func (r *UserCacheRepository) GetData(user_id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	ret_user := &models.User{}
	data, err := r.Rdb.HGet(ctx, user_id, "user").Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), ret_user)
	if err != nil {
		return nil, err
	}
	return ret_user, nil
}

func (r *UserCacheRepository) SetData(u *models.User) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	r.Rdb.HSet(ctx, strconv.FormatUint(uint64(u.ID), 10), "user", data)

	defer cancel()
	return nil
}

func (r *UserCacheRepository) GetDatas() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	ret_user := &[]*models.User{}
	data, err := r.Rdb.HGet(ctx, "users", "users").Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), ret_user)
	if err != nil {
		return nil, err
	}
	return *ret_user, nil
}

func (r *UserCacheRepository) SetDatas(u []*models.User) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	r.Rdb.HSet(ctx, "users", "users", data)

	defer cancel()
	return nil
}

func (r *UserCacheRepository) DelDatasByKey(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	r.Rdb.Del(ctx, "users")

	defer cancel()
	return nil
}
