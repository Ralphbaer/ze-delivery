package usecase

// Point represents a Physical Point in geographic notation [lat, lng].
type Point struct {
	lat float64
	lng float64
}

const (
	// EarthRadius is the radius of our earth. According to Wikipedia, EarthRadius is about 6,371km
	EarthRadius = 6371
)

// NewPoint returns a new Point populated by the passed in latitude (lat) and longitude (lng) values.
func NewPoint(lat float64, lng float64) *Point {
	return &Point{lat: lat, lng: lng}
}

// Lat returns Point p's latitude.
func (p *Point) Lat() float64 {
	return p.lat
}

// Lng returns Point p's longitude.
func (p *Point) Lng() float64 {
	return p.lng
}