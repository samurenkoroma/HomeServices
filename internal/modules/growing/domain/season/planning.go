package season

import (
	"time"
)

// SeasonPlan - план на сезон
type SeasonPlan struct {
	SeasonID        string
	AreaAllocations map[string]AreaAllocation // areaID -> allocation
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// AreaAllocation - что и где сажаем
type AreaAllocation struct {
	AreaID        string
	CropPlanID    string
	ExpectedYield float64
	PlantingDate  time.Time
	HarvestDate   *time.Time
	Status        AllocationStatus
}

type AllocationStatus string

const (
	AllocationStatusPlanned    AllocationStatus = "planned"
	AllocationStatusInProgress AllocationStatus = "in_progress"
	AllocationStatusCompleted  AllocationStatus = "completed"
	AllocationStatusCancelled  AllocationStatus = "cancelled"
)

func NewSeasonPlan(seasonID string) *SeasonPlan {
	return &SeasonPlan{
		SeasonID:        seasonID,
		AreaAllocations: make(map[string]AreaAllocation),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (p *SeasonPlan) AllocateArea(areaID string, allocation AreaAllocation) error {
	if _, exists := p.AreaAllocations[areaID]; exists {
		return ErrAreaAlreadyAllocated
	}

	p.AreaAllocations[areaID] = allocation
	p.UpdatedAt = time.Now()

	return nil
}

func (p *SeasonPlan) UpdateAllocation(areaID string, allocation AreaAllocation) error {
	if _, exists := p.AreaAllocations[areaID]; !exists {
		return ErrAllocationNotFound
	}

	p.AreaAllocations[areaID] = allocation
	p.UpdatedAt = time.Now()

	return nil
}
