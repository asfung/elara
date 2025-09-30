package impl

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
)

type portfolioAssetServiceImpl struct {
	repo repositories.PortfolioAssetRepository
}

func NewPortfolioAssetServiceImpl(repo repositories.PortfolioAssetRepository) services.PortfolioAssetService {
	return &portfolioAssetServiceImpl{
		repo: repo,
	}
}

func (p *portfolioAssetServiceImpl) CreatePortfolioAsset(req models.AddAPortfolioAssetRequest) (entities.PortfolioAsset, error) {
	portfolioAsset, err := entities.NewPortfolioAsset(req.PortfolioID, req.AssetID, req.Quantity, req.AverageBuyPrice, req.CurrentValue)
	if err != nil {
		return entities.PortfolioAsset{}, err
	}

	// current_value = quantity * latest_price
	// portfolioAsset.CurrentValue = ...

	createPortfolioAsset, err := p.repo.Create(*portfolioAsset)
	if err != nil {
		return entities.PortfolioAsset{}, err
	}

	return createPortfolioAsset, nil
}

func (p *portfolioAssetServiceImpl) UpdatePortfolioAsset(req models.UpdatePortfolioAssetRequest) (entities.PortfolioAsset, error) {
	portfolioAsset, err := p.repo.FindById(req.ID)
	if err != nil {
		return entities.PortfolioAsset{}, err
	}

	if req.Quantity != "" {
		portfolioAsset.Quantity = req.Quantity
	}
	if req.AverageBuyPrice != "" {
		portfolioAsset.AverageBuyPrice = req.AverageBuyPrice
	}
	if req.CurrentValue != "" {
		portfolioAsset.CurrentValue = req.CurrentValue
	}

	// current_value = quantity * latest_price
	// portfolioAsset.CurrentValue = ...

	updatePortfolioAsset, err := p.repo.Update(*portfolioAsset)
	if err != nil {
		return entities.PortfolioAsset{}, err
	}

	return updatePortfolioAsset, nil
}

func (p *portfolioAssetServiceImpl) GetPortfolioAssetById(id string) (entities.PortfolioAsset, error) {
	portfolioAsset, err := p.repo.FindById(id)
	if err != nil {
		return entities.PortfolioAsset{}, err
	}
	return *portfolioAsset, nil
}

func (p *portfolioAssetServiceImpl) DeletePortfolioAsset(id string) error {
	return p.repo.Delete(id)
}
