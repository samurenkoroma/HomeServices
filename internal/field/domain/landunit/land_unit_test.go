package landunit

import (
	"samurenkoroma/services/internal/field/domain/valueobject"
	"testing"
)

func TestGreenhouseCannotAddSection(t *testing.T) {
	d, _ := valueobject.NewDimension(10, 10)
	gh := NewGreenhouse(1, "GH", d)

	sec := NewSection(1, "S", d)
	err := gh.AddSection(sec)

	if err == nil {
		t.Fail()
	}
}
