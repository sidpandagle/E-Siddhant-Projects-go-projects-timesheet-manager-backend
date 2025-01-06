package task

import (
	"context"
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateTask(task *entities.Task) (*entities.Task, error)
	ReadTask() (*[]presenter.Task, error)
	ReadTaskByUserId(userId string, page int, pageSize int) (*[]presenter.Task, error)
	UpdateTask(task *entities.Task) (*entities.Task, error)
	DeleteTask(ID string) error
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

// CreateTask is a mongo repository that helps to create tasks
func (r *repository) CreateTask(task *entities.Task) (*entities.Task, error) {
	task.ID = primitive.NewObjectID()
	// task.CreatedAt = time.Now()
	// task.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// ReadTask is a mongo repository that helps to fetch tasks
func (r *repository) ReadTask() (*[]presenter.Task, error) {
	var tasks []presenter.Task
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var task presenter.Task
		_ = cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return &tasks, nil
}

// ReadTaskByUserId is a mongo repository that helps to fetch tasks by userId
func (r *repository) ReadTaskByUserId(userId string, page int, pageSize int) (*[]presenter.Task, error) {
	var tasks []presenter.Task

	// Calculate the number of documents to skip
	skip := (page - 1) * pageSize

	// Convert userId to ObjectID
	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	// Set up filter to fetch documents for the specific userId
	filter := bson.M{"_userId": objectID}

	// Set pagination options
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(pageSize))

	// Execute the query with pagination
	cursor, err := r.Collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor to decode documents
	for cursor.Next(context.Background()) {
		var task presenter.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &tasks, nil
}

// UpdateTask is a mongo repository that helps to update tasks
func (r *repository) UpdateTask(task *entities.Task) (*entities.Task, error) {
	// task.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": task.ID}, bson.M{"$set": task})
	if err != nil {
		return nil, err
	}
	return task, nil
}

// DeleteTask is a mongo repository that helps to delete tasks
func (r *repository) DeleteTask(ID string) error {
	taskID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": taskID})
	if err != nil {
		return err
	}
	return nil
}
