package spatial

func ContainsQuery(table string) string {

	return `
    SELECT id
    FROM ` + table + `
    WHERE ST_Contains(
        geom,
        ST_SetSRID(ST_Point($1,$2),4326)
    )
    `
}

func IntersectsQuery(table string) string {

	return `
    SELECT *
    FROM ` + table + `
    WHERE ST_Intersects(
        geom,
        ST_GeomFromGeoJSON($1)
    )
    `
}

func WithinQuery(table string) string {

	return `
    SELECT *
    FROM ` + table + `
    WHERE ST_Within(
        geom,
        ST_GeomFromGeoJSON($1)
    )
    `
}
