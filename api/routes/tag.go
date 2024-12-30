package routes

import (
	"timesheet-manager-backend/api/handlers"
	"timesheet-manager-backend/pkg/tag"

	"github.com/gofiber/fiber/v2"
)

func TagRouter(app fiber.Router, service tag.Service) {
	app.Get("/tags", handlers.GetTags(service))
	app.Get("/tags/:userId", handlers.GetTagByUserID(service))
	app.Post("/tags", handlers.AddTag(service))
	app.Put("/tags", handlers.UpdateTag(service))
	app.Delete("/tags", handlers.RemoveTag(service))
}
