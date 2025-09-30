package repositories

import "github.com/asfung/elara/internal/entities"

type PortfolioRepository interface {
	Repository[entities.Portfolio]
}
