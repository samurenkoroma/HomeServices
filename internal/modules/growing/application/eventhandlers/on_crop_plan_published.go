package eventhandlers

import (
	"context"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/croptemplate"
)

// OnCropPlanPublished - реакция на публикацию плана культуры из crop модуля
func OnCropPlanPublished(ctx context.Context, cropPlanID string, name string, stages []crop.Stage, requirements crop.Requirements) error {
	uow, ok := repository.FromContext(ctx)
	if !ok {
		return nil
	}

	growingProvider := uow.(*postgres.GrowingProvider)

	// Создаем шаблон
	template := croptemplate.NewCropTemplate(cropPlanID, name)

	// Добавляем этапы
	for _, s := range stages {
		stage := croptemplate.GrowthStage{
			Order:       s.Order,
			Name:        s.Name,
			Duration:    s.Duration,
			MinTemp:     s.MinTemp,
			MaxTemp:     s.MaxTemp,
			MinHumidity: s.MinHumidity,
			MaxHumidity: s.MaxHumidity,
			WaterPerDay: s.WaterPerDay,
		}
		if err := template.AddStage(stage); err != nil {
			return err
		}
	}

	// Добавляем требования
	template.Requirements = croptemplate.Requirements{
		MinPH:      requirements.MinPH,
		MaxPH:      requirements.MaxPH,
		Nitrogen:   requirements.Nitrogen,
		Phosphorus: requirements.Phosphorus,
		Potassium:  requirements.Potassium,
		SoilType:   requirements.SoilType,
	}

	// Публикуем
	if err := template.Publish(); err != nil {
		return err
	}

	// Сохраняем
	if err := growingProvider.CropTemplates().Save(ctx, template); err != nil {
		return err
	}

	uow.RegisterAggregate(template)
	return nil
}
