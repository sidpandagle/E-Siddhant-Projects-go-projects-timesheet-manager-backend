package user

import (
	"timesheet-manager-backend/api/presenter"
	"timesheet-manager-backend/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertUser(user *entities.User) (*entities.User, error)
	FetchUsers() (*[]presenter.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	RemoveUser(ID string) error
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

// InsertUser is a service layer that helps insert user in Users
func (s *service) InsertUser(user *entities.User) (*entities.User, error) {
	return s.repository.CreateUser(user)
}

// FetchUsers is a service layer that helps fetch all users in Users
func (s *service) FetchUsers() (*[]presenter.User, error) {
	return s.repository.ReadUser()
}

// UpdateUser is a service layer that helps update users in Users
func (s *service) UpdateUser(user *entities.User) (*entities.User, error) {
	return s.repository.UpdateUser(user)
}

// RemoveUser is a service layer that helps remove users from Users
func (s *service) RemoveUser(ID string) error {
	return s.repository.DeleteUser(ID)
}
