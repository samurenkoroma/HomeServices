package domain

import (
	"samurenkoroma/services/internal/modules/farm"
	domain2 "samurenkoroma/services/internal/modules/farm/bed/domain"
	"samurenkoroma/services/internal/modules/farm/valueobject"
	"testing"

	"github.com/google/uuid"
)

func TestSectionAreaExceeded(t *testing.T) {
	d, _ := valueobject.NewDimension(10, 10)
	section := NewFieldBlock(farm.GrowingAreaID(uuid.New().String()), "A", d)

	bedDim, _ := valueobject.NewDimension(20, 20)
	bed := domain2.NewBed(farm.GrowingAreaID(uuid.New().String()), "Bed1", bedDim)

	section.AddBed(bed)

}
