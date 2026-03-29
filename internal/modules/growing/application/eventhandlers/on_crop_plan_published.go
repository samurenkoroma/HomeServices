package eventhandlers

import (
	"context"
	"fmt"
	"log"
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/croptemplate"
)

// OnCropPlanPublished — обработчик события публикации плана культуры
func OnCropPlanPublished(ctx context.Context, e event.DomainEvent) error {
	event, ok := e.(cropplan.CropPlanPublished)
	if !ok {
		return fmt.Errorf("on_crop_plan_published event is not of type CropPlanPublished")
	}

	uow, ok := repository.FromContext(ctx)
	if !ok {
		return nil
	}

	growingProvider, ok := postgres.NewGrowingProvider(uow.Tx()).(*postgres.GrowingProvider)
	if !ok {
		return fmt.Errorf("invalid provider type")
	}

	// Получаем полную информацию о плане (нужно расширить событие)
	// В реальности событие должно содержать все необходимые данные
	// или нужно делать дополнительный запрос к crop модулю

	// Создаём шаблон
	template := croptemplate.NewCropTemplate(
		event.PlanID,
		event.Name,
		event.Version,
	)

	// Добавляем этапы из события
	for _, stage := range event.Stages {
		growthStage := croptemplate.GrowthStage{
			Order:           stage.Order,
			Name:            stage.Name,
			Duration:        int(stage.Duration),
			Recommendations: stage.Recommendations.ToMap(),
			Description:     stage.Description,
		}
		if err := template.AddStage(growthStage); err != nil {
			return err
		}
	}

	// Устанавливаем требования
	template.SetRequirements(croptemplate.Requirements{
		MinPH:      event.Environment.MinPH,
		MaxPH:      event.Environment.MaxPH,
		Nitrogen:   event.Nutrients.Nitrogen,
		Phosphorus: event.Nutrients.Phosphorus,
		Potassium:  event.Nutrients.Potassium,
		SoilTypes:  event.Environment.SoilTypes,
	})

	// Публикуем шаблон
	if err := template.Publish(); err != nil {
		return err
	}

	// Сохраняем
	if err := growingProvider.CropTemplates().Save(ctx, template); err != nil {
		return err
	}

	uow.RegisterAggregate(template)

	log.Printf("Crop template created from plan %s version %d", event.PlanID, event.Version)
	return nil
}
