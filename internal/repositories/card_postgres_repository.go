package repositories

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

type CardPostrgresRepository struct {
	*BaseRepository[entities.Card]
}

func NewCardPostrgresRepository(db database.Database) CardRepository {
	return &CardPostrgresRepository{
		BaseRepository: NewBaseRepository[entities.Card](db),
	}
}

func (c *CardPostrgresRepository) FindByUserId(userID string) (entities.Card, error) {
	var card entities.Card
	if err := c.db.GetDb().Where("user_id = ?", userID).First(&card).Error; err != nil {
		return entities.Card{}, err
	}
	return card, nil
}
