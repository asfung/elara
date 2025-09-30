package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type PortfolioAssetService interface {
	CreatePortfolioAsset(req models.AddAPortfolioAssetRequest) (entities.PortfolioAsset, error)
	UpdatePortfolioAsset(req models.UpdatePortfolioAssetRequest) (entities.PortfolioAsset, error)
	GetPortfolioAssetById(id string) (entities.PortfolioAsset, error)
	DeletePortfolioAsset(id string) error
}
