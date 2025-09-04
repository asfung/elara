package impl

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
	"github.com/asfung/elara/utils"
)

type userServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserServiceImpl(userRepository repositories.UserRepository) services.UserService {
	return &userServiceImpl{
		repo: userRepository,
	}
}

func (u *userServiceImpl) CreateUser(req models.AddUserRequest) (entities.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return entities.User{}, err
	}

	user := entities.User{
		Username: req.Username,
		Name:     req.Username,
		Email:    req.Email,
		Password: &hashedPassword,
		Provider: "local",
	}

	createdUser, err := u.repo.Create(user)
	if err != nil {
		return entities.User{}, err
	}
	return createdUser, nil
}

func (u *userServiceImpl) UpdateUser(req models.UpdateUserRequest) (entities.User, error) {
	user, err := u.repo.FindById(req.Id)
	if err != nil {
		return entities.User{}, err
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.FirstName != nil {
		user.FirstName = req.FirstName
	}
	if req.LastName != nil {
		user.LastName = req.LastName
	}

	updatedUser, err := u.repo.Update(*user)
	if err != nil {
		return entities.User{}, err
	}

	return updatedUser, nil
}

func (u *userServiceImpl) GetUserById(id string) (entities.User, error) {
	user, err := u.repo.FindById(id)
	if err != nil {
		return entities.User{}, err
	}
	return *user, nil
}

func (u *userServiceImpl) DeleteUser(id string) error {
	return u.repo.Delete(id)
}

func (u *userServiceImpl) GetUserByUserId(userId string) (entities.User, error) {
	user, err := u.repo.FindByUserId(userId)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
