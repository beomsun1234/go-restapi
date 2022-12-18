package main

import (
	"github/beomsun1234/go-restapi/database"
	"github/beomsun1234/go-restapi/models"
	"github/beomsun1234/go-restapi/repository"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := database.NewPostgres()

	if db.Connection() != nil {
		panic("db error")
	}
	db.PostgresDB.AutoMigrate(&models.User{})

	userRepo := repository.NewUserRepository(db.PostgresDB)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("hello")
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := userRepo.FindAllUsers()
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		return c.Status(200).JSON(fiber.Map{
			"data": users,
		})
	})
	log.Fatal(app.Listen(":3000"))
}
