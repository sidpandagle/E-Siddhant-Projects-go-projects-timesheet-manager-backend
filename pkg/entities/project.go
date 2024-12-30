package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userId"  bson:"_userId,omitempty"`
	Project   string             `json:"project" bson:"project"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
