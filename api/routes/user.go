package routes

import (
	"timesheet-manager-backend/api/handlers"
	"timesheet-manager-backend/pkg/user"

	"github.com/gofiber/fiber/v2"
)

// UserRouter is the Router for GoFiber App
func UserRouter(app fiber.Router, service user.Service) {
	app.Get("/users", handlers.GetUsers(service))
	app.Post("/users", handlers.AddUser(service))
	app.Put("/users", handlers.UpdateUser(service))
	app.Delete("/users", handlers.RemoveUser(service))
	app.Post("/login", handlers.LoginUser(service))
}
