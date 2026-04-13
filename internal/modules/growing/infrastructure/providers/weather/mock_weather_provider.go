package weather

import (
	"context"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/phenology"
	"time"
)

// MockWeatherProvider мок-провайдер для тестов и разработки
type MockWeatherProvider struct {
	// Генерирует идеальные температуры (20-25°C)
}

func NewMockWeatherProvider() phenology.WeatherProvider {
	return &MockWeatherProvider{}
}

func (m *MockWeatherProvider) GetHistoricalTemperatures(
	ctx context.Context,
	lat, lon float64,
	from, to time.Time,
) ([]phenology.DailyTemp, error) {
	var temps []phenology.DailyTemp

	days := int(to.Sub(from).Hours() / 24)
	for i := 0; i <= days; i++ {
		date := from.AddDate(0, 0, i)

		// Идеальная температура для роста
		// Весна/осень: прохладнее, лето: теплее
		month := date.Month()
		var baseTemp float64

		switch month {
		case 3, 4, 5: // весна
			baseTemp = 15.0
		case 6, 7, 8: // лето
			baseTemp = 22.0
		default: // осень/зима
			baseTemp = 12.0
		}

		temps = append(temps, phenology.DailyTemp{
			Date: date,
			Min:  baseTemp - 5,
			Max:  baseTemp + 5,
		})
	}

	return temps, nil
}

func (m *MockWeatherProvider) GetForecast(
	ctx context.Context,
	lat, lon float64,
	days int,
) ([]phenology.DailyTemp, error) {
	return m.GetHistoricalTemperatures(ctx, lat, lon, time.Now(), time.Now().AddDate(0, 0, days))
}

func (m *MockWeatherProvider) GetCurrentTemperature(
	ctx context.Context,
	lat, lon float64,
) (float64, error) {
	return 22.0, nil
}
