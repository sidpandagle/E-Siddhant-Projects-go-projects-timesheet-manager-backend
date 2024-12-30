package presenter

import (
	"time"
	"timesheet-manager-backend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Task      string             `json:"task" bson:"task"`
	Project   string             `json:"project" bson:"project"`
	Tags      []string           `json:"tags" bson:"tags"`
	StartTime time.Time          `json:"startTime" bson:"startTime"`
	EndTime   time.Time          `json:"endTime" bson:"endTime"`
	UserID    primitive.ObjectID `json:"userID,omitempty" bson:"_userId,omitempty"`
}

func TaskSuccessResponse(data *entities.Task) *fiber.Map {
	task := Task{
		ID:        data.ID,
		Task:      data.Task,
		Project:   data.Project,
		Tags:      data.Tags,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
		UserID:    data.UserID,
	}
	return &fiber.Map{
		"status": true,
		"data":   task,
		"error":  nil,
	}
}

func TasksSuccessResponse(data *[]Task) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func TaskErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
