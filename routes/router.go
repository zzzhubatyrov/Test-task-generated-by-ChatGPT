package routes

import (
	"MiddleTestTask/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/tasks", handlers.GetTasks)
	app.Post("/tasks", handlers.CreateTasks)
	app.Get("/tasks/:id", handlers.GetTasksId)
	app.Put("/tasks/:id", handlers.EditTasks)
	app.Delete("/tasks/:id", handlers.DeleteTasks)
}
