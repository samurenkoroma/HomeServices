package phenology

import (
	"math"
)

// GDDCalculator рассчитывает накопленные градусо-дни
type GDDCalculator struct {
}

// NewGDDCalculator создает новый калькулятор GDD
func NewGDDCalculator() *GDDCalculator {
	return &GDDCalculator{}
}

// DailyGDD принимает температуры как параметры
func (c *GDDCalculator) DailyGDD(tMin, tMax, baseTemp, maxTemp float64) float64 {
	// Ограничиваем температуры
	if tMin < baseTemp {
		tMin = baseTemp
	}
	if tMax > maxTemp {
		tMax = maxTemp
	}

	tAvg := (tMin + tMax) / 2
	gdd := tAvg - baseTemp

	if gdd < 0 {
		return 0
	}
	return math.Round(gdd*10) / 10
}

// DailyGDDWithSinusoidal рассчитывает GDD с использованием синусоидальной кривой
// Более точный метод, учитывающий, что температура в течение дня меняется не линейно
// Формула: GDD = (Tmin+Tmax)/2 - Tbase + (Tmax-Tmin)/(2π)
func (c *GDDCalculator) DailyGDDWithSinusoidal(tMin, tMax, baseTemp, maxTemp float64) float64 {
	// Ограничиваем температуры
	if tMin < baseTemp {
		tMin = baseTemp
	}
	if tMax > maxTemp {
		tMax = maxTemp
	}

	// Средняя температура
	tAvg := (tMin + tMax) / 2

	// Амплитуда
	amplitude := (tMax - tMin) / 2

	// Синусоидальная коррекция
	// Интеграл синусоиды за полпериода = 2 * amplitude / π
	// Для целого дня: amplitude / π
	sinusoidalCorrection := amplitude / math.Pi

	// GDD с коррекцией
	gdd := (tAvg - baseTemp) + sinusoidalCorrection

	if gdd < 0 {
		return 0
	}
	return math.Round(gdd*10) / 10
}

// AccumulateGDD суммирует GDD за период
func (c *GDDCalculator) AccumulateGDD(temps []DailyTemp, baseTemp, maxTemp float64) float64 {
	var total float64
	for _, t := range temps {
		total += c.DailyGDD(t.Min, t.Max, baseTemp, maxTemp)
	}
	return math.Round(total*10) / 10
}

// AccumulateGDDWithSinusoidal суммирует GDD с синусоидальным методом
func (c *GDDCalculator) AccumulateGDDWithSinusoidal(temps []DailyTemp, baseTemp, maxTemp float64) float64 {
	var total float64
	for _, t := range temps {
		total += c.DailyGDDWithSinusoidal(t.Min, t.Max, baseTemp, maxTemp)
	}
	return math.Round(total*10) / 10
}

// PredictDaysToTarget прогнозирует количество дней до достижения целевого GDD
func (c *GDDCalculator) PredictDaysToTarget(
	currentGDD, targetGDD float64,
	recentTemps []DailyTemp,
	baseTemp, maxTemp float64,
	useSinusoidal bool,
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
		return 14
	}

	var sum float64
	for i := len(recentTemps) - days; i < len(recentTemps); i++ {
		if useSinusoidal {
			sum += c.DailyGDDWithSinusoidal(recentTemps[i].Min, recentTemps[i].Max, baseTemp, maxTemp)
		} else {
			sum += c.DailyGDD(recentTemps[i].Min, recentTemps[i].Max, baseTemp, maxTemp)
		}
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
