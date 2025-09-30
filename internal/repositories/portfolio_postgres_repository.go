package repositories

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

type PortfolioPostgresRepository struct {
	*BaseRepository[entities.Portfolio]
}

func NewPortfolioPostgresRepository(db database.Database) PortfolioRepository {
	return &PortfolioPostgresRepository{
		BaseRepository: NewBaseRepository[entities.Portfolio](db),
	}
}
