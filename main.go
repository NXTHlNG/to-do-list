package main

import (
	"github.com/gofiber/fiber/v2"
	"to-do-list/configs"
	"to-do-list/routes"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.TaskRoute(app)

	app.Listen(":8000")
}
