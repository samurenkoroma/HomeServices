package phenology

import "errors"

var (
	ErrNoWeatherData        = errors.New("no weather data available for period")
	ErrInvalidTemperature   = errors.New("invalid temperature values")
	ErrVarietyNotFound      = errors.New("variety not found in catalog")
	ErrPlantingDateInFuture = errors.New("planting date cannot be in future")
	ErrNoPhenophaseData     = errors.New("no phenophase data for this variety")
)
