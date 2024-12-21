package project

import (
	"context"
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateProject(project *entities.Project) (*entities.Project, error)
	ReadProject() (*[]presenter.Project, error)
	UpdateProject(project *entities.Project) (*entities.Project, error)
	DeleteProject(ID string) error
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

// CreateProject is a mongo repository that helps to create projects
func (r *repository) CreateProject(project *entities.Project) (*entities.Project, error) {
	project.ID = primitive.NewObjectID()
	// project.CreatedAt = time.Now()
	// project.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), project)
	if err != nil {
		return nil, err
	}
	return project, nil
}

// ReadProject is a mongo repository that helps to fetch projects
func (r *repository) ReadProject() (*[]presenter.Project, error) {
	var projects []presenter.Project
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var project presenter.Project
		_ = cursor.Decode(&project)
		projects = append(projects, project)
	}
	return &projects, nil
}

// UpdateProject is a mongo repository that helps to update projects
func (r *repository) UpdateProject(project *entities.Project) (*entities.Project, error) {
	// project.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": project.ID}, bson.M{"$set": project})
	if err != nil {
		return nil, err
	}
	return project, nil
}

// DeleteProject is a mongo repository that helps to delete projects
func (r *repository) DeleteProject(ID string) error {
	projectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": projectID})
	if err != nil {
		return err
	}
	return nil
}
