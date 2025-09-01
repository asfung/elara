package repositories

import "github.com/asfung/elara/internal/entities"

type CardRepository interface {
	Repository[entities.Card]
}
