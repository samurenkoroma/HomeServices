package spatial

import "encoding/json"

func ParseGeoJSON(data []byte) (*GeoJSON, error) {

	var g GeoJSON

	if err := json.Unmarshal(data, &g); err != nil {
		return nil, err
	}

	return &g, nil
}
