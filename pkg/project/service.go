package project

import (
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertProject(project *entities.Project) (*entities.Project, error)
	FetchProjects() (*[]presenter.Project, error)
	UpdateProject(project *entities.Project) (*entities.Project, error)
	RemoveProject(ID string) error
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

func (s *service) InsertProject(project *entities.Project) (*entities.Project, error) {
	return s.repository.CreateProject(project)
}

func (s *service) FetchProjects() (*[]presenter.Project, error) {
	return s.repository.ReadProject()
}

func (s *service) UpdateProject(project *entities.Project) (*entities.Project, error) {
	return s.repository.UpdateProject(project)
}

func (s *service) RemoveProject(ID string) error {
	return s.repository.DeleteProject(ID)
}
