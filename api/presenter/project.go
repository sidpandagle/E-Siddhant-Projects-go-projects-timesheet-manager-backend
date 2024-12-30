package presenter

import (
	"time"
	"timesheet-manager-backend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userId"  bson:"_userId,omitempty"`
	Project   string             `json:"project" bson:"project"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func ProjectSuccessResponse(data *entities.Project) *fiber.Map {
	project := Project{
		ID:        data.ID,
		UserID:    data.UserID,
		Project:   data.Project,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &fiber.Map{
		"status": true,
		"data":   project,
		"error":  nil,
	}
}

func ProjectsSuccessResponse(data *[]Project) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func ProjectErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
