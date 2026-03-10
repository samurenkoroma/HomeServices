package spatial

import "database/sql"

func ValidateGeometry(g GeoJSON) error {

	switch g.Type {

	case Point:
		return validatePoint(g)

	case Polygon:
		return validatePolygon(g)

	case MultiPolygon:
		return validateMultiPolygon(g)

	default:
		return ErrUnsupportedGeometry
	}
}

func validatePoint(g GeoJSON) error {
	return nil
}

func validateMultiPolygon(g GeoJSON) error {
	return nil
}

func validatePolygon(g GeoJSON) error {

	if len(g.Coordinates) == 0 {
		return ErrInvalidGeometry
	}

	return nil
}

func ValidateBlockInsideField(
	tx *sql.Tx,
	fieldGeom string,
	blockGeom string,
) (bool, error) {

	query := `
    SELECT ST_Contains(
        ST_GeomFromGeoJSON($1),
        ST_GeomFromGeoJSON($2)
    )
    `

	var ok bool

	err := tx.QueryRow(query, fieldGeom, blockGeom).
		Scan(&ok)

	return ok, err
}
