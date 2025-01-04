package routes

import (
	"timesheet-manager-backend/api/handlers"
	"timesheet-manager-backend/pkg/project"

	"github.com/gofiber/fiber/v2"
)

func ProjectRouter(app fiber.Router, service project.Service) {
	app.Get("/projects", handlers.GetProjects(service))
	app.Get("/projects/:userId", handlers.GetProjectsByUserId(service))
	app.Post("/projects", handlers.AddProject(service))
	app.Put("/projects", handlers.UpdateProject(service))
	app.Delete("/projects", handlers.RemoveProject(service))
}
