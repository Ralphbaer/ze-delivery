package repository

import (
	"context"

	e "github.com/Ralphbaer/ze-delivery/partner/entity"
)

// PartnerRepository manages event repository operations
type PartnerRepository interface {
	Find(ctx context.Context, ID string) (*e.Partner, error)
	FindNearest(ctx context.Context, long, lat float64) (*e.Partner, error)
	Save(ctx context.Context, partner *e.Partner) (*string, error)
}
