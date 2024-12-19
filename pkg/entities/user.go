package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User Constructs your User model under entities.
type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email"`
	PasswordHash string             `json:"-" bson:"passwordHash"`    // Store hashed password, not plaintext
	Role         string             `json:"role" bson:"role"`         // e.g., "admin", "user"
	IsActive     bool               `json:"isActive" bson:"isActive"` // Whether the user is active or not
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
	Profile      Profile            `json:"profile,omitempty" bson:"profile,omitempty"` // Profile information (optional)
}

// Profile contains optional user details.
type Profile struct {
	FirstName string `json:"firstName" bson:"firstName,omitempty"`
	LastName  string `json:"lastName" bson:"lastName,omitempty"`
}
