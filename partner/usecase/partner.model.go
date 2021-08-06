package usecase

// CreatePartnerInput is the set of information that will be used to enter data through our handlers.
// We can understand it as a Command. It is used in CREATE operations.
// swagger:model CreatePartnerInput
type CreatePartnerInput struct {
	TradingName  string    `json:"tradingName"`
	OwnerName    string    `json:"ownerName"`
	Document     string    `json:"document"`
	CoverageArea *CoverageArea `json:"coverageArea"`
	Address  *Address `json:"address"`
}

// CoverageArea follows the GeoJSON MultiPolygon format (https://en.wikipedia.org/wiki/GeoJSON)
// swagger:model CoverageArea
type CoverageArea struct {
	Type        string      `json:"type" validate:"oneof=Point MultiPolygon"`
	Coordinates [][][][]float64 `json:"coordinates"`
}

// Address follows the GeoJSON Point format (https://en.wikipedia.org/wiki/GeoJSON)
// swagger:model Address
type Address struct {
	Type        string     `json:"type" validate:"oneof=Point MultiPolygon"`
	Coordinates []float64 `json:"coordinates"`
}