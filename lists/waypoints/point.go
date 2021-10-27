package waypoints

import "fmt"

// A Geo-Point.
type Point struct {
	Lat  float64
	Long float64
}

func (p *Point) String() string {
	return fmt.Sprintf("[Latitude: %.4f; Longitude: %.4f]", p.Lat, p.Long)
}
