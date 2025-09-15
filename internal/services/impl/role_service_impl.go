package impl

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
)

type roleServiceImpl struct {
	roleRepo repositories.RoleRepository
}

func NewRoleServiceImpl(roleRepo repositories.RoleRepository) services.RoleService {
	return &roleServiceImpl{
		roleRepo: roleRepo,
	}
}

func (r *roleServiceImpl) CreateRole(req models.AddRoleRequest) (entities.Role, error) {
	role := entities.Role{
		Name:        req.Name,
		Description: &req.Description,
	}

	created, err := r.roleRepo.Create(role)
	if err != nil {
		return entities.Role{}, err
	}
	return created, err
}

func (r *roleServiceImpl) UpdateRole(req models.UpdateRoleRequest) (entities.Role, error) {
	role, err := r.roleRepo.FindById(req.ID)
	if err != nil {
		return entities.Role{}, err
	}

	if req.Name != "" {
		role.Name = req.Name
	}
	if req.Description != "" {
		role.Description = &req.Description
	}

	updated, err := r.roleRepo.Update(*role)
	if err != nil {
		return entities.Role{}, err
	}
	return updated, nil
}

func (r *roleServiceImpl) GetRoleById(id uint) (entities.Role, error) {
	role, err := r.roleRepo.FindById(id)
	if err != nil {
		return entities.Role{}, err
	}
	return *role, nil
}

func (r *roleServiceImpl) DeleteRole(id uint) error {
	return r.roleRepo.Delete(id)
}
