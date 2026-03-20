package physicalobject

import (
	"samurenkoroma/services/internal/core/domain/event"
)

type GreenhouseCreated struct {
	event.BaseEvent
	ID   string
	Name string
	Area float64
}

func (e GreenhouseCreated) EventName() string {
	return "farm.greenhouse.created"
}

func NewGreenhouseCreated(Id, name string, area float64) GreenhouseCreated {
	return GreenhouseCreated{
		ID:   Id,
		Name: name,
		Area: area,
	}
}

type FieldCreated struct {
	event.BaseEvent

	ID        string
	Name      string
	TotalArea float64
}

func (e FieldCreated) EventName() string {
	return "farm.field.created"
}

func NewFieldCreated(id string, name string, area float64) FieldCreated {
	return FieldCreated{
		BaseEvent: event.NewBaseEvent(),
		ID:        id,
		Name:      name,
		TotalArea: area,
	}
}
