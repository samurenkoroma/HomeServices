package landunit

import (
	"samurenkoroma/services/internal/field/domain/valueobject"
)

type Section struct {
	id        SectionID
	name      string
	dimension valueobject.Dimension
	beds      []*Bed
}

func NewSection(id SectionID, name string, dim valueobject.Dimension) *Section {
	return &Section{id: id, name: name, dimension: dim}
}

func (s *Section) AddBed(b *Bed) error {
	if s.totalBedArea()+b.Area() > s.dimension.Area() {
		return ErrAreaExceeded
	}
	s.beds = append(s.beds, b)
	return nil
}

func (s *Section) totalBedArea() float64 {
	total := 0.0
	for _, b := range s.beds {
		total += b.Area()
	}
	return total
}

func (s *Section) ID() SectionID {
	return s.id
}

func (s *Section) Dimension() valueobject.Dimension {
	return s.dimension
}
func (s *Section) Beds() []*Bed {
	return append([]*Bed(nil), s.beds...)
}

func (s *Section) Name() string {
	return s.name
}
func (s *Section) rehydrateAddBed(b *Bed) {
	s.beds = append(s.beds, b)
}
