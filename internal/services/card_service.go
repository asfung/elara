package services

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
)

type CardService interface {
	CreateCard(req models.AddCardRequest) (entities.Card, error)
	UpdateCard(req models.UpdateCardRequest) (entities.Card, error)
	GetCardById(id string) (entities.Card, error)
	DeleteCard(id string) error
	GetCardByUserId(userID string) (entities.Card, error)
}
