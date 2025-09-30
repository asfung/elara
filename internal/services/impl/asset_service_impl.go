package impl

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
)

type assetServiceImpl struct {
	repo repositories.AssetRepository
}

func NewAssetServiceImpl(assetRepo repositories.AssetRepository) services.AssetService {
	return &assetServiceImpl{
		repo: assetRepo,
	}
}

func (a *assetServiceImpl) CreateAsset(req models.AddAssetRequest) (entities.Asset, error) {
	asset, err := entities.NewAsset(req.Symbol, req.Type, req.Name, req.Exchange, req.Currency)
	if err != nil {
		return entities.Asset{}, err
	}

	createAsset, err := a.repo.Create(*asset)
	if err != nil {
		return entities.Asset{}, err
	}
	return createAsset, nil
}

func (a *assetServiceImpl) UpdateAsset(req models.UpdateAssetRequest) (entities.Asset, error) {
	asset, err := a.repo.FindById(req.ID)
	if err != nil {
		return entities.Asset{}, err
	}

	if req.Symbol != "" {
		asset.Symbol = req.Symbol
	}
	if req.Type != "" {
		asset.Type = req.Type
	}
	if req.Name != "" {
		asset.Name = req.Name
	}
	if req.Exchange != "" {
		asset.Exchange = req.Exchange
	}
	if req.Currency != "" {
		asset.Currency = req.Currency
	}

	updateAsset, err := a.repo.Update(*asset)
	if err != nil {
		return entities.Asset{}, err
	}

	return updateAsset, nil
}

func (a *assetServiceImpl) GetAssetById(id string) (entities.Asset, error) {
	asset, err := a.repo.FindById(id)
	if err != nil {
		return entities.Asset{}, err
	}
	return *asset, nil
}

func (a *assetServiceImpl) DeleteAsset(id string) error {
	return a.repo.Delete(id)
}
