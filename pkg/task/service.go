package task

import (
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertTask(task *entities.Task) (*entities.Task, error)
	FetchTasks() (*[]presenter.Task, error)
	FetchTasksByUserId(userId string, page int, pageSize int) (*[]presenter.Task, error)
	UpdateTask(task *entities.Task) (*entities.Task, error)
	RemoveTask(ID string) error
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

func (s *service) InsertTask(task *entities.Task) (*entities.Task, error) {
	return s.repository.CreateTask(task)
}

func (s *service) FetchTasks() (*[]presenter.Task, error) {
	return s.repository.ReadTask()
}

func (s *service) FetchTasksByUserId(userId string, page int, pageSize int) (*[]presenter.Task, error) {
	return s.repository.ReadTaskByUserId(userId, page, pageSize)
}

func (s *service) UpdateTask(task *entities.Task) (*entities.Task, error) {
	return s.repository.UpdateTask(task)
}

func (s *service) RemoveTask(ID string) error {
	return s.repository.DeleteTask(ID)
}
