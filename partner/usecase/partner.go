package usecase

import (
	"context"
	"log"

	"github.com/Ralphbaer/ze-delivery/common"
	e "github.com/Ralphbaer/ze-delivery/partner/entity"
	r "github.com/Ralphbaer/ze-delivery/partner/repository"
)

// PartnerUseCase represents a collection of use cases for partner operations
type PartnerUseCase struct {
	PartnerRepo r.PartnerRepository
}

// Get get a partner by its id
func (uc *PartnerUseCase) Get(ctx context.Context, ID string) (*e.Partner, error) {
	return uc.PartnerRepo.Find(ctx, ID)
}

// GetNearestPartner returns the nearest partner given coordinates longitude and latitude
func (uc *PartnerUseCase) GetNearestPartner(ctx context.Context, q *r.PartnerQuery) (*e.Partner, error) {
	result, err := uc.PartnerRepo.FindNearest(ctx, q.Longitude, q.Latitude)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, common.EntityNotFoundError{
			Message: ErrNoNearestPartner.Error(),
			Err: err,
		}
	}

	return result, nil
}

// Create creates a new Partner
func (uc *PartnerUseCase) Create(ctx context.Context, cpi *CreatePartnerInput) (*e.Partner, error) {
	p := &e.Partner{
		TradingName: cpi.TradingName,
		OwnerName:       cpi.OwnerName,
		Document: cpi.Document,
		CoverageArea: &e.CoverageArea{
            Type: cpi.CoverageArea.Type,
			Coordinates: cpi.CoverageArea.Coordinates,
		},
		Address: &e.Address{
            Type: cpi.Address.Type,
			Coordinates: cpi.Address.Coordinates,
		},
	}

	log.Printf("Creating event %+v", p)

	id, err := uc.PartnerRepo.Save(ctx, p)
    if err != nil {
        if err == common.ErrMongoDuplicatedDocument {
			return nil, common.EntityConflictError{
                Message: ErrPartnerDocumentConflict.Error(),
                Err: err,
            }
		}
        return nil, err
    }

    p.ID = *id

    return p, nil
}