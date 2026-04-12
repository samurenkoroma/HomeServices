package phenology

import (
	"context"
	"time"
)

// WeatherProvider интерфейс для получения погодных данных
// Может быть реализован для разных источников:
// - OpenMeteo (бесплатный)
// - Яндекс.Погода
// - Tomorrow.io
// - Тестовый мок
type WeatherProvider interface {
	// GetHistoricalTemperatures возвращает исторические температуры за период
	GetHistoricalTemperatures(
		ctx context.Context,
		lat, lon float64,
		from, to time.Time,
	) ([]DailyTemp, error)

	// GetForecast возвращает прогноз погоды на указанное количество дней
	GetForecast(
		ctx context.Context,
		lat, lon float64,
		days int,
	) ([]DailyTemp, error)

	// GetCurrentTemperature возвращает текущую температуру
	GetCurrentTemperature(
		ctx context.Context,
		lat, lon float64,
	) (float64, error)
}
