package facility

import (
	"samurenkoroma/services/internal/modules/growing/domain/valueobject"
	"testing"

	"github.com/google/uuid"
)

func TestGreenhouseCannotAddSection(t *testing.T) {
	d, _ := valueobject.NewDimension(10, 10)
	gh := NewGreenhouseFacility(FacilityID(uuid.New().String()), "GH", d)

	sec := NewFieldBlock(GrowingAreaID(uuid.New().String()), "S", d)
	err := gh.AddBlock(sec)

	if err == nil {
		t.Fail()
	}
}
