package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userID"  bson:"_userID,omitempty"`
	Tag       string             `json:"tag" bson:"tag"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
