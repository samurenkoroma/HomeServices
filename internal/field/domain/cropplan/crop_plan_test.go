package cropplan

import "testing"

func TestHarvestTwice(t *testing.T) {
	cp := New("1", "bed1", "Tomato")
	cp.StartGrowing()
	cp.Harvest(100)

	err := cp.Harvest(50)
	if err == nil {
		t.Fail()
	}
}
