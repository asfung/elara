package repositories

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

type PortfolioAssetPostgresRepository struct {
	*BaseRepository[entities.PortfolioAsset]
}

func NewPortfolioAssetPostgresRepository(db database.Database) PortfolioAssetRepository {
	return &PortfolioAssetPostgresRepository{
		BaseRepository: NewBaseRepository[entities.PortfolioAsset](db),
	}
}
