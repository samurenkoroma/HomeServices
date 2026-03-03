package cropplan

import (
	"errors"
	"samurenkoroma/services/internal/common/domain"
	"time"
)

const (
	StatusDraft     Status = "draft"
	StatusActive    Status = "active"
	StatusCompleted Status = "completed"
	StatusCancelled Status = "cancelled"
)

type CropPlan struct {
	domain.BaseAggregate

	id          PlanID
	facilityID  GrowingAreaID
	name        string
	cropName    string
	status      Status
	stages      []Stage
	createdAt   time.Time
	updatedAt   time.Time
	startedAt   *time.Time
	completedAt *time.Time
	harvestKg   float64
}

// Конструктор
func NewCropPlan(id PlanID, facilityID GrowingAreaID, name, cropName string) *CropPlan {
	now := time.Now()
	plan := &CropPlan{
		id:         id,
		facilityID: facilityID,
		name:       name,
		cropName:   cropName,
		status:     StatusDraft,
		stages:     []Stage{},
		createdAt:  now,
		updatedAt:  now,
	}

	plan.AddEvent(CropPlanCreatedEvent{
		PlanID:   string(id),
		AreaID:   string(facilityID),
		CropName: cropName,
		Time:     now,
	})

	return plan
}

// Getters
func (p *CropPlan) ID() PlanID                { return p.id }
func (p *CropPlan) FacilityID() GrowingAreaID { return p.facilityID }
func (p *CropPlan) Name() string              { return p.name }
func (p *CropPlan) CropName() string          { return p.cropName }
func (p *CropPlan) Status() Status            { return p.status }
func (p *CropPlan) Stages() []Stage           { return append([]Stage(nil), p.stages...) }
func (p *CropPlan) CreatedAt() time.Time      { return p.createdAt }
func (p *CropPlan) StartedAt() *time.Time     { return p.startedAt }
func (p *CropPlan) CompletedAt() *time.Time   { return p.completedAt }
func (p *CropPlan) HarvestKg() float64        { return p.harvestKg }

// Управление этапами

// AddStage добавляет этап в план
func (p *CropPlan) AddStage(stage Stage) error {
	if p.status != StatusDraft {
		return errors.New("cannot add stages to non-draft plan")
	}

	// Проверяем уникальность порядка
	for _, s := range p.stages {
		if s.Order == stage.Order {
			return ErrStageOrderDuplicate
		}
	}

	stage.PlanID = string(p.id)
	p.stages = append(p.stages, stage)
	p.updatedAt = time.Now()

	p.AddEvent(StageAddedEvent{
		PlanID:  string(p.id),
		StageID: stage.ID,
		Type:    string(stage.Type),
		Order:   stage.Order,
		Time:    time.Now(),
	})

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

// StartStage начинает этап
func (p *CropPlan) StartStage(stageID string) error {
	if p.status != StatusActive {
		return errors.New("can only start stages in active plan")
	}

	stage, err := p.GetStage(stageID)
	if err != nil {
		return err
	}

	if err := stage.Start(); err != nil {
		return err
	}

	p.updatedAt = time.Now()

	p.AddEvent(StageStartedEvent{
		PlanID:    string(p.id),
		StageID:   stage.ID,
		StartedAt: *stage.StartDate,
		Time:      time.Now(),
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
		PlanID:      string(p.id),
		StageID:     stage.ID,
		CompletedAt: *stage.EndDate,
		Time:        time.Now(),
	})

	// Проверяем, все ли этапы завершены
	allCompleted := true
	for _, s := range p.stages {
		if !s.IsFinished() {
			allCompleted = false
			break
		}
	}

	if allCompleted && p.status == StatusActive {
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
		PlanID:    string(p.id),
		StageID:   stage.ID,
		SkippedAt: time.Now(),
		Time:      time.Now(),
	})

	return nil
}

// GetStages возвращает все этапы плана
func (p *CropPlan) GetStages() []Stage {
	return append([]Stage(nil), p.stages...)
}

// GetNextStage возвращает следующий незавершенный этап
func (p *CropPlan) GetNextStage() *Stage {
	for _, stage := range p.stages {
		if stage.Status == StageStatusPending {
			return &stage
		}
	}
	return nil
}

// Жизненный цикл плана
func (p *CropPlan) Activate() error {
	if p.status != StatusDraft {
		return errors.New("only draft plans can be activated")
	}
	if len(p.stages) == 0 {
		return errors.New("cannot activate plan without stages")
	}

	now := time.Now()
	p.status = StatusActive
	p.startedAt = &now
	p.updatedAt = now

	p.AddEvent(CropPlanActivatedEvent{
		PlanID:      string(p.id),
		ActivatedAt: now,
	})

	return nil
}

func (p *CropPlan) Harvest(kg float64) error {
	if p.status != StatusActive {
		return errors.New("only active plans can be harvested")
	}
	if kg <= 0 {
		return errors.New("harvest weight must be positive")
	}

	p.harvestKg = kg
	p.complete()
	return nil
}

func (p *CropPlan) complete() {
	now := time.Now()
	p.status = StatusCompleted
	p.completedAt = &now
	p.updatedAt = now

	p.AddEvent(CropPlanCompletedEvent{
		PlanID:      string(p.id),
		CompletedAt: now,
		HarvestKg:   p.harvestKg,
	})
}

func (p *CropPlan) Cancel() error {
	if p.status == StatusCompleted || p.status == StatusCancelled {
		return errors.New("plan already completed or cancelled")
	}

	now := time.Now()
	p.status = StatusCancelled
	p.updatedAt = now

	p.AddEvent(CropPlanCancelledEvent{
		PlanID:      string(p.id),
		CancelledAt: now,
	})

	return nil
}
