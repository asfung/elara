package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type P2PTransferService interface {
	CreateP2PTransfer(req models.AddP2PTransferRequest) (entities.P2pTransfer, error)
	UpdateP2PTransfer(req models.UpdateP2PTransferRequest) (entities.P2pTransfer, error)
	GetP2PTransferById(id string) (entities.P2pTransfer, error)
	DeleteP2PTransfer(id string) error
}
