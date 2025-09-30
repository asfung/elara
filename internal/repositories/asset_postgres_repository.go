package repositories

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

type AssetPostgresRepository struct {
	*BaseRepository[entities.Asset]
}

func NewAssetPostgresRepository(db database.Database) AssetRepository {
	return &AssetPostgresRepository{
		BaseRepository: NewBaseRepository[entities.Asset](db),
	}
}
