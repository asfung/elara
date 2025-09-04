package repositories

import (
	"github.com/asfung/elara/database"
	"github.com/asfung/elara/internal/entities"
)

type P2PTransferPostgresRepository struct {
	*BaseRepository[entities.P2pTransfer]
}

func NewP2PTransferPostgresRepository(db database.Database) P2PTransferRepository {
	return &P2PTransferPostgresRepository{
		BaseRepository: NewBaseRepository[entities.P2pTransfer](db),
	}
}
