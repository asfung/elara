package impl

import (
	"github.com/asfung/elara/internal/entities"
	"github.com/asfung/elara/internal/models"
	"github.com/asfung/elara/internal/repositories"
	"github.com/asfung/elara/internal/services"
)

type cardServiceImpl struct {
	repo        repositories.CardRepository
	userService services.UserService
}

func NewCardServiceImpl(
	repo repositories.CardRepository,
	userService services.UserService,
) services.CardService {
	return &cardServiceImpl{
		repo:        repo,
		userService: userService,
	}
}

func (c *cardServiceImpl) CreateCard(req models.AddCardRequest) (entities.Card, error) {
	_, err := c.userService.GetUserById(req.UserID)
	if err != nil {
		return entities.Card{}, err
	}

	card, err := entities.NewCard(req.UserID, req.CardNumberHash, req.CardType, req.ExpiryDate)
	if err != nil {
		return entities.Card{}, err
	}

	createdCard, err := c.repo.Create(*card)
	if err != nil {
		return entities.Card{}, err
	}

	return createdCard, nil
}

func (c *cardServiceImpl) UpdateCard(req models.UpdateCardRequest) (entities.Card, error) {
	card, err := c.repo.FindById(req.ID)
	if err != nil {
		return entities.Card{}, err
	}

	if req.CardNumberHash != "" {
		card.CardNumberHash = req.CardNumberHash
	}
	if req.CardType != "" {
		card.CardType = req.CardType
	}
	if req.ExpiryDate != "" {
		card.ExpiryDate = req.ExpiryDate
	}
	if req.Status != "" {
		card.Status = req.Status
	}
	if req.TokenizedReference != "" {
		card.TokenizedReference = req.TokenizedReference
	}

	updatedCard, err := c.repo.Update(*card)
	if err != nil {
		return entities.Card{}, err
	}

	return updatedCard, nil
}

func (c *cardServiceImpl) GetCardById(id string) (entities.Card, error) {
	card, err := c.repo.FindById(id)
	if err != nil {
		return entities.Card{}, err
	}
	return *card, nil
}

func (c *cardServiceImpl) DeleteCard(id string) error {
	return c.repo.Delete(id)
}

func (c *cardServiceImpl) GetCardByUserId(userID string) (entities.Card, error) {
	card, err := c.repo.FindByUserId(userID)
	if err != nil {
		return entities.Card{}, err
	}
	return card, nil
}
