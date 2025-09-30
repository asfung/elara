package repositories

import "github.com/asfung/elara/internal/entities"

type AssetRepository interface {
	Repository[entities.Asset]
}
