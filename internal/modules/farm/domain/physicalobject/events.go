package physicalobject

import (
	"encoding/json"
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
)

var (
	FarmGreenhouseCreatedEvent   = "farm.greenhouse.created"
	FarmObjectSchemaUpdatedEvent = "farm.object.schema.updated"
)

type GreenhouseCreated struct {
	event.BaseEvent
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Geometry spatial.GeoJSON `json:"geometry"`
	Dim      types.Dimension `json:"dim"`
}

func (e GreenhouseCreated) EventName() string {
	return FarmGreenhouseCreatedEvent
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

type PlotCreated struct {
	event.BaseEvent
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Geometry spatial.GeoJSON `json:"geometry"`
	Area     float64         `json:"area"`
}

func (e PlotCreated) EventName() string {
	return "farm.plot.created"
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

type PhysicalObjectSchemaUpdated struct {
	event.BaseEvent
	ObjectID string          `json:"object_id"`
	Name     string          `json:"name"`
	Schema   json.RawMessage `json:"schema"`
	Geometry spatial.GeoJSON `json:"geometry"`
}

func (e PhysicalObjectSchemaUpdated) EventName() string {
	return FarmObjectSchemaUpdatedEvent
}
