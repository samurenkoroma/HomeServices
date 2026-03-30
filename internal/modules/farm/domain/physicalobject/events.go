package physicalobject

import (
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
)

type GreenhouseCreated struct {
	event.BaseEvent
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Geometry spatial.GeoJSON `json:"geometry"`
	Dim      types.Dimension `json:"dim"`
	Width    float64         `json:"width"`
}

func (e GreenhouseCreated) EventName() string {
	return "farm.greenhouse.created"
}

type FieldCreated struct {
	event.BaseEvent
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Geometry spatial.GeoJSON `json:"geometry"`
	Area     float64         `json:"area"`
}

func (e FieldCreated) EventName() string {
	return "farm.field.created"
}

// PhysicalObjectActivated — событие активации объекта
type PhysicalObjectActivated struct {
	event.BaseEvent
	ObjectID string `json:"object_id"`
	Type     string `json:"type"`
}

func (e PhysicalObjectActivated) EventName() string {
	return "farm.object.activated"
}

// PhysicalObjectDeactivated — событие деактивации объекта
type PhysicalObjectDeactivated struct {
	event.BaseEvent
	ObjectID string `json:"object_id"`
	Type     string `json:"type"`
}

func (e PhysicalObjectDeactivated) EventName() string {
	return "farm.object.deactivated"
}
