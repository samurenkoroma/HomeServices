package spatial

/*
query := `
INSERT INTO fields(id,geom)
VALUES($1,` + spatial.GeoJSONToPostGIS() + `)
*/

func GeoJSONToPostGIS() string {

	return `
        ST_SetSRID(
            ST_GeomFromGeoJSON($1),
            4326
        )
    `
}
