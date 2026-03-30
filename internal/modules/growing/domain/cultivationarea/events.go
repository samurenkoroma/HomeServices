package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/event"
)

// FieldConfiguredAsMonoculture — поле настроено как монокультура
type FieldConfiguredAsMonoculture struct {
	event.BaseEvent
	FieldID    string  `json:"field_id"`
	SeasonID   string  `json:"season_id"`
	CropPlanID string  `json:"crop_plan_id"`
	Area       float64 `json:"area"`
}

func (e FieldConfiguredAsMonoculture) EventName() string {
	return "growing.field.configured_as_monoculture"
}

// FieldConfiguredAsPolyculture — поле настроено как поликультура
type FieldConfiguredAsPolyculture struct {
	event.BaseEvent
	FieldID  string `json:"field_id"`
	SeasonID string `json:"season_id"`
}

func (e FieldConfiguredAsPolyculture) EventName() string {
	return "growing.field.configured_as_polyculture"
}

// BlockConfigured — участок настроен на сезон
type BlockConfigured struct {
	event.BaseEvent
	BlockID    string  `json:"block_id"`
	SeasonID   string  `json:"season_id"`
	CropPlanID string  `json:"crop_plan_id"`
	Area       float64 `json:"area"`
}

func (e BlockConfigured) EventName() string {
	return "growing.block.configured"
}

// BedConfigured — грядка настроена на сезон
type BedConfigured struct {
	event.BaseEvent
	BedID    string `json:"bed_id"`
	SeasonID string `json:"season_id"`
	//CropPlanID string  `json:"crop_plan_id"`
	Area float64 `json:"area"`
}

func (e BedConfigured) EventName() string {
	return "growing.bed.configured"
}

// GreenhouseConfigured — теплица настроена на сезон
type GreenhouseConfigured struct {
	event.BaseEvent
	GreenhouseID string `json:"greenhouse_id"`
	SeasonID     string `json:"season_id"`
}

func (e GreenhouseConfigured) EventName() string {
	return "growing.greenhouse.configured"
}
