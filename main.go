package main

import (
	"github/beomsun1234/go-restapi/database"
	"github/beomsun1234/go-restapi/middleware"
	"github/beomsun1234/go-restapi/models"
	"github/beomsun1234/go-restapi/repository"
	"github/beomsun1234/go-restapi/service"
	"log"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	//path 중첩 문제 해결법 https://github.com/gofiber/fiber/issues/623
	app := fiber.New(
		fiber.Config{StrictRouting: true},
	)
	db := database.NewPostgres()

	if db.Connection() != nil {
		panic("db error")
	}
	db.PostgresDB.AutoMigrate(&models.User{})
	cacheDb := database.NewRedisDB()
	cacheDb.RedisConnect()

	userRepo := repository.NewUserRepository(db.PostgresDB)
	userCacheRepo := repository.NewUserCacheRepository(cacheDb.Rdb)
	userCacheMiddleware := middleware.NewUserCacheMiddleware(userCacheRepo)
	userService := service.NewUserService(userRepo, userCacheRepo)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("hello")
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		p := &models.User{}
		err := c.BodyParser(p)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"data": err})
		}
		user, err := userService.CreateUser(p)
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		return c.Status(200).JSON(fiber.Map{
			"data": user,
		})
	})
	app.Get("/users", userCacheMiddleware.GetCacheUsers(), func(c *fiber.Ctx) error {
		users, err := userService.FindUsers()
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		return c.Status(200).JSON(fiber.Map{
			"data": users,
		})
	})
	//순서 중요
	app.Get("/users/search", func(c *fiber.Ctx) error {
		parma := c.Query("name")
		user, err := userService.FindUserByName(parma)
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		return c.Status(200).JSON(fiber.Map{
			"data": user,
		})
	})
	app.Get("/users/:userId", userCacheMiddleware.GetCacheUserById(), func(c *fiber.Ctx) error {
		parma := c.Params("userId")
		id, err := strconv.Atoi(parma)

		if err != nil {
			return c.Status(404).SendString(err.Error())
		}

		user, err := userService.FindUserById(id)
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		return c.Status(200).JSON(fiber.Map{
			"data": user,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
