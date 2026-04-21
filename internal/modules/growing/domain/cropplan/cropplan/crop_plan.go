package cropplan

import (
	"errors"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"time"
)

// Status статус плана
type Status string

const (
	StatusDraft     Status = "draft"
	StatusActive    Status = "active"
	StatusCompleted Status = "completed"
	StatusCancelled Status = "cancelled"
)

type Area interface {
	GetId() string
}
type Variety interface {
	GetId() string
	GetName() string
	GetSpeciesName() string
}
type Season interface {
	GetId() string
	IsFinished() bool
	GetStartDate() time.Time
	GetEndDate() time.Time
}

type CropPlan struct {
	aggregate.BaseAggregate

	// Идентификация
	id      string
	name    string
	area    Area
	variety Variety
	season  Season

	plantingDate time.Time

	harvestKg float64

	metadata map[string]interface{}

	// Статус и этапы
	status Status
	stages []Stage

	// Даты
	createdAt   time.Time
	updatedAt   time.Time
	startedAt   *time.Time
	completedAt *time.Time

	// Ответственный агроном
	assignedTo string
}

// NewCropPlan создает новый план
func NewCropPlan(
	id,
	name string,
	plantingDate time.Time,
	area Area,
	season Season,
	variety Variety,
	assignedTo string,
) (*CropPlan, error) {

	if id == "" {
		return nil, ErrInvalidPlanID
	}
	if name == "" {
		return nil, errors.New("plan name is required")
	}
	if variety.GetId() == "" {
		return nil, ErrVarietyRequired
	}
	if area.GetId() == "" {
		return nil, ErrAreaRequired
	}
	if season.GetId() == "" {
		return nil, ErrSeasonRequired
	}
	if types.UUIDIsValid(assignedTo) {
		return nil, errors.New("assigned plan is required")
	}

	now := time.Now()

	plan := &CropPlan{
		id:           id,
		area:         area,
		name:         name,
		variety:      variety,
		season:       season,
		plantingDate: plantingDate,
		status:       StatusDraft,
		stages:       []Stage{},
		createdAt:    now,
		updatedAt:    now,
		assignedTo:   assignedTo,
		metadata:     make(map[string]interface{}),
	}

	plan.AddEvent(CropPlanCreatedEvent{
		PlanID:       id,
		AreaId:       area.GetId(),
		Name:         name,
		VarietyID:    variety.GetId(),
		VarietyName:  variety.GetName(),
		SpeciesName:  variety.GetSpeciesName(),
		PlantingDate: plantingDate,
	})

	return plan, nil
}

// ========== GETTERS ==========

func (p *CropPlan) ID() string              { return p.id }
func (p *CropPlan) Area() Area              { return p.area }
func (p *CropPlan) Name() string            { return p.name }
func (p *CropPlan) Variety() Variety        { return p.variety }
func (p *CropPlan) Season() Season          { return p.season }
func (p *CropPlan) PlantingDate() time.Time { return p.plantingDate }
func (p *CropPlan) Status() Status          { return p.status }
func (p *CropPlan) Stages() []Stage         { return append([]Stage(nil), p.stages...) }

// func (p *CropPlan) SeedsPlanted() int       { return p.seedsPlanted }
// func (p *CropPlan) ExpectedYield() float64  { return p.expectedYield }
func (p *CropPlan) HarvestKg() float64               { return p.harvestKg }
func (p *CropPlan) CreatedAt() time.Time             { return p.createdAt }
func (p *CropPlan) UpdatedAt() time.Time             { return p.updatedAt }
func (p *CropPlan) StartedAt() *time.Time            { return p.startedAt }
func (p *CropPlan) CompletedAt() *time.Time          { return p.completedAt }
func (p *CropPlan) AssignedTo() string               { return p.assignedTo }
func (p *CropPlan) Metadata() map[string]interface{} { return p.metadata }

// ========== УПРАВЛЕНИЕ ЭТАПАМИ ==========

// AddStage добавляет этап в план
func (p *CropPlan) AddStage(stage Stage) error {
	if p.status != StatusDraft {
		return ErrPlanNotDraft
	}

	// Проверка уникальности порядка
	for _, s := range p.stages {
		if s.Order == stage.Order {
			return ErrStageOrderDuplicate
		}
	}

	stage.PlanID = p.id
	p.stages = append(p.stages, stage)
	p.updatedAt = time.Now()

	p.AddEvent(StageAddedEvent{
		PlanID:    p.id,
		StageID:   stage.ID,
		StageName: stage.Name,
		BBCHStart: stage.BBCHStart,
		BBCHEnd:   stage.BBCHEnd,
		Order:     stage.Order,
	})

	return nil
}

// AddStagesFromTemplates добавляет этапы из шаблонов
func (p *CropPlan) AddStagesFromTemplates(templates []catalog.StageTemplate) error {
	for i, template := range templates {
		stage, err := NewStageFromTemplate(
			types.NewUUID(),
			p.id,
			template,
			i+1,
		)
		if err != nil {
			return err
		}
		stage.Description = template.Description
		if err := p.AddStage(*stage); err != nil {
			return err
		}
	}
	return nil
}

// GetStage возвращает этап по ID
func (p *CropPlan) GetStage(stageID string) (*Stage, error) {
	for i := range p.stages {
		if p.stages[i].ID == stageID {
			return &p.stages[i], nil
		}
	}
	return nil, ErrStageNotFound
}

// GetNextStage возвращает следующий незавершенный этап
func (p *CropPlan) GetNextStage() *Stage {
	for i := range p.stages {
		if !p.stages[i].IsFinished() {
			return &p.stages[i]
		}
	}
	return nil
}

// GetPendingStages возвращает все ожидающие этапы
func (p *CropPlan) GetPendingStages() []Stage {
	var pending []Stage
	for _, s := range p.stages {
		if s.Status == StageStatusPending {
			pending = append(pending, s)
		}
	}
	return pending
}

// StartStage начинает этап
func (p *CropPlan) StartStage(stageID string, currentBBCH int) error {
	if p.status != StatusActive {
		return errors.New("can only start stages in active plan")
	}

	stage, err := p.GetStage(stageID)
	if err != nil {
		return err
	}

	if err := stage.Start(currentBBCH); err != nil {
		return err
	}

	p.updatedAt = time.Now()

	p.AddEvent(StageStartedEvent{
		PlanID:    p.id,
		StageID:   stage.ID,
		StageName: stage.Name,
		BBCHCode:  currentBBCH,
		StartedAt: *stage.StartDate,
	})

	return nil
}

// CompleteStage завершает этап
func (p *CropPlan) CompleteStage(stageID string) error {
	if p.status != StatusActive {
		return errors.New("can only complete stages in active plan")
	}

	stage, err := p.GetStage(stageID)
	if err != nil {
		return err
	}

	if err := stage.Complete(); err != nil {
		return err
	}

	p.updatedAt = time.Now()

	p.AddEvent(StageCompletedEvent{
		PlanID:      p.id,
		StageID:     stage.ID,
		StageName:   stage.Name,
		CompletedAt: *stage.EndDate,
	})

	// Проверяем, все ли этапы завершены
	if p.allStagesFinished() {
		p.complete()
	}

	return nil
}

// SkipStage пропускает этап
func (p *CropPlan) SkipStage(stageID string) error {
	if p.status != StatusActive && p.status != StatusDraft {
		return errors.New("cannot skip stages in this status")
	}

	stage, err := p.GetStage(stageID)
	if err != nil {
		return err
	}

	if err := stage.Skip(); err != nil {
		return err
	}

	p.updatedAt = time.Now()

	p.AddEvent(StageSkippedEvent{
		PlanID:    p.id,
		StageID:   stage.ID,
		StageName: stage.Name,
		SkippedAt: time.Now(),
	})

	return nil
}

// GetApplicableStagesForBBCH возвращает этапы, которые можно выполнять при данном BBCH
func (p *CropPlan) GetApplicableStagesForBBCH(bbchCode int) []Stage {
	var applicable []Stage
	for _, stage := range p.stages {
		if stage.Status == StageStatusPending && stage.IsApplicableForBBCH(bbchCode) {
			applicable = append(applicable, stage)
		}
	}
	return applicable
}

// allStagesFinished проверяет, все ли этапы завершены
func (p *CropPlan) allStagesFinished() bool {
	for _, s := range p.stages {
		if !s.IsFinished() {
			return false
		}
	}
	return true
}

// ========== ЖИЗНЕННЫЙ ЦИКЛ ПЛАНА ==========

// CanActivate проверяет, можно ли активировать план
func (p *CropPlan) CanActivate() error {
	if p.status != StatusDraft {
		return ErrPlanNotDraft
	}
	if len(p.stages) == 0 {
		return ErrNoStages
	}
	if p.season.IsFinished() {
		return errors.New("cannot activate plan after season end")
	}
	return nil
}

// Activate активирует план
func (p *CropPlan) Activate() error {
	if err := p.CanActivate(); err != nil {
		return err
	}

	now := time.Now()
	p.status = StatusActive
	p.startedAt = &now
	p.updatedAt = now

	p.AddEvent(CropPlanActivatedEvent{
		PlanID:      p.id,
		ActivatedAt: now,
	})

	return nil
}

// CanComplete проверяет, можно ли завершить план
func (p *CropPlan) CanComplete() error {
	if p.status != StatusActive {
		return ErrPlanNotActive
	}
	if !p.allStagesFinished() {
		return errors.New("cannot complete plan with unfinished stages")
	}
	return nil
}

// complete внутренний метод завершения плана
func (p *CropPlan) complete() {
	now := time.Now()
	p.status = StatusCompleted
	p.completedAt = &now
	p.updatedAt = now

	p.AddEvent(CropPlanCompletedEvent{
		PlanID:      p.id,
		CompletedAt: now,
		HarvestKg:   p.harvestKg,
	})
}

// Complete завершает план со сбором урожая
func (p *CropPlan) Complete(harvestKg float64) error {
	if err := p.CanComplete(); err != nil {
		return err
	}

	p.metadata["harvestKg"] = harvestKg
	p.complete()
	return nil
}

// Cancel отменяет план
func (p *CropPlan) Cancel(reason string) error {
	if p.status == StatusCompleted || p.status == StatusCancelled {
		return errors.New("plan already completed or cancelled")
	}

	p.status = StatusCancelled
	p.updatedAt = time.Now()

	p.AddEvent(CropPlanCancelledEvent{
		PlanID:      p.id,
		Reason:      reason,
		CancelledAt: time.Now(),
	})

	return nil
}

// ========== АГРОНОМИЧЕСКИЕ РАСЧЕТЫ ==========

// SetSeedsPlanted устанавливает количество посаженных семян
func (p *CropPlan) SetSeedsPlanted(seeds int) {
	p.metadata["seedsPlanted"] = seeds
	p.updatedAt = time.Now()
}

// SetExpectedYield устанавливает ожидаемую урожайность
func (p *CropPlan) SetExpectedYield(yield float64) {
	p.metadata["expectedYield"] = yield
	p.updatedAt = time.Now()
}

// CalculateExpectedYield рассчитывает ожидаемую урожайность на основе сорта и площади
func (p *CropPlan) CalculateExpectedYield(areaM2 float64, yieldPotential float64) float64 {
	return areaM2 * yieldPotential
}
