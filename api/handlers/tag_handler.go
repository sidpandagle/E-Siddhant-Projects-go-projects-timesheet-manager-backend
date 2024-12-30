package handlers

import (
	"net/http"
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"
	"timesheet-manager-backend/pkg/tag"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func AddTag(service tag.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Tag
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TagErrorResponse(err))
		}
		if requestBody.Tag == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TagErrorResponse(errors.New(
				"Please specify tag details")))
		}
		result, err := service.InsertTag(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TagErrorResponse(err))
		}
		return c.JSON(presenter.TagSuccessResponse(result))
	}
}

func UpdateTag(service tag.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Tag
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TagErrorResponse(err))
		}
		result, err := service.UpdateTag(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TagErrorResponse(err))
		}
		return c.JSON(presenter.TagSuccessResponse(result))
	}
}

func RemoveTag(service tag.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TagErrorResponse(err))
		}
		tagID := requestBody.ID
		err = service.RemoveTag(tagID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TagErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func GetTags(service tag.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchTags()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TagErrorResponse(err))
		}
		return c.JSON(presenter.TagsSuccessResponse(fetched))
	}
}

func GetTagByUserID(service tag.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		fetched, err := service.FetchTagByUserID(userId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TagErrorResponse(err))
		}
		return c.JSON(presenter.TagsSuccessResponse(fetched))
	}
}
