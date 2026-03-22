package physicalobject

import (
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/core/spatial"
)

type GreenhouseCreated struct {
	event.BaseEvent
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Geometry spatial.GeoJSON `json:"geometry"`
}

func (e GreenhouseCreated) EventName() string {
	return "farm.greenhouse.created"
}

type FieldCreated struct {
	event.BaseEvent
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Geometry spatial.GeoJSON `json:"geometry"`
}

func (e FieldCreated) EventName() string {
	return "farm.field.created"
}
