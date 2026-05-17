package eventhandlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type SchemaElement struct {
	Id      string  `json:"id"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Icon    string  `json:"icon"`
	TypeObj string  `json:"type"`
	Color   string  `json:"color"`
	Label   string  `json:"label"`
	Width   float64 `json:"width"`
	Length  float64 `json:"height"`
}

func OnFarmObjectSchemaUpdated(ctx context.Context, event event.DomainEvent) error {
	factory, ok := repository.FactoryFromContext(ctx)
	if !ok {
		return nil
	}
	uow, err := factory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = uow.Execute(ctx, postgres.NewPostgresGrowingProvider, func(provider repository.RepositoryProvider) (any, error) {
		growingProvider, ok := provider.(*postgres.PostgresGrowingProvider)
		if !ok {
			return nil, fmt.Errorf("invalid provider type")
		}
		var area cultivationarea.CultivationArea

		switch e := event.(type) {
		case physicalobject.PhysicalObjectSchemaUpdated:
			var elements []SchemaElement
			// Создано поле → создаём FieldArea
			if err := json.Unmarshal(e.Schema, &elements); err != nil {
				return nil, err
			}

			for _, bed := range elements {
				if bed.TypeObj != string(cultivationarea.AreaTypeBed) {
					continue
				}
				area, err = growingProvider.CultivationAreas().FindById(ctx, bed.Id)
				if err != nil {
					if !errors.As(err, &cultivationarea.ErrAreaNotFound) {
						return nil, err
					}
				}
				if area == nil {
					area = cultivationarea.NewBed(
						bed.Id,
						e.ObjectID,
						fmt.Sprintf("%s - %s", e.Name, bed.Label),
						bed.Width*bed.Length/10000,
					)
				} else {
					area.SetArea(bed.Width * bed.Length / 10000)
					area.SetName(fmt.Sprintf("%s - %s", e.Name, bed.Label))
				}
				a, ok := area.(*cultivationarea.Bed)
				if !ok {
					return nil, fmt.Errorf("invalid area type")
				}
				a.SetAttributes(bed.Width, bed.Length, bed.X, bed.Y)
				// Сохраняем операционное место
				if err := growingProvider.CultivationAreas().Save(ctx, a); err != nil {
					return nil, err
				}
			}

		default:
			// Не интересуют другие типы событий
			return nil, nil
		}

		uow.RegisterAggregate(area)
		return nil, nil
	})
	return err
}
