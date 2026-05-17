package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
)

type AreaType string

const (
	AreaTypeField    AreaType = "field"
	AreaTypeBed      AreaType = "bed"
	AreaTypeCassette AreaType = "cassette"
)

type CultivationArea interface {
	aggregate.Aggregate
	// Базовые методы
	GetId() string
	GetFarmRefID() string
	GetType() AreaType
	GetName() string
	GetArea() float64
	SetName(name string)
	SetArea(area float64)
}
