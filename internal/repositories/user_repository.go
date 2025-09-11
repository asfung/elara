package repositories

import "github.com/asfung/elara/internal/entities"

type UserRepository interface {
	Repository[entities.User]
	FindByEmail(email string) (entities.User, error)
	FindByRefreshToken(refreshToken string) (entities.User, error)
	FindByUserId(userId string) (entities.User, error)
	FindByProvider(provider string, providerUserId string) (entities.User, error)
}
