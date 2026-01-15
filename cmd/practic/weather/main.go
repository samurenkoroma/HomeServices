package main

import (
	"flag"
	"fmt"

	"samurenkoroma/services/internal/weather/geo"
	"samurenkoroma/services/internal/weather/weather"
)

func main() {

	city := flag.String("city", "", "Выберите город")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	data, err := geo.GetMylocation(*city)
	if err != nil {
		panic(err)
	}

	weather, err := weather.GetWeather(*data, *format)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Погода в  %s:  %s", data.City, weather)
}
