package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"samurenkoroma/services/internal/weather/geo"
)

var ErrWrongFormat = errors.New("WRONG_FORMAT")

func GetWeather(geo geo.GeoData, format int) (string, error) {

	if format < 1 || format > 4 {
		return "", ErrWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", errors.New("WRONG_URL")
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))

	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Print(resp)
		return "nil", errors.New("NOT200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "nil", err
	}
	return string(body), nil
}
