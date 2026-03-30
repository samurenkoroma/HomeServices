package spatial

import (
	"database/sql"
	"encoding/json"
)

func CalculateAreaFromGeometry(tx *sql.Tx, geom GeoJSON) float64 {
	var area float64
	data, err := json.Marshal(geom)
	if err != nil {
		return 0
	}
	err = tx.QueryRow("SELECT ST_Area(ST_SetSRID(ST_GeomFromGeoJSON($1), 4326)::geography) / 10000", data).Scan(&area)
	if err != nil {
		return 0
	}
	return area
}

func CalculateCenter(tx *sql.Tx, geom GeoJSON) [2]float64 {
	var point = struct {
		Lat  float64 `db:"lat"`
		Long float64 `db:"lng"`
	}{}

	var area [2]float64
	data, err := json.Marshal(geom)
	if err != nil {
		return [2]float64{}
	}
	err = tx.QueryRow("select ST_X(ST_Centroid($1)) AS lng, ST_Y(ST_Centroid($1)) AS lat", data).Scan(&point)
	if err != nil {
		return [2]float64{}
	}
	return area
}
