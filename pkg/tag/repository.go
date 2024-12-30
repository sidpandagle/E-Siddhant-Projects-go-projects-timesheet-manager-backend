package tag

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
	CreateTag(tag *entities.Tag) (*entities.Tag, error)
	ReadTag() (*[]presenter.Tag, error)
	ReadTagByUserID(userId string) (*[]presenter.Tag, error)
	UpdateTag(tag *entities.Tag) (*entities.Tag, error)
	DeleteTag(ID string) error
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

// CreateTag is a mongo repository that helps to create tags
func (r *repository) CreateTag(tag *entities.Tag) (*entities.Tag, error) {
	tag.ID = primitive.NewObjectID()
	// tag.CreatedAt = time.Now()
	// tag.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

// Readtag is a mongo repository that helps to fetch tags
func (r *repository) ReadTag() (*[]presenter.Tag, error) {
	var tags []presenter.Tag
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var tag presenter.Tag
		_ = cursor.Decode(&tag)
		tags = append(tags, tag)
	}
	return &tags, nil
}

func (r *repository) ReadTagByUserID(userId string) (*[]presenter.Tag, error) {
	var tags []presenter.Tag
	objectID, err := primitive.ObjectIDFromHex(userId)
	cursor, err := r.Collection.Find(context.Background(), bson.M{"_userId": objectID})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var tag presenter.Tag
		_ = cursor.Decode(&tag)
		tags = append(tags, tag)
	}
	return &tags, nil
}

// Updatetag is a mongo repository that helps to update tags
func (r *repository) UpdateTag(tag *entities.Tag) (*entities.Tag, error) {
	// tag.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": tag.ID}, bson.M{"$set": tag})
	if err != nil {
		return nil, err
	}
	return tag, nil
}

// Deletetag is a mongo repository that helps to delete tags
func (r *repository) DeleteTag(ID string) error {
	tagID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": tagID})
	if err != nil {
		return err
	}
	return nil
}
