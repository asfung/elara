package repositories

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

type RolePostgresRepository struct {
	*BaseRepository[entities.Role]
}

func NewRolePostgresRepository(db database.Database) RoleRepository {
	return &RolePostgresRepository{
		BaseRepository: NewBaseRepository[entities.Role](db),
	}
}
