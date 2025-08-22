package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type UserService interface {
	CreateUser(req models.AddUserRequest) (entities.User, error)
	UpdateUser(req models.UpdateUserRequest) (entities.User, error)
	GetUserById(id uint32) (entities.User, error)
	DeleteUser(id uint32) error
}
