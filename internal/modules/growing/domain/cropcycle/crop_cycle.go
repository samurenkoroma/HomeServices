package cropcycle

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"time"
)

type CycleStatus string

const (
	CycleStatusDraft     CycleStatus = "draft"
	CycleStatusActive    CycleStatus = "active"
	CycleStatusGrowing   CycleStatus = "growing"
	CycleStatusHarvested CycleStatus = "harvested"
	CycleStatusCompleted CycleStatus = "completed"
	CycleStatusFailed    CycleStatus = "failed"
	CycleStatusCancelled CycleStatus = "cancelled"
)

// CropCycle - цикл выращивания культуры
type CropCycle struct {
	aggregate.BaseAggregate

	ID              string
	AreaID          string // ID из cultivationarea
	SeasonID        string
	CropPlanID      string
	CropPlanVersion int

	Status     CycleStatus
	StartedAt  *time.Time
	FinishedAt *time.Time

	// Операции
	Operations []Operation

	// Результаты
	Yield *Yield

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Yield - урожай
type Yield struct {
	ActualWeight    float64
	EstimatedWeight float64
	HarvestedAt     time.Time
	Quality         string
	Notes           string
}

func NewCropCycle(areaID, seasonID, cropPlanID string) *CropCycle {
	return &CropCycle{
		ID:         generateID(),
		AreaID:     areaID,
		SeasonID:   seasonID,
		CropPlanID: cropPlanID,
		Status:     CycleStatusDraft,
		Operations: []Operation{},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (c *CropCycle) Start() error {
	if c.Status != CycleStatusDraft {
		return ErrInvalidState
	}

	now := time.Now()
	c.Status = CycleStatusActive
	c.StartedAt = &now
	c.UpdatedAt = now

	c.AddEvent(CropCycleStarted{
		CycleID:   c.ID,
		AreaID:    c.AreaID,
		SeasonID:  c.SeasonID,
		StartedAt: now,
	})

	return nil
}

func (c *CropCycle) RecordOperation(op Operation) error {
	if c.Status != CycleStatusActive && c.Status != CycleStatusGrowing {
		return ErrInvalidState
	}

	op.CreatedAt = time.Now()
	c.Operations = append(c.Operations, op)
	c.UpdatedAt = time.Now()

	c.AddEvent(OperationRecorded{
		CycleID:   c.ID,
		Operation: op,
	})

	return nil
}

func (c *CropCycle) RecordYield(weight float64, quality string) error {
	if c.Status != CycleStatusGrowing && c.Status != CycleStatusActive {
		return ErrInvalidState
	}

	now := time.Now()
	c.Yield = &Yield{
		ActualWeight: weight,
		HarvestedAt:  now,
		Quality:      quality,
	}
	c.Status = CycleStatusHarvested
	c.FinishedAt = &now
	c.UpdatedAt = now

	c.AddEvent(CropCycleHarvested{
		CycleID: c.ID,
		Yield:   *c.Yield,
	})

	return nil
}

func (c *CropCycle) Complete() error {
	if c.Status != CycleStatusHarvested {
		return ErrInvalidState
	}

	c.Status = CycleStatusCompleted
	c.UpdatedAt = time.Now()

	c.AddEvent(CropCycleCompleted{
		CycleID: c.ID,
	})

	return nil
}
