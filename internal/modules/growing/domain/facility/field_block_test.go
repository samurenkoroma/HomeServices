package facility

import (
	"samurenkoroma/services/internal/modules/growing/domain/valueobject"
	"testing"

	"github.com/google/uuid"
)

func TestSectionAreaExceeded(t *testing.T) {
	d, _ := valueobject.NewDimension(10, 10)
	section := NewFieldBlock(GrowingAreaID(uuid.New().String()), "A", d)

	bedDim, _ := valueobject.NewDimension(20, 20)
	bed := NewBed(GrowingAreaID(uuid.New().String()), "Bed1", bedDim)

	section.AddBed(bed)

}
