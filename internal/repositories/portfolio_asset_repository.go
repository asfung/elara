package repositories

import "github.com/asfung/elara/internal/entities"

type PortfolioAssetRepository interface {
	Repository[entities.PortfolioAsset]
}
