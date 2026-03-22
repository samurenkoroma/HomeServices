package cropcycle

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"time"
)

// CycleID — идентификатор цикла
type CycleID string

// Status — статус цикла
type Status string

const (
	StatusDraft     Status = "draft"     // черновик
	StatusActive    Status = "active"    // активен (посеян)
	StatusGrowing   Status = "growing"   // в процессе роста
	StatusHarvested Status = "harvested" // собран урожай
	StatusCompleted Status = "completed" // завершён
	StatusFailed    Status = "failed"    // не удался
	StatusCancelled Status = "cancelled" // отменён
)

// CropCycle — цикл выращивания культуры
type CropCycle struct {
	aggregate.Entity[CycleID]

	templateID      string // ID шаблона (из croptemplate)
	areaID          string // ID места выращивания (из cultivationarea)
	seasonID        string // ID сезона
	cropPlanID      string // ID плана культуры (из crop)
	cropPlanVersion int    // версия плана

	status     Status
	startedAt  *time.Time
	finishedAt *time.Time

	// Операции
	operations []Operation

	// Результаты
	yield   *Yield
	version int
}

// Operation — операция в цикле
type Operation struct {
	ID          string        `json:"id"`
	Type        OperationType `json:"type"`
	Description string        `json:"description"`
	Amount      float64       `json:"amount"`       // количество (литры, кг)
	Unit        string        `json:"unit"`         // единица измерения
	PerformedBy string        `json:"performed_by"` // кто выполнил
	Notes       string        `json:"notes"`
	CreatedAt   time.Time     `json:"created_at"`
}

// OperationType — тип операции
type OperationType string

const (
	OperationPlanting    OperationType = "planting"     // посадка
	OperationWatering    OperationType = "watering"     // полив
	OperationFertilizing OperationType = "fertilizing"  // подкормка
	OperationPestControl OperationType = "pest_control" // защита от вредителей
	OperationWeeding     OperationType = "weeding"      // прополка
	OperationHarvesting  OperationType = "harvesting"   // уборка
	OperationOther       OperationType = "other"        // другое
)

// Yield — урожай
type Yield struct {
	ActualWeight    float64   `json:"actual_weight"`    // кг
	EstimatedWeight float64   `json:"estimated_weight"` // кг (прогноз)
	HarvestedAt     time.Time `json:"harvested_at"`
	Quality         string    `json:"quality"` // качество (высокое, среднее, низкое)
	Notes           string    `json:"notes"`
}

// NewCropCycle создаёт новый цикл
func NewCropCycle(templateID, areaID, seasonID, cropPlanID string, planVersion int) *CropCycle {
	return &CropCycle{
		Entity:          aggregate.NewEntity(CycleID(types.NewUUID())),
		templateID:      templateID,
		areaID:          areaID,
		seasonID:        seasonID,
		cropPlanID:      cropPlanID,
		cropPlanVersion: planVersion,
		status:          StatusDraft,
		operations:      []Operation{},
	}
}

// Start — начать цикл (посадка)
func (c *CropCycle) Start() error {
	if c.status != StatusDraft {
		return ErrInvalidState
	}

	now := time.Now()
	c.status = StatusActive
	c.startedAt = &now
	c.Update()

	c.AddEvent(CropCycleStarted{
		CycleID:   string(c.Id),
		AreaID:    c.areaID,
		SeasonID:  c.seasonID,
		StartedAt: now,
	})

	return nil
}

// RecordOperation — записать операцию
func (c *CropCycle) RecordOperation(op Operation) error {
	if c.status != StatusActive && c.status != StatusGrowing {
		return ErrInvalidState
	}

	op.ID = types.NewUUID()
	op.CreatedAt = time.Now()
	c.operations = append(c.operations, op)
	c.Update()

	c.AddEvent(OperationRecorded{
		CycleID:   string(c.Id),
		Operation: op,
	})

	return nil
}

// RecordYield — записать урожай
func (c *CropCycle) RecordYield(actualWeight, estimatedWeight float64, quality, notes string) error {
	if c.status != StatusGrowing && c.status != StatusActive {
		return ErrInvalidState
	}

	now := time.Now()
	c.yield = &Yield{
		ActualWeight:    actualWeight,
		EstimatedWeight: estimatedWeight,
		HarvestedAt:     now,
		Quality:         quality,
		Notes:           notes,
	}
	c.status = StatusHarvested
	c.finishedAt = &now
	c.Update()

	c.AddEvent(CropCycleHarvested{
		CycleID: string(c.Id),
		Yield:   *c.yield,
	})

	return nil
}

// Complete — завершить цикл (после уборки)
func (c *CropCycle) Complete() error {
	if c.status != StatusHarvested {
		return ErrInvalidState
	}

	c.status = StatusCompleted
	c.Update()

	c.AddEvent(CropCycleCompleted{
		CycleID: string(c.Id),
	})

	return nil
}

// Fail — отметить цикл как неудачный
func (c *CropCycle) Fail(reason string) error {
	if c.status == StatusCompleted || c.status == StatusCancelled {
		return ErrInvalidState
	}

	c.status = StatusFailed
	c.finishedAt = &time.Time{}
	c.Update()

	c.AddEvent(CropCycleFailed{
		CycleID: string(c.Id),
		Reason:  reason,
	})

	return nil
}

// Cancel — отменить цикл
func (c *CropCycle) Cancel(reason string) error {
	if c.status == StatusCompleted {
		return ErrInvalidState
	}

	c.status = StatusCancelled
	c.Update()

	c.AddEvent(CropCycleCancelled{
		CycleID: string(c.Id),
		Reason:  reason,
	})

	return nil
}

// GetCurrentStage — получить текущий этап (на основе прошедшего времени)
func (c *CropCycle) GetCurrentStage(templateStages []cropplan.GrowthStage) *cropplan.GrowthStage {
	if c.startedAt == nil {
		return nil
	}

	daysSinceStart := int(time.Since(*c.startedAt).Hours() / 24)
	accumulated := 0

	for _, stage := range templateStages {
		accumulated += int(stage.Duration)
		if daysSinceStart <= accumulated {
			return &stage
		}
	}

	return nil
}

// GetProgress — получить прогресс выполнения в процентах
func (c *CropCycle) GetProgress(templateStages []cropplan.GrowthStage) float64 {
	if c.startedAt == nil {
		return 0
	}

	totalDuration := 0
	for _, stage := range templateStages {
		totalDuration += int(stage.Duration)
	}

	if totalDuration == 0 {
		return 0
	}

	daysSinceStart := int(time.Since(*c.startedAt).Hours() / 24)
	progress := float64(daysSinceStart) / float64(totalDuration) * 100

	if progress > 100 {
		return 100
	}
	return progress
}

// Геттеры
func (c *CropCycle) GetID() CycleID             { return c.Id }
func (c *CropCycle) GetTemplateID() string      { return c.templateID }
func (c *CropCycle) GetAreaID() string          { return c.areaID }
func (c *CropCycle) GetSeasonID() string        { return c.seasonID }
func (c *CropCycle) GetCropPlanID() string      { return c.cropPlanID }
func (c *CropCycle) GetCropPlanVersion() int    { return c.cropPlanVersion }
func (c *CropCycle) GetStatus() Status          { return c.status }
func (c *CropCycle) GetStartedAt() *time.Time   { return c.startedAt }
func (c *CropCycle) GetFinishedAt() *time.Time  { return c.finishedAt }
func (c *CropCycle) GetOperations() []Operation { return c.operations }
func (c *CropCycle) GetYield() *Yield           { return c.yield }
func (c *CropCycle) GetCreatedAt() time.Time    { return c.CreatedAt }
func (c *CropCycle) GetUpdatedAt() time.Time    { return c.UpdatedAt }

// Добавляем геттеры для yield полей
func (c *CropCycle) GetYieldActual() float64 {
	if c.yield == nil {
		return 0
	}
	return c.yield.ActualWeight
}

func (c *CropCycle) GetYieldEstimated() float64 {
	if c.yield == nil {
		return 0
	}
	return c.yield.EstimatedWeight
}

func (c *CropCycle) GetYieldQuality() string {
	if c.yield == nil {
		return ""
	}
	return c.yield.Quality
}

func (c *CropCycle) GetYieldNotes() string {
	if c.yield == nil {
		return ""
	}
	return c.yield.Notes
}

func (c *CropCycle) GetVersion() int {
	return c.version
}

// Обновляем Rehydrate
func (c *CropCycle) Rehydrate(
	id CycleID,
	templateID, areaID, seasonID, cropPlanID string,
	cropPlanVersion int,
	status Status,
	startedAt, finishedAt *time.Time,
	operations []Operation,
	yield *Yield,
	createdAt, updatedAt time.Time,
	version int,
) {
	c.Id = id
	c.templateID = templateID
	c.areaID = areaID
	c.seasonID = seasonID
	c.cropPlanID = cropPlanID
	c.cropPlanVersion = cropPlanVersion
	c.status = status
	c.startedAt = startedAt
	c.finishedAt = finishedAt
	c.operations = operations
	c.yield = yield
	c.CreatedAt = createdAt
	c.UpdatedAt = updatedAt
	c.version = version
}
