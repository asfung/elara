package repositories

import "github.com/asfung/elara/internal/entities"

type RoleRepository interface {
	Repository[entities.Role]
}
