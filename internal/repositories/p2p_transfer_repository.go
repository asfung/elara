package repositories

import "github.com/asfung/elara/internal/entities"

type P2PTransferRepository interface {
	Repository[entities.P2pTransfer]
}
