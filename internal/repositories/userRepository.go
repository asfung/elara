package repositories

import "github.com/asfung/elara/internal/entities"

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
	Update(user entities.User) (entities.User, error)
	FindById(id uint32) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
	FindByRefreshToken(refreshToken string) (entities.User, error)
	Delete(id uint32) error
}
