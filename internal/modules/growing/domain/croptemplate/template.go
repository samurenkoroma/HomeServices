package croptemplate

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"time"
)

// TemplateID — идентификатор шаблона
type TemplateID string

// TemplateStatus — статус шаблона
type TemplateStatus string

const (
	TemplateStatusDraft     TemplateStatus = "draft"
	TemplateStatusPublished TemplateStatus = "published"
	TemplateStatusArchived  TemplateStatus = "archived"
)

// CropTemplate — шаблон выращивания культуры (из CropPlan)
type CropTemplate struct {
	aggregate.Entity[TemplateID]

	id         TemplateID
	cropPlanID string // ID из модуля crop
	name       string
	version    int
	status     TemplateStatus

	// Этапы роста (из CropPlan)
	stages []GrowthStage

	// Требования к условиям
	requirements Requirements

	createdAt   time.Time
	publishedAt *time.Time
	updatedAt   time.Time
}
type Attributes map[string]interface{}

// GrowthStage — этап роста (копия из crop, но с добавленными полями для growing)
type GrowthStage struct {
	Order           int     `json:"order"`
	Name            string  `json:"name"`
	Duration        int     `json:"duration"` // дней
	MinTemp         float64 `json:"min_temp"`
	MaxTemp         float64 `json:"max_temp"`
	OptimalTemp     float64 `json:"optimal_temp"`
	Recommendations Attributes
	WaterPerDay     float64 `json:"water_per_day"` // л/м² в день
	Description     string  `json:"description"`
}

// Requirements — требования к условиям
type Requirements struct {
	MinPH      float64  `json:"min_ph"`
	MaxPH      float64  `json:"max_ph"`
	Nitrogen   float64  `json:"nitrogen"`   // кг/га
	Phosphorus float64  `json:"phosphorus"` // кг/га
	Potassium  float64  `json:"potassium"`  // кг/га
	SoilTypes  []string `json:"soil_types"`
}

// NewCropTemplate создаёт новый шаблон из CropPlan
func NewCropTemplate(cropPlanID, name string, version int) *CropTemplate {
	return &CropTemplate{
		Entity:     aggregate.NewEntity(TemplateID(types.NewUUID())),
		cropPlanID: cropPlanID,
		name:       name,
		version:    version,
		status:     TemplateStatusDraft,
		stages:     []GrowthStage{},
	}
}

// AddStage добавляет этап
func (t *CropTemplate) AddStage(stage GrowthStage) error {
	if t.status == TemplateStatusPublished {
		return ErrCannotModifyPublished
	}

	// Проверяем уникальность порядка
	for _, s := range t.stages {
		if s.Order == stage.Order {
			return ErrStageOrderDuplicate
		}
	}

	t.stages = append(t.stages, stage)
	t.updatedAt = time.Now()
	return nil
}

// SetRequirements устанавливает требования
func (t *CropTemplate) SetRequirements(req Requirements) {
	t.requirements = req
	t.updatedAt = time.Now()
}

// Publish публикует шаблон
func (t *CropTemplate) Publish() error {
	if t.status == TemplateStatusPublished {
		return ErrAlreadyPublished
	}

	if len(t.stages) == 0 {
		return ErrNoStages
	}

	now := time.Now()
	t.status = TemplateStatusPublished
	t.publishedAt = &now
	t.updatedAt = now

	t.AddEvent(CropTemplatePublished{
		TemplateID: string(t.id),
		CropPlanID: t.cropPlanID,
		Version:    t.version,
	})

	return nil
}

// Archive архивирует шаблон
func (t *CropTemplate) Archive() error {
	if t.status == TemplateStatusArchived {
		return ErrAlreadyArchived
	}

	t.status = TemplateStatusArchived
	t.updatedAt = time.Now()

	t.AddEvent(CropTemplateArchived{
		TemplateID: string(t.id),
	})

	return nil
}

// GetStage возвращает этап по порядковому номеру
func (t *CropTemplate) GetStage(order int) (*GrowthStage, error) {
	for _, s := range t.stages {
		if s.Order == order {
			return &s, nil
		}
	}
	return nil, ErrStageNotFound
}

// ValidateDuration проверяет соответствие длительности
func (t *CropTemplate) ValidateDuration() int {
	total := 0
	for _, s := range t.stages {
		total += s.Duration
	}
	return total
}

// Геттеры
func (t *CropTemplate) GetID() TemplateID             { return t.id }
func (t *CropTemplate) GetCropPlanID() string         { return t.cropPlanID }
func (t *CropTemplate) GetName() string               { return t.name }
func (t *CropTemplate) GetVersion() int               { return t.version }
func (t *CropTemplate) GetStatus() TemplateStatus     { return t.status }
func (t *CropTemplate) GetStages() []GrowthStage      { return t.stages }
func (t *CropTemplate) GetRequirements() Requirements { return t.requirements }
func (t *CropTemplate) GetCreatedAt() time.Time       { return t.createdAt }
func (t *CropTemplate) GetPublishedAt() *time.Time    { return t.publishedAt }
func (t *CropTemplate) GetUpdatedAt() time.Time       { return t.updatedAt }

// Rehydrate восстанавливает шаблон из БД
func (t *CropTemplate) Rehydrate(
	id TemplateID,
	cropPlanID string,
	name string,
	version int,
	status TemplateStatus,
	stages []GrowthStage,
	requirements Requirements,
	createdAt time.Time,
	publishedAt *time.Time,
	updatedAt time.Time,
) {
	t.id = id
	t.cropPlanID = cropPlanID
	t.name = name
	t.version = version
	t.status = status
	t.stages = stages
	t.requirements = requirements
	t.createdAt = createdAt
	t.publishedAt = publishedAt
	t.updatedAt = updatedAt
}
