package valueobject

import (
	"errors"
)

var (
	// Общие ошибки
	ErrNegativeValue           = errors.New("value cannot be negative")
	ErrInvalidPeriod           = errors.New("start date must be before end date")
	ErrInvalidTemperatureRange = errors.New("min temperature cannot be greater than max temperature")
	ErrInvalidHumidityRange    = errors.New("min humidity cannot be greater than max humidity")
	ErrInvalidPHRange          = errors.New("min pH cannot be greater than max pH")
	ErrInvalidLightHours       = errors.New("light hours cannot be negative")
	ErrInvalidWaterRequirement = errors.New("water requirement cannot be negative")

	// Ошибки питательных веществ
	ErrNegativeNutrientValue = errors.New("nutrient value cannot be negative")

	// Ошибки урожайности
	ErrNegativeYield      = errors.New("yield value cannot be negative")
	ErrInvalidYieldRange  = errors.New("min yield cannot be greater than max yield")
	ErrExpectedOutOfRange = errors.New("expected yield must be between min and max")

	// Ошибки длины
	ErrInvalidDuration = errors.New("duration must be greater than 0")
	ErrInvalidDays     = errors.New("number of days must be greater than 0")
)
