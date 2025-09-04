package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type UserService interface {
	CreateUser(req models.AddUserRequest) (entities.User, error)
	UpdateUser(req models.UpdateUserRequest) (entities.User, error)
	GetUserById(id string) (entities.User, error)
	DeleteUser(id string) error
	GetUserByUserId(userId string) (entities.User, error)
}
