package repositories

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

type AuthPostgresRepository struct {
	*BaseRepository[entities.User]
}

func NewAuthPostgresRepository(db database.Database) AuthRepository {
	return &AuthPostgresRepository{
		BaseRepository: NewBaseRepository[entities.User](db),
	}
}

func (a *AuthPostgresRepository) findAcessToken(accessToken string) (entities.User, error) {
	panic("unimplemented")
}
