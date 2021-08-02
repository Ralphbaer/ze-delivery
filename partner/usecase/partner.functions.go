package usecase

// Represents a Physical Point in geographic notation [lat, lng].
type Point struct {
	lat float64
	lng float64
}

const (
	// According to Wikipedia, the Earth's radius is about 6,371km
	EARTH_RADIUS = 6371
)

// Returns a new Point populated by the passed in latitude (lat) and longitude (lng) values.
func NewPoint(lat float64, lng float64) *Point {
	return &Point{lat: lat, lng: lng}
}

// Returns Point p's latitude.
func (p *Point) Lat() float64 {
	return p.lat
}

// Returns Point p's longitude.
func (p *Point) Lng() float64 {
	return p.lng
}