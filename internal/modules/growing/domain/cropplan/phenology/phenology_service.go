package phenology

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"time"
)

type PhenologyService interface {
	GetCurrentPhenology(
		ctx context.Context,
		planID string,
		varietyID string,
		plantingDate time.Time,
		lat, lon float64,
	) (*CurrentPhenology, error)
}

// PhenologyService основной сервис для работы с фенологией
type phenologyService struct {
	catalogRepo     catalog.Repository
	weatherProvider WeatherProvider
	gddCalculator   *GDDCalculator
	useSinusoidal   bool
}

func NewPhenologyService(
	catalogRepo catalog.Repository,
	weatherProvider WeatherProvider,
	useSinusoidal bool,
) PhenologyService {
	return &phenologyService{
		catalogRepo:     catalogRepo,
		weatherProvider: weatherProvider,
		gddCalculator:   NewGDDCalculator(),
		useSinusoidal:   useSinusoidal, // больше не передаем температуры
	}
}

func (s *phenologyService) GetCurrentPhenology(
	ctx context.Context,
	planID string,
	varietyID string,
	plantingDate time.Time,
	lat, lon float64,
) (*CurrentPhenology, error) {
	today := time.Now()
	// 1. Проверяем дату посадки
	if plantingDate.After(today) {
		return nil, ErrPlantingDateInFuture
	}

	// 2. Получаем сорт из каталога
	variety, err := s.catalogRepo.GetVariety(ctx, varietyID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrVarietyNotFound, err)
	}

	// 3. Получаем погоду с даты посадки
	temps, err := s.weatherProvider.GetHistoricalTemperatures(
		ctx, lat, lon, plantingDate, today,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrNoWeatherData, err)
	}

	// 4. Рассчитываем GDD с температурами из сорта!
	var accumulatedGDD float64
	if s.useSinusoidal {
		accumulatedGDD = s.gddCalculator.AccumulateGDDWithSinusoidal(
			temps, variety.BaseTemperature, variety.MaxTemperature,
		)
	} else {
		accumulatedGDD = s.gddCalculator.AccumulateGDD(
			temps, variety.BaseTemperature, variety.MaxTemperature,
		)
	}

	// 5. Определяем текущую фазу
	currentPhase := variety.GetPhaseByGDD(accumulatedGDD)
	if currentPhase == nil {
		return nil, ErrNoPhenophaseData
	}

	// 6. Определяем следующую фазу
	nextPhase := variety.GetNextPhase(accumulatedGDD)

	var requiredGDD float64
	var progress float64
	if nextPhase != nil {
		requiredGDD = nextPhase.GDDRequired
		progress = (accumulatedGDD / requiredGDD) * 100
		if progress > 100 {
			progress = 100
		}
	} else {
		// Все фазы пройдены
		progress = 100
		requiredGDD = accumulatedGDD
	}

	// 7. Прогнозируем дни до следующей фазы
	var estimatedDays int
	if nextPhase != nil {
		estimatedDays = s.gddCalculator.PredictDaysToTarget(
			accumulatedGDD,
			nextPhase.GDDRequired,
			temps,
			variety.BaseTemperature, // ← из сорта
			variety.MaxTemperature,  // ← из сорта
			s.useSinusoidal,
		)
	}

	// 8. Рассчитываем отклонение от идеала
	deviationDays, deviationReason := s.calculateDeviation(
		variety, plantingDate, today, accumulatedGDD, temps,
	)

	// 9. Генерируем рекомендации
	actions := s.generateRecommendations(
		currentPhase.Code,
		deviationDays,
		nextPhase,
	)

	// 10. Прогнозируем дату сбора урожая
	var harvestDate *time.Time
	lastPhase := variety.PhenophaseGDD[len(variety.PhenophaseGDD)-1]
	if lastPhase.GDDRequired > accumulatedGDD {
		daysToHarvest := s.gddCalculator.PredictDaysToTarget(
			accumulatedGDD,
			lastPhase.GDDRequired,
			temps,
			variety.BaseTemperature,
			variety.MaxTemperature,
			s.useSinusoidal,
		)
		if daysToHarvest > 0 && daysToHarvest < 365 {
			date := time.Now().AddDate(0, 0, daysToHarvest)
			harvestDate = &date
		}
	}

	return &CurrentPhenology{
		PlanID:                   planID,
		VarietyID:                varietyID,
		VarietyName:              variety.Name,
		AccumulatedGDD:           accumulatedGDD,
		RequiredGDDForNext:       requiredGDD,
		CurrentPhaseCode:         currentPhase.Code,
		CurrentPhaseName:         currentPhase.Name,
		ProgressPercent:          progress,
		EstimatedDaysToNextPhase: estimatedDays,
		EstimatedHarvestDate:     harvestDate,
		DeviationDays:            deviationDays,
		DeviationReason:          deviationReason,
		IsCritical:               currentPhase.IsCritical,
		RecommendedActions:       actions,
	}, nil
}

// calculateDeviation рассчитывает отклонение от идеального развития
func (s *phenologyService) calculateDeviation(
	variety *catalog.Variety,
	plantingDate, currentDate time.Time,
	actualGDD float64,
	temps []DailyTemp,
) (int, string) {

	daysSincePlanting := int(currentDate.Sub(plantingDate).Hours() / 24)
	if daysSincePlanting <= 0 {
		return 0, ""
	}

	// Идеальное GDD зависит от культуры!
	// Для томата 8 GDD/день, для огурца 10 GDD/день
	var idealDailyGDD float64
	switch variety.SpeciesKey {
	case "tomato":
		idealDailyGDD = 8.0
	case "cucumber":
		idealDailyGDD = 10.0
	case "eggplant":
		idealDailyGDD = 9.0
	default:
		idealDailyGDD = 8.0
	}

	idealGDD := float64(daysSincePlanting) * idealDailyGDD
	deviation := actualGDD - idealGDD
	deviationDays := int(deviation / idealDailyGDD)

	var reason string
	if deviationDays > 5 {
		reason = "heat_wave"
	} else if deviationDays < -5 {
		reason = "cold_spell"
	}

	return deviationDays, reason
}

// generateRecommendations генерирует рекомендации на основе текущей фазы
func (s *phenologyService) generateRecommendations(
	currentPhaseCode string,
	deviationDays int,
	nextPhase *catalog.PhenophaseGDD,
) []RecommendedAction {

	var actions []RecommendedAction

	switch currentPhaseCode {
	case "BBCH-10", "BBCH-19":
		actions = append(actions, RecommendedAction{
			Title:       "Осмотр всходов",
			Description: "Проверить густоту всходов, при необходимости проредить",
			Priority:    "medium",
			DueDays:     3,
		})

	case "BBCH-51":
		actions = append(actions, RecommendedAction{
			Title:       "Подготовка к цветению",
			Description: "Внести калийные удобрения для лучшего цветения",
			Priority:    "high",
			DueDays:     2,
		})

	case "BBCH-61":
		actions = append(actions, RecommendedAction{
			Title:       "Обработка от вредителей",
			Description: "Фаза цветения критическая для защиты от вредителей",
			Priority:    "urgent",
			DueDays:     1,
		})

		if deviationDays > 3 {
			actions = append(actions, RecommendedAction{
				Title:       "Усиленный полив",
				Description: "Из-за жары требуется дополнительное увлажнение",
				Priority:    "high",
				DueDays:     1,
			})
		}

	case "BBCH-71":
		actions = append(actions, RecommendedAction{
			Title:       "Нормирование завязей",
			Description: "Удалить лишние завязи для получения крупных плодов",
			Priority:    "medium",
			DueDays:     5,
		})

		if deviationDays < -3 {
			actions = append(actions, RecommendedAction{
				Title:       "Стимулятор роста",
				Description: "Задержка развития, требуется стимуляция",
				Priority:    "high",
				DueDays:     2,
			})
		}

	case "BBCH-81", "BBCH-89":
		actions = append(actions, RecommendedAction{
			Title:       "Подготовка к уборке",
			Description: "Проверить готовность плодов к сбору",
			Priority:    "high",
			DueDays:     3,
		})

		if nextPhase != nil && nextPhase.Code == "BBCH-89" {
			actions = append(actions, RecommendedAction{
				Title:       "Сбор урожая",
				Description: "Плоды достигли технической спелости",
				Priority:    "urgent",
				DueDays:     2,
			})
		}
	}

	return actions
}

//
//// ForecastDevelopment прогнозирует развитие на указанное количество дней
//func (s *PhenologyService) ForecastDevelopment(
//	ctx context.Context,
//	planID string,
//	varietyID string,
//	plantingDate time.Time,
//	lat, lon float64,
//	forecastDays int,
//) (*PhenologyForecast, error) {
//
//	// Получаем сорт
//	variety, err := s.catalogRepo.GetVariety(ctx, "", varietyID)
//	if err != nil {
//		return nil, err
//	}
//
//	// Получаем исторические температуры
//	today := time.Now()
//	historyTemps, err := s.weatherProvider.GetHistoricalTemperatures(
//		ctx, lat, lon, plantingDate, today,
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	// Получаем прогноз погоды
//	forecastTemps, err := s.weatherProvider.GetForecast(ctx, lat, lon, forecastDays)
//	if err != nil {
//		return nil, err
//	}
//
//	// Объединяем историю и прогноз
//	allTemps := append(historyTemps, forecastTemps...)
//
//	// Рассчитываем накопленное GDD
//	accumulatedGDD := s.gddCalculator.AccumulateGDD(historyTemps, variety.BaseTemperature, // ← из сорта
//		variety.MaxTemperature, // ← из сорта
//	)
//
//	// Прогнозируем даты фаз
//	var phases []ForecastPhase
//
//	for _, phase := range variety.PhenophaseGDD {
//		if phase.GDDRequired <= accumulatedGDD {
//			// Фаза уже пройдена
//			continue
//		}
//
//		// Прогнозируем дату достижения фазы
//		daysNeeded := s.gddCalculator.PredictDaysToTarget(
//			accumulatedGDD, phase.GDDRequired, allTemps,
//		)
//		expectedDate := today.AddDate(0, 0, daysNeeded)
//
//		phases = append(phases, ForecastPhase{
//			PhaseCode:    phase.Code,
//			PhaseName:    phase.Name,
//			ExpectedDate: expectedDate,
//			GDDRequired:  phase.GDDRequired,
//			IsCritical:   phase.IsCritical,
//		})
//	}
//
//	// Генерируем рекомендации
//	currentPhase := variety.GetPhaseByGDD(accumulatedGDD)
//	var actions []RecommendedAction
//	if currentPhase != nil {
//		actions = s.generateRecommendations(currentPhase.Code, 0, nil)
//	}
//
//	return &PhenologyForecast{
//		PlanID:             planID,
//		PlantingDate:       plantingDate,
//		ForecastDate:       today,
//		Phases:             phases,
//		RecommendedActions: actions,
//	}, nil
//}
