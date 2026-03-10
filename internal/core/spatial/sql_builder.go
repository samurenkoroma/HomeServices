package spatial

func AsGeoJSON(column string) string {

	return ` ST_AsGeoJSON(` + column + `)`
}

func SetSRID(column string) string {

	return ` ST_SetSRID(` + column + `,4326)`
}
