package valueobject

import "testing"

func TestDimensionArea(t *testing.T) {
	d, _ := NewDimension(10, 2)
	if d.Area() != 20 {
		t.Fail()
	}
}
