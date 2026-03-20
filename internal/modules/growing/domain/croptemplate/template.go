package croptemplate

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"time"
)

type TemplateStatus string

const (
	TemplateStatusDraft     TemplateStatus = "draft"
	TemplateStatusPublished TemplateStatus = "published"
	TemplateStatusArchived  TemplateStatus = "archived"
)

// CropTemplate - шаблон выращивания культуры
type CropTemplate struct {
	aggregate.BaseAggregate

	ID         string
	CropPlanID string // ID из модуля crop
	Name       string
	Version    int
	Status     TemplateStatus

	Stages       []GrowthStage
	Requirements Requirements

	CreatedAt   time.Time
	PublishedAt *time.Time
	UpdatedAt   time.Time
}

// GrowthStage - этап роста
type GrowthStage struct {
	Order       int
	Name        string
	Duration    int // дней
	MinTemp     float64
	MaxTemp     float64
	MinHumidity float64
	MaxHumidity float64
	WaterPerDay float64 // литров на м² в день
}

// Requirements - требования к условиям
type Requirements struct {
	MinPH      float64
	MaxPH      float64
	Nitrogen   float64 // кг/га
	Phosphorus float64
	Potassium  float64
	SoilType   []string // подходящие типы почв
}

func NewCropTemplate(cropPlanID string, name string) *CropTemplate {
	return &CropTemplate{
		ID:         generateID(),
		CropPlanID: cropPlanID,
		Name:       name,
		Version:    1,
		Status:     TemplateStatusDraft,
		Stages:     []GrowthStage{},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (t *CropTemplate) AddStage(stage GrowthStage) error {
	if t.Status == TemplateStatusPublished {
		return ErrCannotModifyPublished
	}

	t.Stages = append(t.Stages, stage)
	t.UpdatedAt = time.Now()

	return nil
}

func (t *CropTemplate) Publish() error {
	if t.Status == TemplateStatusPublished {
		return ErrAlreadyPublished
	}

	if len(t.Stages) == 0 {
		return ErrNoStages
	}

	now := time.Now()
	t.Status = TemplateStatusPublished
	t.PublishedAt = &now
	t.UpdatedAt = now

	t.AddEvent(CropTemplatePublished{
		TemplateID: t.ID,
		Version:    t.Version,
	})

	return nil
}
