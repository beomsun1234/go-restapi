package middleware

import (
	"fmt"
	"github/beomsun1234/go-restapi/repository"

	fiber "github.com/gofiber/fiber/v2"
)

type UserCacheMiddleware struct {
	userCacheRepo repository.UserCacheRepositoryInterface
}

func NewUserCacheMiddleware(di_userCacheRepo repository.UserCacheRepositoryInterface) *UserCacheMiddleware {
	return &UserCacheMiddleware{
		userCacheRepo: di_userCacheRepo,
	}
}

func (u *UserCacheMiddleware) GetCacheUserById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		parma := c.Params("userId")

		data, err := u.userCacheRepo.GetData(parma)

		if err != nil {
			fmt.Println(err)
			return c.Next()
		}
		return c.Status(200).JSON(fiber.Map{"data": data})
	}
}

func (u *UserCacheMiddleware) GetCacheUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {

		data, err := u.userCacheRepo.GetDatas()

		if err != nil {
			fmt.Println(err)
			return c.Next()
		}
		return c.Status(200).JSON(fiber.Map{"data": data})
	}
}
