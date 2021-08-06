package repository

import (
	"context"

	e "github.com/Ralphbaer/ze-delivery/partner/entity"
)

//go:generate mockgen -destination=../gen/mock/repository_mock.go -package=mock . PartnerRepository

// PartnerRepository manages partner repository operations
type PartnerRepository interface {
	Find(ctx context.Context, ID string) (*e.Partner, error)
	FindByDocument(ctx context.Context, document string) (*e.Partner, error)
	FindNearest(ctx context.Context, long, lat float64) (*e.Partner, error)
	Save(ctx context.Context, partner *e.Partner) (*string, error)
}
