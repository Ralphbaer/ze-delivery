package entity

// Partner represents a collection of identification data about a Zé Delivery Partner,
// including its coordinates represented by the coverageArea and address fields.
// swagger:model Partner
type Partner struct {
	// ID is a unique field
	ID           string    `json:"id" bson:"_id"`
	TradingName  string    `json:"tradingName"`
	OwnerName    string    `json:"ownerName"`
	// Document is a unique field
	Document     string    `json:"document"`
	CoverageArea *CoverageArea `json:"coverageArea"`
	Address  *Address `json:"address"`
}

// CoverageArea follows the GeoJSON MultiPolygon format (https://en.wikipedia.org/wiki/GeoJSON)
type CoverageArea struct {
	Type        string      `json:"type"`
	Coordinates [][][][]float64 `json:"coordinates"`
}

// Address follows the GeoJSON Point format (https://en.wikipedia.org/wiki/GeoJSON)
type Address struct {
	Type        string      `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
