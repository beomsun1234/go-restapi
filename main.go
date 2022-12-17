package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/users/:userId", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("hello")
	})
	log.Fatal(app.Listen(":3000"))
}
