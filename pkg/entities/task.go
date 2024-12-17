package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Task      string             `json:"task" bson:"task"`
	Project   string             `json:"project" bson:"project"`
	Tags      []string           `json:"tags" bson:"tags"`
	StartTime time.Time          `json:"startTime" bson:"startTime"`
	EndTime   time.Time          `json:"endTime" bson:"endTime"`
	UserID    primitive.ObjectID `json:"userId,omitempty" bson:"_userId,omitempty"`
}
