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
	Time      time.Duration      `json:"time" bson:"time"`
	UserID    primitive.ObjectID `json:"userId,omitempty" bson:"_userId,omitempty"`
}

func TaskSuccessResponse(data *entities.Task) *fiber.Map {
	task := Task{
		ID:        data.ID,
		Task:      data.Task,
		Project:   data.Project,
		Tags:      data.Tags,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
		Time:      data.EndTime.Sub(data.StartTime),
		UserID:    data.UserID,
	}
	return &fiber.Map{
		"status": true,
		"data":   task,
		"error":  nil,
	}
}

func TasksSuccessResponse(data *[]Task) *fiber.Map {
	var tasks []Task

	// Iterate over the slice of tasks and calculate the Time field
	for _, t := range *data {
		task := Task{
			ID:        t.ID,
			Task:      t.Task,
			Project:   t.Project,
			Tags:      t.Tags,
			StartTime: t.StartTime,
			EndTime:   t.EndTime,
			Time:      t.EndTime.Sub(t.StartTime),
			UserID:    t.UserID,
		}
		tasks = append(tasks, task)
	}

	return &fiber.Map{
		"status": true,
		"data":   tasks,
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
