package cultivationarea

import (
	"fmt"
	"samurenkoroma/services/internal/core/spatial"
)

type CreateAreaConfig struct {
	Type      AreaType
	FarmRefID string
	ParentID  string // для Block и Bed
	Name      string
	Geometry  spatial.GeoJSON
}

func CreateArea(config CreateAreaConfig) (CultivationArea, error) {
	switch config.Type {
	case AreaTypeField:
		return NewFieldArea(config.FarmRefID, config.Name, config.Geometry), nil

	case AreaTypeBlock:
		if config.ParentID == "" {
			return nil, ErrBlockRequiresParent
		}
		return NewBlock(config.ParentID, config.Name, config.Geometry), nil

	case AreaTypeBed:
		if config.ParentID == "" {
			return nil, ErrBedRequiresParent
		}
		return NewBed(config.ParentID, config.Name, config.Geometry), nil

	case AreaTypeGreenhouse:
		return NewGreenhouseArea(config.FarmRefID, config.Name, config.Geometry), nil

	default:
		return nil, fmt.Errorf("unknown area type: %s", config.Type)
	}
}
