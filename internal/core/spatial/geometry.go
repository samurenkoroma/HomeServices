package spatial

import "encoding/json"

type Geometry struct {
	Type        GeometryType
	Coordinates json.RawMessage
}

func NewGeometry(typ GeometryType, coords json.RawMessage) Geometry {
	return Geometry{
		Type:        typ,
		Coordinates: coords,
	}
}
