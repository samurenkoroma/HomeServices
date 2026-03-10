package spatial

import "encoding/json"

type GeoJSON struct {
	Type        GeometryType    `json:"type"`
	Coordinates json.RawMessage `json:"coordinates"`
}

type Feature struct {
	Type       string                 `json:"type"`
	Geometry   GeoJSON                `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}
