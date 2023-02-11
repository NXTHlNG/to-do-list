package routes

import (
	"github.com/gofiber/fiber/v2"
	"to-do-list/controllers"
)

func TaskRoute(app *fiber.App) {
	app.Get("/api/task/:id?", controllers.GetTask)
	app.Get("/api/tasks/", controllers.GetAllTasks)
	app.Post("/api/task", controllers.CreateTask)
	app.Put("api/task/:id", controllers.UpdateTask)
	app.Delete("api/task/:id", controllers.DeleteTask)
}
