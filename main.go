package main

import (
	"MiddleTestTask/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	if err := app.Listen(":5000"); err != nil {
		panic(err)
	}
}
