package season

import (
	"time"
)

// SeasonPlan — план на сезон
type SeasonPlan struct {
	SeasonID        string                    `json:"season_id"`
	AreaAllocations map[string]AreaAllocation `json:"area_allocations"` // areaID -> allocation
	CreatedAt       time.Time                 `json:"created_at"`
	UpdatedAt       time.Time                 `json:"updated_at"`
}

// AreaAllocation — что и где сажаем
type AreaAllocation struct {
	AreaID          string           `json:"area_id"`           // ID места выращивания
	CropPlanID      string           `json:"crop_plan_id"`      // ID плана культуры
	CropPlanVersion int              `json:"crop_plan_version"` // Версия плана
	ExpectedYield   float64          `json:"expected_yield"`    // Ожидаемый урожай (кг/га)
	PlantingDate    time.Time        `json:"planting_date"`     // Дата посадки
	HarvestDate     *time.Time       `json:"harvest_date"`      // Дата уборки
	Status          AllocationStatus `json:"status"`
}

// AllocationStatus — статус выделения
type AllocationStatus string

const (
	AllocationStatusPlanned    AllocationStatus = "planned"     // запланировано
	AllocationStatusInProgress AllocationStatus = "in_progress" // в процессе
	AllocationStatusCompleted  AllocationStatus = "completed"   // завершено
	AllocationStatusCancelled  AllocationStatus = "cancelled"   // отменено
)

// NewSeasonPlan создаёт новый план на сезон
func NewSeasonPlan(seasonID string) *SeasonPlan {
	return &SeasonPlan{
		SeasonID:        seasonID,
		AreaAllocations: make(map[string]AreaAllocation),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

// AllocateArea выделяет место под культуру
func (p *SeasonPlan) AllocateArea(areaID string, allocation AreaAllocation) error {
	if _, exists := p.AreaAllocations[areaID]; exists {
		return ErrAreaAlreadyAllocated
	}

	p.AreaAllocations[areaID] = allocation
	p.UpdatedAt = time.Now()
	return nil
}

// UpdateAllocation обновляет выделение
func (p *SeasonPlan) UpdateAllocation(areaID string, allocation AreaAllocation) error {
	if _, exists := p.AreaAllocations[areaID]; !exists {
		return ErrAllocationNotFound
	}

	p.AreaAllocations[areaID] = allocation
	p.UpdatedAt = time.Now()
	return nil
}

// RemoveAllocation удаляет выделение
func (p *SeasonPlan) RemoveAllocation(areaID string) error {
	if _, exists := p.AreaAllocations[areaID]; !exists {
		return ErrAllocationNotFound
	}

	delete(p.AreaAllocations, areaID)
	p.UpdatedAt = time.Now()
	return nil
}

// GetAllocation возвращает выделение по ID места
func (p *SeasonPlan) GetAllocation(areaID string) (AreaAllocation, bool) {
	alloc, ok := p.AreaAllocations[areaID]
	return alloc, ok
}

// GetCropPlanID возвращает ID плана культуры для места
func (p *SeasonPlan) GetCropPlanID(areaID string) (string, bool) {
	alloc, ok := p.AreaAllocations[areaID]
	if !ok {
		return "", false
	}
	return alloc.CropPlanID, true
}
