package presenter

import (
	"time"
	"timesheet-manager-backend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userID"  bson:"_userID,omitempty"`
	Tag       string             `json:"tag" bson:"tag"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func TagSuccessResponse(data *entities.Tag) *fiber.Map {
	tag := Tag{
		ID:        data.ID,
		UserID:    data.UserID,
		Tag:       data.Tag,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &fiber.Map{
		"status": true,
		"data":   tag,
		"error":  nil,
	}
}

func TagsSuccessResponse(data *[]Tag) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func TagErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
