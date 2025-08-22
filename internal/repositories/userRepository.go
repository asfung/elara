package repositories

import "github.com/asfung/elara/internal/entities"

type UserRepository interface {
	Repository[entities.User]
	FindByEmail(email string) (entities.User, error)
	FindByRefreshToken(refreshToken string) (entities.User, error)
}
