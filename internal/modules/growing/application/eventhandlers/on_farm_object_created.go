package eventhandlers

import (
	"context"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

// OnFarmObjectCreated - реакция на создание объекта в farm модуле
func OnFarmObjectCreated(ctx context.Context, farmRefID string, objType string, name string, geom spatial.GeoJSON) error {
	uow, ok := repository.FromContext(ctx)
	if !ok {
		return nil
	}

	growingProvider := uow.(*postgres.GrowingProvider)

	var area cultivationarea.CultivationArea
	var err error

	switch objType {
	case "field":
		area = cultivationarea.NewFieldArea(farmRefID, name, geom)
	case "greenhouse":
		area = cultivationarea.NewGreenhouseArea(farmRefID, name, geom)
	default:
		return nil // Не интересуют другие типы
	}

	if err := growingProvider.CultivationAreas().Save(ctx, area); err != nil {
		return err
	}

	uow.RegisterAggregate(area)
	return nil
}
