package routes

import (
	"timesheet-manager-backend/api/handlers"
	"timesheet-manager-backend/pkg/task"

	"github.com/gofiber/fiber/v2"
)

func TaskRouter(app fiber.Router, service task.Service) {
	app.Get("/tasks", handlers.GetTasks(service))
	app.Get("/tasks/:userId", handlers.GetTasksByUserId(service))
	app.Post("/tasks", handlers.AddTask(service))
	app.Put("/tasks", handlers.UpdateTask(service))
	app.Delete("/tasks", handlers.RemoveTask(service))
}
