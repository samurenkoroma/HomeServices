package production

import (
	"time"

	"github.com/samurenkoroma/HomeServices/internal/domain/shared/valueobjects"
	"github.com/samurenkoroma/HomeServices/internal/domain/spatial"
)

// GrowingCycle - Aggregate for crop growing cycle
type GrowingCycle struct {
	ID              string                     `json:"id"`
	ProductionUnit  *spatial.ProductionUnit   `json:"production_unit"`
	Crop            string                     `json:"crop"`
	Variety         string                     `json:"variety"`
	StartDate       time.Time                  `json:"start_date"`
	ExpectedEndDate time.Time                  `json:"expected_end_date"`
	Status          string                     `json:"status"`
}

func NewGrowingCycle(id string, pu *spatial.ProductionUnit, crop, variety string, start, end time.Time) *GrowingCycle {
	return &GrowingCycle{
		ID:              id,
		ProductionUnit:  pu,
		Crop:            crop,
		Variety:         variety,
		StartDate:       start,
		ExpectedEndDate: end,
		Status:          "planned",
	}
}
