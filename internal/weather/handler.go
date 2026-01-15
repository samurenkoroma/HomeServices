package weather

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"samurenkoroma/services/internal/app"
	"samurenkoroma/services/pkg/payloads"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sonh/qs"
)

type WeatherHandler struct {
	router fiber.Router
}

func NewWeatherHandler(app *app.Polevod) {
	h := &WeatherHandler{
		router: app.App,
	}
	h.router.Get("/weather", h.weather)
}

func (h *WeatherHandler) weather(c *fiber.Ctx) error {
	query := &payloads.OpenMeteoQuery{
		Latitude:     45.91,
		Longitude:    30.02,
		Hourly:       "temperature_2m,weather_code,wind_speed_10m,wind_direction_10m,wind_gusts_10m,is_day",
		Daily:        "weather_code,temperature_2m_max,temperature_2m_min,wind_direction_10m_dominant,wind_gusts_10m_max,wind_speed_10m_max",
		Timezone:     "Europe/Moscow",
		ForecastDays: 16,
	}

	encoder := qs.NewEncoder()
	values, err := encoder.Values(query)

	if err != nil {
		// Handle error
	}

	baseUrl, err := url.Parse("https://api.open-meteo.com/v1/forecast")
	if err != nil {
		return errors.New("WRONG_URL")
	}
	baseUrl.RawQuery = values.Encode()
	log.Println(baseUrl.String())
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var response payloads.OpenMeteoResponse

	json.Unmarshal(body, &response)

	var items payloads.WeatherData
	for i, value := range response.Hourly.Datetime {
		datetime, _ := time.Parse("2006-01-02T15:04", value)
		// if datetime.Format("2006-01-02") != time.Now().Format("2006-01-02") {
		// 	continue
		// }
		items.Hourly = append(items.Hourly, payloads.HourlyItem{
			WC:          response.Hourly.WC[i],
			Date:        datetime.Format("15:04"),
			Temperature: response.Hourly.Temperature_2m[i],
			IsDay:       response.Hourly.IsDay[i] == 1,
			Wind: payloads.Wind{
				Speed:     response.Hourly.WindSpeed_10m[i],
				Direction: response.Hourly.WindDirection_10m[i],
				Gusts:     response.Hourly.WindGusts_10m[i],
			},
		})
	}
	for i, value := range response.Daily.Datetime {
		datetime, _ := time.Parse("2006-01-02", value)
		items.Daily = append(items.Daily, payloads.DailyItem{
			WC:   response.Daily.WC[i],
			Date: datetime.Format("2006-01-02"),
			Temperature: payloads.Temperature{
				Max: response.Daily.TemperatureMax_2m[i],
				Min: response.Daily.TemperatureMin_2m[i],
			}, Wind: payloads.Wind{
				Speed:     response.Daily.WindSpeed_10m[i],
				Direction: response.Daily.WindDirection_10m[i],
				Gusts:     response.Daily.WindGusts_10m[i],
			},
		})
	}
	// component := pages.Weather(
	// 	[]string{"01.07", "02.07", "03.07", "04.07", "05.07", "06.07", "07.07"},
	// 	items,
	// )
	return c.JSON(items)
}
