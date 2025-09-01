package repositories

import "github.com/asfung/elara/internal/entities"

type AuthRepository interface {
	Repository[entities.User]
	findAcessToken(accessToken string) (entities.User, error)
}
