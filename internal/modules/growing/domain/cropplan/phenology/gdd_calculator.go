package phenology

import (
	"math"
)

// GDDCalculator рассчитывает накопленные градусо-дни
// Использует метод Baskerville-Emin для более точного расчета
type GDDCalculator struct {
	baseTemperature float64 // базовая температура (ниже которой рост останавливается)
	maxTemperature  float64 // максимальная температура для расчета
}

// NewGDDCalculator создает новый калькулятор GDD
// baseTemp: базовая температура (для томатов 10°C, для огурцов 12°C)
// maxTemp: максимальная температура (обычно 30°C)
func NewGDDCalculator(baseTemp, maxTemp float64) *GDDCalculator {
	return &GDDCalculator{
		baseTemperature: baseTemp,
		maxTemperature:  maxTemp,
	}
}

// DailyGDD рассчитывает GDD за один день
// Формула: GDD = (Tmin + Tmax)/2 - Tbase
// С ограничениями: если Tmin < Tbase, то Tmin = Tbase
//
//	если Tmax > Tmax, то Tmax = Tmax
func (c *GDDCalculator) DailyGDD(tMin, tMax float64) float64 {
	// Ограничиваем температуры
	if tMin < c.baseTemperature {
		tMin = c.baseTemperature
	}
	if tMax > c.maxTemperature {
		tMax = c.maxTemperature
	}

	// Средняя температура
	tAvg := (tMin + tMax) / 2

	// GDD не может быть отрицательным
	gdd := tAvg - c.baseTemperature
	if gdd < 0 {
		return 0
	}

	// Округляем до 1 десятичного знака
	return math.Round(gdd*10) / 10
}

// DailyGDDWithSinusoidal рассчитывает GDD с использованием синусоидальной кривой
// Более точный метод, учитывающий, что температура в течение дня меняется не линейно
func (c *GDDCalculator) DailyGDDWithSinusoidal(tMin, tMax float64) float64 {
	// Ограничиваем температуры
	if tMin < c.baseTemperature {
		tMin = c.baseTemperature
	}
	if tMax > c.maxTemperature {
		tMax = c.maxTemperature
	}

	// Синусоидальная аппроксимация
	// GDD = ∫(T(t) - Tbase)dt за день
	// Приближенно: (Tmax + Tmin)/2 - Tbase + (Tmax - Tmin)/(2π)

	tAvg := (tMin + tMax) / 2
	amplitude := (tMax - tMin) / 2

	// Коррекция на синусоидальность
	correction := amplitude / math.Pi

	gdd := (tAvg - c.baseTemperature) + correction
	if gdd < 0 {
		return 0
	}

	return math.Round(gdd*10) / 10
}

// AccumulateGDD суммирует GDD за период
func (c *GDDCalculator) AccumulateGDD(temps []DailyTemp) float64 {
	var total float64
	for _, t := range temps {
		total += c.DailyGDD(t.Min, t.Max)
	}
	return math.Round(total*10) / 10
}

// AccumulateGDDWithSinusoidal суммирует GDD с синусоидальным методом
func (c *GDDCalculator) AccumulateGDDWithSinusoidal(temps []DailyTemp) float64 {
	var total float64
	for _, t := range temps {
		total += c.DailyGDDWithSinusoidal(t.Min, t.Max)
	}
	return math.Round(total*10) / 10
}

// PredictDaysToTarget прогнозирует количество дней до достижения целевого GDD
func (c *GDDCalculator) PredictDaysToTarget(
	currentGDD, targetGDD float64,
	recentTemps []DailyTemp,
) int {
	if targetGDD <= currentGDD {
		return 0
	}

	// Берем последние 7 дней для расчета средней скорости
	days := 7
	if len(recentTemps) < days {
		days = len(recentTemps)
	}
	if days == 0 {
		return 14 // дефолтное значение, если нет данных
	}

	var sum float64
	for i := len(recentTemps) - days; i < len(recentTemps); i++ {
		sum += c.DailyGDD(recentTemps[i].Min, recentTemps[i].Max)
	}
	avgDailyGDD := sum / float64(days)

	if avgDailyGDD <= 0 {
		return 14
	}

	remainingGDD := targetGDD - currentGDD
	daysNeeded := int(math.Ceil(remainingGDD / avgDailyGDD))

	if daysNeeded < 1 {
		return 1
	}
	return daysNeeded
}
