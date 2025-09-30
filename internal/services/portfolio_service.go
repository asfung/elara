package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type PortfolioService interface {
	CreatePortfolio(req models.AddAPortfolioRequest) (entities.Portfolio, error)
	UpdatePortfolio(req models.UpdatePortfolioRequest) (entities.Portfolio, error)
	GetPortfolioById(id string) (entities.Portfolio, error)
	DeletePortfolio(id string) error
}
