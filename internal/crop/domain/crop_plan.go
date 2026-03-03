package domain

import (
	common "samurenkoroma/services/internal/common/domain"
	"samurenkoroma/services/internal/shared/events"
)

type CropPlanID string

type CropPlan struct {
	common.BaseAggregate
	id       CropPlanID
	cropType CropTypeID
	variety  *CropVarietyID

	name     string
	duration int
	version  int
	status   CropPlanStatus

	stages   []GrowthStage
	rotation []CropRotationRule

	envReq  EnvironmentalRequirements
	nutrReq NutrientRequirements
}

func (c *CropPlan) RotationRules() []CropRotationRule {
	return c.rotation
}

func (c *CropPlan) ID() CropPlanID {
	return c.id
}

func (c *CropPlan) CropTypeID() CropTypeID {
	return c.cropType
}

func (c *CropPlan) VarietyID() *CropVarietyID {
	return c.variety
}

func (c *CropPlan) Name() string {
	return c.name
}

func (c *CropPlan) Duration() int {
	return c.duration
}

func (c *CropPlan) Version() int {
	return c.version
}

func (c *CropPlan) Status() CropPlanStatus {
	return c.status
}

func (c *CropPlan) Env() EnvironmentalRequirements {
	return c.envReq
}

func (c *CropPlan) Nutrients() NutrientRequirements {
	return c.nutrReq
}

func NewCropPlan(id CropPlanID, cropType CropTypeID, name string, duration int) (*CropPlan, error) {
	if duration <= 0 {
		return nil, ErrInvalidDuration
	}

	plan := &CropPlan{
		id:       id,
		cropType: cropType,
		name:     name,
		duration: duration,
		version:  1,
		status:   StatusDraft,
	}
	plan.AddEvent(events.NewCropPlanCreated(string(id), string(cropType)))
	return plan, nil
}

// AddStage добавляет этап в план
func (c *CropPlan) AddStage(stage GrowthStage) error {
	if c.status == StatusPublished {
		return ErrCannotModifyPublished
	}

	// Проверяем уникальность порядка
	for _, s := range c.stages {
		if s.order == stage.order {
			return ErrStageOrderDuplicate
		}
	}
	c.stages = append(c.stages, stage)

	c.AddEvent(events.NewGrowthStageAdded(string(c.id), stage.Order(), stage.Name()))
	return nil
}

func (c *CropPlan) ValidateDuration() error {

	total := 0
	for _, s := range c.stages {
		total += s.duration
	}

	if total != c.duration {
		return ErrStageDurationMismatch
	}

	return nil
}

func (c *CropPlan) Publish() error {

	if err := c.ValidateDuration(); err != nil {
		return err
	}

	c.status = StatusPublished
	c.AddEvent(events.NewCropPlanPublished(string(c.id), c.version))
	return nil
}

func (c *CropPlan) AddRotationRule(rule CropRotationRule) error {

	if c.status == StatusPublished {
		return ErrCannotModifyPublished
	}

	for _, r := range c.rotation {
		if r.predecessor == rule.predecessor {
			return ErrRotationDuplicate
		}
	}

	c.rotation = append(c.rotation, rule)
	//c.AddEvent(NewRotationRuleAdded(c.ID(), ))
	return nil
}
func (c *CropPlan) Rehydrate(version int, status CropPlanStatus) {
	c.version = version
	c.status = status
}

func (c *CropPlan) Stages() []GrowthStage {
	return c.stages
}

//
//// GetStage возвращает этап по ID
//func (p *CropPlan) GetStage(stageID string) (*Stage, error) {
//	for i := range p.stages {
//		if p.stages[i].ID == stageID {
//			return &p.stages[i], nil
//		}
//	}
//	return nil, ErrStageNotFound
//}
//
//// StartStage начинает этап
//func (p *CropPlan) StartStage(stageID string) error {
//	if p.status != StatusActive {
//		return errors.New("can only start stages in active plan")
//	}
//
//	stage, err := p.GetStage(stageID)
//	if err != nil {
//		return err
//	}
//
//	if err := stage.Start(); err != nil {
//		return err
//	}
//
//	p.updatedAt = time.Now()
//
//	p.AddEvent(StageStartedEvent{
//		PlanID:    string(p.id),
//		StageID:   stage.ID,
//		StartedAt: *stage.StartDate,
//		Time:      time.Now(),
//	})
//
//	return nil
//}
//
//// CompleteStage завершает этап
//func (p *CropPlan) CompleteStage(stageID string) error {
//	if p.status != StatusActive {
//		return errors.New("can only complete stages in active plan")
//	}
//
//	stage, err := p.GetStage(stageID)
//	if err != nil {
//		return err
//	}
//
//	if err := stage.Complete(); err != nil {
//		return err
//	}
//
//	p.updatedAt = time.Now()
//
//	p.AddEvent(StageCompletedEvent{
//		PlanID:      string(p.id),
//		StageID:     stage.ID,
//		CompletedAt: *stage.EndDate,
//		Time:        time.Now(),
//	})
//
//	// Проверяем, все ли этапы завершены
//	allCompleted := true
//	for _, s := range p.stages {
//		if !s.IsFinished() {
//			allCompleted = false
//			break
//		}
//	}
//
//	if allCompleted && p.status == StatusActive {
//		p.complete()
//	}
//
//	return nil
//}
//
//// SkipStage пропускает этап
//func (p *CropPlan) SkipStage(stageID string) error {
//	if p.status != StatusActive && p.status != StatusDraft {
//		return errors.New("cannot skip stages in this status")
//	}
//
//	stage, err := p.GetStage(stageID)
//	if err != nil {
//		return err
//	}
//
//	if err := stage.Skip(); err != nil {
//		return err
//	}
//
//	p.updatedAt = time.Now()
//
//	p.AddEvent(StageSkippedEvent{
//		PlanID:    string(p.id),
//		StageID:   stage.ID,
//		SkippedAt: time.Now(),
//		Time:      time.Now(),
//	})
//
//	return nil
//}
//
//// GetStages возвращает все этапы плана
//func (p *CropPlan) GetStages() []Stage {
//	return append([]Stage(nil), p.stages...)
//}
//
//// GetNextStage возвращает следующий незавершенный этап
//func (p *CropPlan) GetNextStage() *Stage {
//	for _, stage := range p.stages {
//		if stage.Status == StageStatusPending {
//			return &stage
//		}
//	}
//	return nil
//}
//
//// Жизненный цикл плана
//func (p *CropPlan) Activate() error {
//	if p.status != StatusDraft {
//		return errors.New("only draft plans can be activated")
//	}
//	if len(p.stages) == 0 {
//		return errors.New("cannot activate plan without stages")
//	}
//
//	now := time.Now()
//	p.status = StatusActive
//	p.startedAt = &now
//	p.updatedAt = now
//
//	p.AddEvent(CropPlanActivatedEvent{
//		PlanID:      string(p.id),
//		ActivatedAt: now,
//	})
//
//	return nil
//}
//
//func (p *CropPlan) Harvest(kg float64) error {
//	if p.status != StatusActive {
//		return errors.New("only active plans can be harvested")
//	}
//	if kg <= 0 {
//		return errors.New("harvest weight must be positive")
//	}
//
//	p.harvestKg = kg
//	p.complete()
//	return nil
//}
//
//func (p *CropPlan) complete() {
//	now := time.Now()
//	p.status = StatusCompleted
//	p.completedAt = &now
//	p.updatedAt = now
//
//	p.AddEvent(CropPlanCompletedEvent{
//		PlanID:      string(p.id),
//		CompletedAt: now,
//		HarvestKg:   p.harvestKg,
//	})
//}
//
//func (p *CropPlan) Cancel() error {
//	if p.status == StatusCompleted || p.status == StatusCancelled {
//		return errors.New("plan already completed or cancelled")
//	}
//
//	now := time.Now()
//	p.status = StatusCancelled
//	p.updatedAt = now
//
//	p.AddEvent(CropPlanCancelledEvent{
//		PlanID:      string(p.id),
//		CancelledAt: now,
//	})
//
//	return nil
//}
