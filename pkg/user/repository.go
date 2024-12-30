package user

import (
	"context"
	"time"
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	ReadUser() (*[]presenter.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(ID string) error
}
type repository struct {
	Collection *mongo.Collection
}

// NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

// CreateUser is a mongo repository that helps to create users
func (r *repository) CreateUser(user *entities.User) (*entities.User, error) {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// ReadUser is a mongo repository that helps to fetch users
func (r *repository) ReadUser() (*[]presenter.User, error) {
	var users []presenter.User
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user presenter.User
		_ = cursor.Decode(&user)
		users = append(users, user)
	}
	return &users, nil
}

// UpdateUser is a mongo repository that helps to update users
func (r *repository) UpdateUser(user *entities.User) (*entities.User, error) {
	user.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser is a mongo repository that helps to delete users
func (r *repository) DeleteUser(ID string) error {
	userId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": userId})
	if err != nil {
		return err
	}
	return nil
}
