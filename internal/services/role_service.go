package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type RoleService interface {
	CreateRole(req models.AddRoleRequest) (entities.Role, error)
	UpdateRole(req models.UpdateRoleRequest) (entities.Role, error)
	GetRoleById(id uint) (entities.Role, error)
	DeleteRole(id uint) error
}
