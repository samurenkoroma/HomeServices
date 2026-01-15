package geo_test

import (
	"testing"

	"samurenkoroma/services/internal/weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	got, err := geo.GetMylocation(city)

	if err != nil {
		t.Error(err.Error())
	}

	if got.City != expected.City {
		t.Errorf("Ожидалось %v | получено %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "NoCity"

	_, err := geo.GetMylocation(city)

	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v | получено %v", geo.ErrNoCity, err)
	}

}
