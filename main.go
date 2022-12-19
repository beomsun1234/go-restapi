package main

import (
	"github/beomsun1234/go-restapi/database"
	"github/beomsun1234/go-restapi/models"
	"github/beomsun1234/go-restapi/repository"
	"github/beomsun1234/go-restapi/service"
	"log"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(
		fiber.Config{StrictRouting: true},
	)

	db := database.NewPostgres()

	if db.Connection() != nil {
		panic("db error")
	}
	db.PostgresDB.AutoMigrate(&models.User{})

	userRepo := repository.NewUserRepository(db.PostgresDB)
	userService := service.NewUserService(userRepo)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("hello")
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := userService.FindUsers()
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		return c.Status(200).JSON(fiber.Map{
			"data": users,
		})
	})
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
	app.Get("/users/:userId", func(c *fiber.Ctx) error {
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
