package spatial

func MakePoint(lon, lat float64) Coordinate {

	return Coordinate{
		Lon: lon,
		Lat: lat,
	}
}

func MakePolygon(coords [][]Coordinate) GeoJSON {

	return GeoJSON{
		//Coordinates: coords,
		Type: Polygon,
	}
}
