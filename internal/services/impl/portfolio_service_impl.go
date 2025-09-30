package impl

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
)

type portfolioServiceImpl struct {
	repo repositories.PortfolioRepository
}

func NewPortfolioServiceImpl(repo repositories.PortfolioRepository) services.PortfolioService {
	return &portfolioServiceImpl{
		repo: repo,
	}
}

func (p *portfolioServiceImpl) CreatePortfolio(req models.AddAPortfolioRequest) (entities.Portfolio, error) {
	portfolio, err := entities.NewPortfolio(req.UserID, req.Name, req.Type)
	if err != nil {
		return entities.Portfolio{}, err
	}

	createPortfolio, err := p.repo.Create(*portfolio)
	if err != nil {
		return entities.Portfolio{}, err
	}

	return createPortfolio, nil
}

func (p *portfolioServiceImpl) UpdatePortfolio(req models.UpdatePortfolioRequest) (entities.Portfolio, error) {
	portfolio, err := p.repo.FindById(req.ID)
	if err != nil {
		return entities.Portfolio{}, err
	}

	if req.Name != "" {
		portfolio.Name = req.Name
	}

	if req.Type != "" {
		portfolio.Type = req.Type
	}

	updatePortfolio, err := p.repo.Update(*portfolio)
	if err != nil {
		return entities.Portfolio{}, err
	}

	return updatePortfolio, nil
}

func (p *portfolioServiceImpl) GetPortfolioById(id string) (entities.Portfolio, error) {
	portfolio, err := p.repo.FindById(id)
	if err != nil {
		return entities.Portfolio{}, err
	}
	return *portfolio, nil
}

func (p *portfolioServiceImpl) DeletePortfolio(id string) error {
	return p.repo.Delete(id)
}
