package weather_test

import (
	"strings"
	"testing"

	"samurenkoroma/services/internal/weather/geo"
	"samurenkoroma/services/internal/weather/weather"
)

func TestGetWeather(t *testing.T) {
	expected := "Odessa"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3

	result, err := weather.GetWeather(geoData, format)

	if err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалась погода в %v, а получено: %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Большой формат", format: 234},
	{name: "Отрицательный формат", format: -21},
	{name: "0 формат", format: 0},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Odessa"
			geoData := geo.GeoData{
				City: expected,
			}

			_, err := weather.GetWeather(geoData, tc.format)

			if err != weather.ErrWrongFormat {
				t.Error(err.Error())
			}
		})
	}
}
