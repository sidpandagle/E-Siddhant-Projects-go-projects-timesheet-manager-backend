package handlers

import (
	"net/http"
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"
	"timesheet-manager-backend/pkg/project"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func AddProject(service project.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Project
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProjectErrorResponse(err))
		}
		if requestBody.Project == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProjectErrorResponse(errors.New(
				"Please specify project details")))
		}
		result, err := service.InsertProject(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProjectErrorResponse(err))
		}
		return c.JSON(presenter.ProjectSuccessResponse(result))
	}
}

func UpdateProject(service project.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Project
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProjectErrorResponse(err))
		}
		result, err := service.UpdateProject(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProjectErrorResponse(err))
		}
		return c.JSON(presenter.ProjectSuccessResponse(result))
	}
}

func RemoveProject(service project.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProjectErrorResponse(err))
		}
		projectID := requestBody.ID
		err = service.RemoveProject(projectID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProjectErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "Updated Successfully!",
			"err":    nil,
		})
	}
}

func GetProjects(service project.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchProjects()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProjectErrorResponse(err))
		}
		return c.JSON(presenter.ProjectsSuccessResponse(fetched))
	}
}

func GetProjectByUserID(service project.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		fetched, err := service.FetchProjectByUserID(userId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProjectErrorResponse(err))
		}
		return c.JSON(presenter.ProjectsSuccessResponse(fetched))
	}
}
