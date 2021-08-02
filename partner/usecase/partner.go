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
func (uc *PartnerUseCase) GetNearestPartner(ctx context.Context, ID string) (*e.Partner, error) {
	/*p := NewPoint(49.014, 8.4043)
    geocoder := new(geo.Point)
    data, err := geocoder.Request(fmt.Sprintf("latlng=%f,%f", p.Lat(), p.Lng()))
    if err != nil {
        log.Println(err)
    }
    var res googleGeocodeResponse
    if err := json.Unmarshal(data, &res); err != nil {
        log.Println(err)
    }
    var city string
    if len(res.Results) > 0 {
        r := res.Results[0]
    outer:
        for _, comp := range r.AddressComponents {
            // See https://developers.google.com/maps/documentation/geocoding/#Types
            // for address types
            for _, compType := range comp.Types {
                if compType == "locality" {
                    city = comp.LongName
                    break outer
                }
            }
        }
    }
    fmt.Printf("City: %s\n", city)*/
	return nil, nil
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
                Message: "partner document already taken",
                Err: err,
            }
		}
        return nil, err
    }

    p.ID = *id

    return p, nil
}