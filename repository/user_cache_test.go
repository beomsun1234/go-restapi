package repository

import (
	"encoding/json"
	"errors"
	"github/beomsun1234/go-restapi/models"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

var (
	redis_db      *redis.Client
	redis_mock    redismock.ClientMock
	userCacheRepo UserCacheRepositoryInterface
)

func setUpRedisMock() {
	redis_db, redis_mock = redismock.NewClientMock()
}

func setUpUserCacheRepo() {
	userCacheRepo = NewUserCacheRepository(redis_db)
}

func convertStructToBytes(u *models.User) []byte {
	user_json_data, _ := json.Marshal(u)
	return user_json_data
}

func convertStructListToBytes(u []*models.User) []byte {
	user_json_data, _ := json.Marshal(u)
	return user_json_data
}

func Test_UserCacheRepository_GetData(t *testing.T) {
	t.Run("성공", func(t *testing.T) {
		//given
		setUpRedisMock()
		setUpUserCacheRepo()
		user_id := "1"
		mock_user := &models.User{ID: 1}
		redis_mock.ExpectHGet(user_id, "user").SetVal(string(convertStructToBytes(mock_user)))
		//when
		u, err := userCacheRepo.GetData(user_id)
		//then
		assert.Equal(t, u.ID, mock_user.ID)
		assert.NoError(t, err)
		defer redis_db.Close()
	})
	t.Run("redis 실패", func(t *testing.T) {
		//given
		setUpRedisMock()
		setUpUserCacheRepo()
		user_id := "1"
		redis_mock.ExpectHGet(user_id, "user").SetErr(errors.New("redis error"))
		//when
		u, err := userCacheRepo.GetData(user_id)
		//then
		assert.Error(t, err)
		assert.Nil(t, u)
		defer redis_db.Close()
	})
	t.Run("json 실패", func(t *testing.T) {
		//given
		setUpRedisMock()
		setUpUserCacheRepo()
		user_id := "1"
		redis_mock.ExpectHGet(user_id, "user").SetVal("ok")
		//when
		u, err := userCacheRepo.GetData("")
		//then
		assert.Error(t, err)
		assert.Nil(t, u)
		defer redis_db.Close()
	})
}

func Test_UserCacheRepository_SetData(t *testing.T) {
	t.Run("성공", func(t *testing.T) {
		//given
		setUpRedisMock()
		setUpUserCacheRepo()
		user_id := "1"
		mock_user := &models.User{ID: 1}
		redis_mock.ExpectHSet(user_id, convertStructToBytes(mock_user))
		//when
		err := userCacheRepo.SetData(mock_user)
		//then
		assert.NoError(t, err)
		defer redis_db.Close()
	})
	t.Run("redis 실패", func(t *testing.T) {
		//given
		setUpRedisMock()
		setUpUserCacheRepo()
		user_id := "1"

		redis_mock.ExpectHSet(user_id).SetErr(errors.New("redis error"))
		//when
		u, err := userCacheRepo.GetData(user_id)
		//then
		assert.Error(t, err)
		assert.Nil(t, u)
		defer redis_db.Close()
	})
}

func Test_UserCacheRepository_GetDatas(t *testing.T) {
	t.Run("성공", func(t *testing.T) {
		//given
		setUpRedisMock()
		setUpUserCacheRepo()
		key := "users"
		mock_users := []*models.User{
			models.NewUser().BuildId(1).BuildName("park"),
			models.NewUser().BuildId(2).BuildName("kim"),
		}
		redis_mock.ExpectHGet(key, "users").SetVal(string(convertStructListToBytes(mock_users)))
		//when
		find_users, err := userCacheRepo.GetDatas()
		//then
		assert.NoError(t, err)
		assert.Equal(t, 2, len(find_users))
		defer redis_db.Close()
	})
	t.Run("redis 실패", func(t *testing.T) {
		//given
		setUpRedisMock()
		setUpUserCacheRepo()
		key := "users"
		redis_mock.ExpectHGet(key, "ok").SetErr(errors.New("redis error"))
		//when
		find_users, err := userCacheRepo.GetDatas()
		//then
		assert.Error(t, err)
		assert.Nil(t, find_users)
		defer redis_db.Close()
	})
	t.Run("json 실패", func(t *testing.T) {
		//given
		setUpRedisMock()
		setUpUserCacheRepo()
		key := "users"
		redis_mock.ExpectHGet(key, "ok").SetVal("")
		//when
		u, err := userCacheRepo.GetDatas()
		//then
		assert.Error(t, err)
		assert.Nil(t, u)
		defer redis_db.Close()
	})
}

func Test_UserCacheRepository_SetDatas(t *testing.T) {
	t.Run("성공", func(t *testing.T) {
		//given
		setUpRedisMock()
		setUpUserCacheRepo()
		key := "users"
		mock_users := []*models.User{
			models.NewUser().BuildId(1).BuildName("park"),
			models.NewUser().BuildId(2).BuildName("kim"),
		}
		redis_mock.ExpectHSet(key, mock_users)
		//when
		err := userCacheRepo.SetDatas(mock_users)
		//then
		assert.NoError(t, err)
		defer redis_db.Close()
	})
}

func Test_UserCacheRepository_DelDatasByKey(t *testing.T) {
	t.Run("성공", func(t *testing.T) {
		//given
		setUpRedisMock()
		setUpUserCacheRepo()
		key := "users"
		redis_mock.ExpectHDel(key, "users")
		//when
		err := userCacheRepo.DelDatasByKey(key)
		//then
		assert.NoError(t, err)
		defer redis_db.Close()
	})
}
