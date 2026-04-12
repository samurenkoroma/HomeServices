package cultivationarea

import (
	"fmt"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
)

// CreateAreaConfig — конфигурация для создания места
type CreateAreaConfig struct {
	Type      AreaType
	FarmRefID string
	ParentID  string // для Block и Bed
	Name      string
	Geometry  spatial.GeoJSON
}

// CreateArea создаёт место выращивания нужного типа
func CreateArea(config CreateAreaConfig) (CultivationArea, error) {
	switch config.Type {
	case AreaTypeField:
		return NewFieldArea(config.FarmRefID, config.Name, config.Geometry, 0), nil

	case AreaTypeBlock:
		if config.ParentID == "" {
			return nil, ErrBlockRequiresParent
		}
		return NewBlock(config.ParentID, config.Name, config.Geometry), nil

	case AreaTypeBed:
		if config.ParentID == "" {
			return nil, ErrBedRequiresParent
		}
		return NewBed(types.NewUUID(), config.ParentID, config.Name, config.Geometry, 0), nil

	case AreaTypeGreenhouse:
		return NewGreenhouseArea(config.FarmRefID, config.Name, types.Dimension{}, config.Geometry), nil

	default:
		return nil, fmt.Errorf("%w: %s", ErrUnknownAreaType, config.Type)
	}
}
