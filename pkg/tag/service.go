package tag

import (
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertTag(tag *entities.Tag) (*entities.Tag, error)
	FetchTags() (*[]presenter.Tag, error)
	UpdateTag(tag *entities.Tag) (*entities.Tag, error)
	RemoveTag(ID string) error
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertTag(tag *entities.Tag) (*entities.Tag, error) {
	return s.repository.CreateTag(tag)
}

func (s *service) FetchTags() (*[]presenter.Tag, error) {
	return s.repository.ReadTag()
}

func (s *service) UpdateTag(tag *entities.Tag) (*entities.Tag, error) {
	return s.repository.UpdateTag(tag)
}

func (s *service) RemoveTag(ID string) error {
	return s.repository.DeleteTag(ID)
}
