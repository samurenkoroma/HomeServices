package payloads

type OpenMeteoResponse struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Hourly    Hourly  `json:"hourly"`
	Daily     Daily   `json:"daily"`
}

type Hourly struct {
	Datetime          []string  `json:"time"`
	WC                []int     `json:"weather_code"`
	Temperature_2m    []float32 `json:"temperature_2m"`
	WindSpeed_10m     []float32 `json:"wind_speed_10m"`
	WindDirection_10m []int     `json:"wind_direction_10m"`
	WindGusts_10m     []float32 `json:"wind_gusts_10m"`
	IsDay             []int     `json:"is_day"`
}

type Daily struct {
	Datetime          []string  `json:"time"`
	WC                []int     `json:"weather_code"`
	TemperatureMax_2m []float32 `json:"temperature_2m_max"`
	TemperatureMin_2m []float32 `json:"temperature_2m_min"`
	WindSpeed_10m     []float32 `json:"wind_speed_10m_max"`
	WindDirection_10m []int     `json:"wind_direction_10m_dominant"`
	WindGusts_10m     []float32 `json:"wind_gusts_10m_max"`
}

/*
https://api.open-meteo.com/v1/forecast?latitude=45.91&longitude=30.02&
// daily=weather_code,temperature_2m_max,temperature_2m_min,wind_direction_10m_dominant,wind_gusts_10m_max,wind_speed_10m_max&
// hourly=temperature_2m,weather_code,wind_speed_10m,wind_direction_10m,wind_gusts_10m&timezone=Europe%2FMoscow
*/
type OpenMeteoQuery struct {
	Latitude     float32 `qs:"latitude"`
	Longitude    float32 `qs:"longitude"`
	Daily        string  `qs:"daily"`
	Hourly       string  `qs:"hourly"`
	Timezone     string  `qs:"timezone"`
	ForecastDays int     `qs:"forecast_days"`
}

type HourlyItem struct {
	Temperature float32
	WC          int
	Wind        Wind
	Date        string
	IsDay       bool
}

type DailyItem struct {
	Temperature Temperature
	WC          int
	Wind        Wind
	Date        string
}

type Temperature struct {
	Min float32
	Max float32
}

type Wind struct {
	Speed     float32
	Direction int
	Gusts     float32
}

type WeatherData struct {
	Daily  []DailyItem
	Hourly []HourlyItem
}
