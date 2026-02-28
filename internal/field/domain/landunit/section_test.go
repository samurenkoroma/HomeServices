package landunit

import (
	"samurenkoroma/services/internal/field/domain/valueobject"
	"testing"
)

func TestSectionAreaExceeded(t *testing.T) {
	d, _ := valueobject.NewDimension(10, 10)
	section := NewSection(1, "A", d)

	bedDim, _ := valueobject.NewDimension(20, 20)
	bed := NewBed(1, "Bed1", bedDim)

	err := section.AddBed(bed)
	if err == nil {
		t.Fail()
	}
}
