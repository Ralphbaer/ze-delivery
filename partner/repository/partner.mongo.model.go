package repository

import (
	"strings"

	"github.com/Ralphbaer/ze-delivery/common/zmongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	e "github.com/Ralphbaer/ze-delivery/partner/entity"
)

// PartnerQuery represents the query to find partners
type PartnerQuery struct {
	Longitude   float64       `schema:"long,required"`
	Latitude    float64       `schema:"lat,required"`
}

// PartnerMongoModel is the model of entity.Partner
type PartnerMongoModel struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty"`
	TradingName  string    `bson:"tradingName"`
	OwnerName    string    `bson:"ownerName"`
	Document     string    `bson:"document"`
	CoverageArea CoverageArea `bson:"coverageArea"`
	Address  Address `bson:"address"`
}

// CoverageArea follows the GeoJSON MultiPolygon format (https://en.wikipedia.org/wiki/GeoJSON)
type CoverageArea struct {
	Type        string          `bson:"type"`
	Coordinates [][][][]float64 `bson:"coordinates"`
}

// Address follows the GeoJSON Point format (https://en.wikipedia.org/wiki/GeoJSON)
type Address struct {
	Type        string     `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

// ToEntity converts a PartnerMongoModel to e.Partner
func (d *PartnerMongoModel) ToEntity() *e.Partner {
	partner := &e.Partner{
		ID:       d.ID.Hex(),
		TradingName: d.TradingName,
		OwnerName: d.OwnerName,
		Document: d.Document,
		CoverageArea: &e.CoverageArea{
			Type: d.CoverageArea.Type,
			Coordinates: d.CoverageArea.Coordinates,
		},
		Address: &e.Address{
			Type: d.Address.Type,
			Coordinates: d.Address.Coordinates,
		},
	}

	return partner
}

// FromEntity converts an entity.Partner to PartnerMongoModel
func (d *PartnerMongoModel) FromEntity(partner *e.Partner) error {
	if strings.TrimSpace(partner.ID) != "" {
		objID, err := zmongo.ObjectIDFromString(partner.ID)
		if err != nil {
			return err
		}
		d.ID = objID
	}

	d.TradingName = partner.TradingName
	d.OwnerName = partner.OwnerName
	d.Document = partner.Document
	d.CoverageArea.Type = partner.CoverageArea.Type
	d.CoverageArea.Coordinates = partner.CoverageArea.Coordinates
	d.Address.Type = partner.Address.Type
	d.Address.Coordinates = partner.Address.Coordinates

	return nil
}
