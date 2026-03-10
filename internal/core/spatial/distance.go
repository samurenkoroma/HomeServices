package spatial

func DistanceSQL() string {

	return `
        ST_Distance(
            geom,
            ST_SetSRID(ST_Point($1,$2),4326)
        )
    `
}
