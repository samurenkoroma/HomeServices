package cropplan

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"time"
)

type PlanID string
type PlanStatus string

const (
	PlanStatusDraft      PlanStatus = "draft"
	PlanStatusPublished  PlanStatus = "published"
	PlanStatusDeprecated PlanStatus = "deprecated"
	PlanStatusArchived   PlanStatus = "archived"
)

// CropPlan - агротехнический план выращивания культуры
type CropPlan struct {
	aggregate.BaseAggregate

	ID          PlanID
	CropTypeID  string  // ID типа культуры (пшеница, кукуруза)
	VarietyID   *string // ID сорта (опционально)
	Name        string  // Название плана (например "Интенсивная пшеница")
	Description string

	// Агротехнические параметры
	Duration int // Общая длительность в днях
	Version  int // Версия плана
	Status   PlanStatus

	// Этапы роста
	Stages []GrowthStage

	// Правила севооборота
	RotationRules []RotationRule

	// Требования
	Environment EnvironmentalRequirements
	Nutrients   NutrientRequirements

	// Метаданные
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt *time.Time
}

func NewCropPlan(
	cropTypeID string,
	name string,
	duration int,
	createdBy string,
) (*CropPlan, error) {
	if duration <= 0 {
		return nil, ErrInvalidDuration
	}
	if name == "" {
		return nil, ErrInvalidName
	}

	plan := &CropPlan{
		ID:            PlanID(generateID()),
		CropTypeID:    cropTypeID,
		Name:          name,
		Duration:      duration,
		Version:       1,
		Status:        PlanStatusDraft,
		Stages:        []GrowthStage{},
		RotationRules: []RotationRule{},
		CreatedBy:     createdBy,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	plan.AddEvent(CropPlanCreated{
		PlanID:     string(plan.ID),
		CropTypeID: cropTypeID,
		Name:       name,
	})

	return plan, nil
}

// AddStage добавляет этап роста
func (p *CropPlan) AddStage(stage GrowthStage) error {
	if p.Status == PlanStatusPublished {
		return ErrCannotModifyPublished
	}

	// Проверяем уникальность порядка
	for _, s := range p.Stages {
		if s.Order == stage.Order {
			return ErrStageOrderDuplicate
		}
	}

	p.Stages = append(p.Stages, stage)
	p.UpdatedAt = time.Now()

	p.AddEvent(GrowthStageAdded{
		PlanID: string(p.ID),
		Order:  stage.Order,
		Name:   stage.Name,
		Stage:  stage,
	})

	return nil
}

// UpdateStage обновляет этап
func (p *CropPlan) UpdateStage(order int, stage GrowthStage) error {
	if p.Status == PlanStatusPublished {
		return ErrCannotModifyPublished
	}

	for i, s := range p.Stages {
		if s.Order == order {
			p.Stages[i] = stage
			p.UpdatedAt = time.Now()
			return nil
		}
	}

	return ErrStageNotFound
}

// RemoveStage удаляет этап
func (p *CropPlan) RemoveStage(order int) error {
	if p.Status == PlanStatusPublished {
		return ErrCannotModifyPublished
	}

	for i, s := range p.Stages {
		if s.Order == order {
			p.Stages = append(p.Stages[:i], p.Stages[i+1:]...)
			p.UpdatedAt = time.Now()
			return nil
		}
	}

	return ErrStageNotFound
}

// AddRotationRule добавляет правило севооборота
func (p *CropPlan) AddRotationRule(rule RotationRule) error {
	if p.Status == PlanStatusPublished {
		return ErrCannotModifyPublished
	}

	// Проверяем дублирование
	for _, r := range p.RotationRules {
		if r.PredecessorCropTypeID == rule.PredecessorCropTypeID {
			return ErrRotationRuleDuplicate
		}
	}

	p.RotationRules = append(p.RotationRules, rule)
	p.UpdatedAt = time.Now()

	p.AddEvent(RotationRuleAdded{
		PlanID:                string(p.ID),
		PredecessorCropTypeID: rule.PredecessorCropTypeID,
		MinYears:              rule.MinYears,
	})

	return nil
}

// ValidateDuration проверяет соответствие длительности сумме этапов
func (p *CropPlan) ValidateDuration() error {
	total := 0
	for _, s := range p.Stages {
		total += s.Duration
	}

	if total != p.Duration {
		return ErrStageDurationMismatch
	}

	return nil
}

// Publish публикует план
func (p *CropPlan) Publish() error {
	if p.Status == PlanStatusPublished {
		return ErrAlreadyPublished
	}

	if len(p.Stages) == 0 {
		return ErrNoStages
	}

	if err := p.ValidateDuration(); err != nil {
		return err
	}

	now := time.Now()
	p.Status = PlanStatusPublished
	p.PublishedAt = &now
	p.UpdatedAt = now

	p.AddEvent(CropPlanPublished{
		PlanID:  string(p.ID),
		Version: p.Version,
		Name:    p.Name,
	})

	return nil
}

// Deprecate деактивирует план
func (p *CropPlan) Deprecate(reason string) error {
	if p.Status != PlanStatusPublished {
		return ErrOnlyPublishedCanBeDeprecated
	}

	p.Status = PlanStatusDeprecated
	p.UpdatedAt = time.Now()

	p.AddEvent(CropPlanDeprecated{
		PlanID: string(p.ID),
		Reason: reason,
	})

	return nil
}

// CreateNewVersion создает новую версию плана
func (p *CropPlan) CreateNewVersion() (*CropPlan, error) {
	if p.Status != PlanStatusPublished && p.Status != PlanStatusDeprecated {
		return nil, ErrOnlyPublishedCanBeVersioned
	}

	newPlan := &CropPlan{
		ID:            PlanID(generateID()),
		CropTypeID:    p.CropTypeID,
		VarietyID:     p.VarietyID,
		Name:          p.Name,
		Description:   p.Description,
		Duration:      p.Duration,
		Version:       p.Version + 1,
		Status:        PlanStatusDraft,
		Stages:        p.Stages, // копируем этапы
		RotationRules: p.RotationRules,
		Environment:   p.Environment,
		Nutrients:     p.Nutrients,
		CreatedBy:     p.CreatedBy,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	newPlan.AddEvent(CropPlanVersionCreated{
		OriginalPlanID: string(p.ID),
		NewPlanID:      string(newPlan.ID),
		Version:        newPlan.Version,
	})

	return newPlan, nil
}

// SetVariety устанавливает сорт
func (p *CropPlan) SetVariety(varietyID string) error {
	if p.Status == PlanStatusPublished {
		return ErrCannotModifyPublished
	}

	p.VarietyID = &varietyID
	p.UpdatedAt = time.Now()

	return nil
}

// SetRequirements устанавливает требования
func (p *CropPlan) SetRequirements(env EnvironmentalRequirements, nut NutrientRequirements) {
	p.Environment = env
	p.Nutrients = nut
	p.UpdatedAt = time.Now()
}

// Геттеры
func (p *CropPlan) GetID() PlanID                             { return p.ID }
func (p *CropPlan) GetCropTypeID() string                     { return p.CropTypeID }
func (p *CropPlan) GetVarietyID() *string                     { return p.VarietyID }
func (p *CropPlan) GetName() string                           { return p.Name }
func (p *CropPlan) GetDuration() int                          { return p.Duration }
func (p *CropPlan) GetVersion() int                           { return p.Version }
func (p *CropPlan) GetStatus() PlanStatus                     { return p.Status }
func (p *CropPlan) GetStages() []GrowthStage                  { return p.Stages }
func (p *CropPlan) GetRotationRules() []RotationRule          { return p.RotationRules }
func (p *CropPlan) GetEnvironment() EnvironmentalRequirements { return p.Environment }
func (p *CropPlan) GetNutrients() NutrientRequirements        { return p.Nutrients }

// Rehydrate восстанавливает план из БД (для репозитория)
func (p *CropPlan) Rehydrate(version int, status PlanStatus) {
	p.Version = version
	p.Status = status
}
