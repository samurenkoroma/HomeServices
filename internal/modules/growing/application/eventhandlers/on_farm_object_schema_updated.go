package eventhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"

	"github.com/google/uuid"
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

	return uow.Execute(ctx, postgres.NewPostgresGrowingProvider, func(provider repository.RepositoryProvider) error {
		growingProvider, ok := provider.(*postgres.PostgresGrowingProvider)
		if !ok {
			return fmt.Errorf("invalid provider type")
		}
		var area cultivationarea.CultivationArea

		switch e := event.(type) {
		case physicalobject.PhysicalObjectSchemaUpdated:
			var elements []SchemaElement
			// Создано поле → создаём FieldArea
			err := json.Unmarshal(e.Schema, &elements)
			if err != nil {
				return err
			}
			for _, bed := range elements {
				if bed.TypeObj != string(cultivationarea.AreaTypeBed) {
					continue
				}
				bUUID, err := uuid.Parse(bed.Id)
				if err != nil {
					return err
				}

				area = cultivationarea.NewBed(
					bUUID.String(),
					e.ObjectID,
					fmt.Sprintf("%s -%s", e.Name, bed.Label),
					e.Geometry,
					bed.Width*bed.Length,
				)
				a, ok := area.(*cultivationarea.Bed)
				if !ok {
					return fmt.Errorf("invalid area type")
				}
				a.SetAttributes(bed.Width, bed.Length, bed.X, bed.Y)
				// Сохраняем операционное место
				if err := growingProvider.CultivationAreas().Save(ctx, a); err != nil {
					return fmt.Errorf("failed to save cultivation area: %w", err)
				}
			}

		default:
			// Не интересуют другие типы событий
			return nil
		}

		uow.RegisterAggregate(area)
		return nil
	})

}
