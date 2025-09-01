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
