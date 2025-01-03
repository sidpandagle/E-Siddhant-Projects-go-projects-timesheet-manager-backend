package presenter

import (
	"time"
	"timesheet-manager-backend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the presenter object which will be passed in the response by Handler
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Role      string             `json:"role" bson:"role"`         // e.g., "admin", "user"
	IsActive  bool               `json:"isActive" bson:"isActive"` // Whether the user is active or not
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Profile   Profile            `json:"profile,omitempty" bson:"profile,omitempty"` // Profile information (optional)
}

type LoginRequest struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Profile struct {
	FirstName string `json:"firstName" bson:"firstName,omitempty"`
	LastName  string `json:"lastName" bson:"lastName,omitempty"`
}

// UserSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func UserSuccessResponse(data *entities.User) *fiber.Map {
	profile := Profile{
		FirstName: data.Profile.FirstName,
		LastName:  data.Profile.LastName,
	}
	user := User{
		ID:        data.ID,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
		IsActive:  data.IsActive,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Profile:   profile,
	}
	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}

// UsersSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func UsersSuccessResponse(data *[]User) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// UserErrorResponse is the ErrorResponse that will be passed in the response by Handler
func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
