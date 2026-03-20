package cultivationarea

import "samurenkoroma/services/internal/core/domain/event"

type FieldConfiguredAsMonoculture struct {
	event.BaseEvent
	FieldID    string
	SeasonID   string
	CropPlanID string
	Area       float64
}

func (e FieldConfiguredAsMonoculture) EventName() string {
	return "growing.field.configured_as_monoculture"
}

type FieldConfiguredAsPolyculture struct {
	event.BaseEvent
	FieldID  string
	SeasonID string
}

func (e FieldConfiguredAsPolyculture) EventName() string {
	return "growing.field.configured_as_polyculture"
}

type BlockConfigured struct {
	event.BaseEvent
	BlockID    string
	SeasonID   string
	CropPlanID string
	Area       float64
}

func (e BlockConfigured) EventName() string {
	return "growing.block.configured"
}

type BedConfigured struct {
	event.BaseEvent
	BedID      string
	SeasonID   string
	CropPlanID string
	Area       float64
}

func (e BedConfigured) EventName() string {
	return "growing.bed.configured"
}

type GreenhouseConfigured struct {
	event.BaseEvent
	GreenhouseID string
	SeasonID     string
}

func (e GreenhouseConfigured) EventName() string {
	return "growing.greenhouse.configured"
}
