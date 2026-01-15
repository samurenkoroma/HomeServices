package geo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulation struct {
	Error bool `json:"error"`
}

var ErrNoCity = errors.New("NOCITY")
var ErrNot200 = errors.New("NOT200")

func GetMylocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			return nil, ErrNoCity
		}
		return &GeoData{
			City: city,
		}, nil

	}

	ipapiClient := http.Client{}
	req, err := http.NewRequest("GET", "https://ipapi.co/json/", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.5")
	resp, err := ipapiClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Print(resp)
		return nil, ErrNot200
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geoData GeoData

	json.Unmarshal(body, &geoData)
	return &geoData, nil
}

func checkCity(city string) bool {
	payload := strings.NewReader(`{
    "city": "lagos"
}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://countriesnow.space/api/v0.1/countries/population/cities", payload)
	if err != nil {
		return false
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var cityPopulation CityPopulation

	json.Unmarshal(body, &cityPopulation)
	return !cityPopulation.Error

}
