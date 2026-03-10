package spatial

type GeometryType string

const (
	Point        GeometryType = "Point"
	LineString   GeometryType = "LineString"
	Polygon      GeometryType = "Polygon"
	MultiPolygon GeometryType = "MultiPolygon"
)

type Coordinate struct {
	Lon float64
	Lat float64
}

type BoundingBox struct {
	MinLon float64
	MinLat float64
	MaxLon float64
	MaxLat float64
}
