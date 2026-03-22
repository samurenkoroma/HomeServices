package croptemplate

import (
	"samurenkoroma/services/internal/core/domain/event"
)

// CropTemplateCreated — событие создания шаблона
type CropTemplateCreated struct {
	event.BaseEvent
	TemplateID string `json:"template_id"`
	CropPlanID string `json:"crop_plan_id"`
	Name       string `json:"name"`
	Version    int    `json:"version"`
}

func (e CropTemplateCreated) EventName() string {
	return "growing.template.created"
}

// CropTemplatePublished — событие публикации шаблона
type CropTemplatePublished struct {
	event.BaseEvent
	TemplateID string `json:"template_id"`
	CropPlanID string `json:"crop_plan_id"`
	Version    int    `json:"version"`
}

func (e CropTemplatePublished) EventName() string {
	return "growing.template.published"
}

// CropTemplateArchived — событие архивации шаблона
type CropTemplateArchived struct {
	event.BaseEvent
	TemplateID string `json:"template_id"`
}

func (e CropTemplateArchived) EventName() string {
	return "growing.template.archived"
}

// StageAdded — событие добавления этапа
type StageAdded struct {
	event.BaseEvent
	TemplateID string      `json:"template_id"`
	Stage      GrowthStage `json:"stage"`
}

func (e StageAdded) EventName() string {
	return "growing.template.stage_added"
}
