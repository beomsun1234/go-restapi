package repository

import (
	"testing"

	"github/beomsun1234/go-restapi/database"
	"github/beomsun1234/go-restapi/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	sqliteDB *gorm.DB
	userRepo UserRepositoryInterface
)

func setUpRepository() {
	sqlite := database.NewSqlite()
	sqlite.Connection()

	sqliteDB = sqlite.SqliteDB

	sqlite.SqliteDB.AutoMigrate(&models.User{})
	userRepo = NewUserRepository(sqlite.SqliteDB)
}

func closeConnection() {
	sqliteDB.Migrator().DropTable(&models.User{})
	sqlDB, _ := sqliteDB.DB()
	sqlDB.Close()
}

func Test_UserRepository_CeateUser(t *testing.T) {
	t.Run("유저 저장 성공", func(t *testing.T) {

		setUpRepository()
		//given
		user := models.NewUser().BuildName("park")
		//when
		created_user, err := userRepo.CreateUser(user)
		//then
		assert.NoError(t, err)
		assert.Equal(t, user.Name, created_user.Name)
		closeConnection()
	})
	t.Run("유저 저장 실패", func(t *testing.T) {
		setUpRepository()
		_, err := userRepo.CreateUser(nil)
		assert.Error(t, err)
		closeConnection()
	})
}

func Test_UserRepository_FindUSerById(t *testing.T) {
	t.Run("유저 조회 성공", func(t *testing.T) {
		setUpRepository()
		//given
		created_user, _ := userRepo.CreateUser(models.NewUser().BuildName("testuser"))
		//when
		find_user, err := userRepo.FindUserById(created_user.ID)
		//then
		assert.NoError(t, err)
		assert.Equal(t, find_user.ID, created_user.ID)

		closeConnection()
	})
	t.Run("조회 실패", func(t *testing.T) {
		setUpRepository()
		_, err := userRepo.FindUserById(-1)
		assert.Error(t, err)
		closeConnection()
	})
}

func Test_UserRepository_FindUSerByName(t *testing.T) {
	t.Run("유저 이름 조회 성공", func(t *testing.T) {
		setUpRepository()
		//given
		created_user, _ := userRepo.CreateUser(models.NewUser().BuildName("testuser"))
		//when
		find_user, err := userRepo.FindUserByName("testuser")

		//then
		assert.NoError(t, err)
		assert.Equal(t, find_user.ID, created_user.ID)

		closeConnection()
	})
	t.Run("조회 실패", func(t *testing.T) {
		setUpRepository()
		_, err := userRepo.FindUserByName("")
		assert.Error(t, err)
		closeConnection()
	})
}

func Test_UserRepository_FindAllUsers(t *testing.T) {
	t.Run("유저 전체 조회 성공", func(t *testing.T) {
		setUpRepository()
		//given
		_, _ = userRepo.CreateUser(models.NewUser().BuildName("testuser"))
		//when
		find_users, err := userRepo.FindAllUsers()

		//then
		assert.NoError(t, err)
		assert.Equal(t, len(find_users), 1)

		closeConnection()
	})
	t.Run("유저 전체 조회 실패", func(t *testing.T) {
		setUpRepository()
		_, err := userRepo.FindAllUsers()
		assert.Error(t, err)
		closeConnection()
	})
}
