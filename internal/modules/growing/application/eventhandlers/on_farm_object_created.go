package eventhandlers

import (
	"context"
	"fmt"
	"log"
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

// OnFarmObjectCreated — обработчик события создания объекта в farm модуле
func OnFarmObjectCreated(ctx context.Context, event event.DomainEvent) error {
	factory, ok := repository.FactoryFromContext(ctx)
	if !ok {
		return nil
	}
	uow, err := factory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	return uow.Execute(ctx, postgres.NewGrowingProvider, func(provider repository.RepositoryProvider) error {
		growingProvider, ok := provider.(*postgres.GrowingProvider)
		if !ok {
			return fmt.Errorf("invalid provider type")
		}
		var area cultivationarea.CultivationArea

		switch e := event.(type) {
		case physicalobject.FieldCreated:
			// Создано поле → создаём FieldArea
			area = cultivationarea.NewFieldArea(
				e.ID,
				e.Name,
				e.Geometry,
				e.Area,
			)
			log.Printf("Created FieldArea with geometry: %+v", e.Geometry)

		case physicalobject.GreenhouseCreated:
			// Создана теплица → создаём GreenhouseArea
			area = cultivationarea.NewGreenhouseArea(
				e.ID,
				e.Name,
				e.Dim,
				e.Geometry,
			)
			log.Printf("Created GreenhouseArea from farm greenhouse: id=%s, name=%s", e.ID, e.Name)

		default:
			// Не интересуют другие типы событий
			return nil
		}

		// Сохраняем операционное место
		if err := growingProvider.CultivationAreas().Save(ctx, area); err != nil {
			return fmt.Errorf("failed to save cultivation area: %w", err)
		}

		uow.RegisterAggregate(area)
		return nil
	})

}
