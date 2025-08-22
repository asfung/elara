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
		Email:    req.Email,
		Password: &hashedPassword,
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

	updatedUser, err := u.repo.Update(user)
	if err != nil {
		return entities.User{}, err
	}

	return updatedUser, nil
}

func (u *userServiceImpl) GetUserById(id uint32) (entities.User, error) {
	user, err := u.repo.FindById(id)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (u *userServiceImpl) DeleteUser(id uint32) error {
	return u.repo.Delete(id)
}

func (u *userServiceImpl) RefreshToken(req models.RefreshTokenRequest) (models.AuthResponse, error) {
	panic("unimplemented")
	// user, err := s.repo.FindByRefreshToken(req.RefreshToken)
	// if err != nil {
	// 	return models.AuthResponse{}, err
	// }
	// if user.RefreshToken != req.RefreshToken {
	// 	return models.AuthResponse{}, errors.New("invalid refresh token")
	// }

	// newAccess := "new-access-token"
	// newRefresh := "new-refresh-token"
	// expiry := time.Now().Add(24 * time.Hour)

	// user.AccessToken = newAccess
	// user.RefreshToken = newRefresh
	// user.ExpiresAt = expiry

	// _, err = s.repo.Update(user)
	// if err != nil {
	// 	return models.AuthResponse{}, err
	// }

	// return models.AuthResponse{
	// 	AccessToken:  newAccess,
	// 	RefreshToken: newRefresh,
	// 	ExpiresAt:    expiry,
	// }, nil
}
