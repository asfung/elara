package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type AssetService interface {
	CreateAsset(req models.AddAssetRequest) (entities.Asset, error)
	UpdateAsset(req models.UpdateAssetRequest) (entities.Asset, error)
	GetAssetById(id string) (entities.Asset, error)
	DeleteAsset(id string) error
}
