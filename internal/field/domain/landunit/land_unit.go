package landunit

import (
	"samurenkoroma/services/internal/domain/shared"
	"samurenkoroma/services/internal/field/domain/valueobject"
)

type LandUnit struct {
	shared.BaseAggregate
	id        LandUnitID
	name      string
	unitType  LandUnitType
	dimension valueobject.Dimension
	sections  []*Section
	beds      []*Bed
}

func NewField(id LandUnitID, name string, dim valueobject.Dimension) *LandUnit {
	return &LandUnit{id: id, name: name, unitType: Field, dimension: dim}
}

func NewGreenhouse(id LandUnitID, name string, dim valueobject.Dimension) *LandUnit {
	return &LandUnit{id: id, name: name, unitType: Greenhouse, dimension: dim}
}

func (l *LandUnit) AddSection(s *Section) error {
	if l.unitType != Field {
		return ErrInvalidUnitType
	}
	if l.totalSectionArea()+s.dimension.Area() > l.dimension.Area() {
		return ErrAreaExceeded
	}
	l.sections = append(l.sections, s)
	return nil
}

func (l *LandUnit) AddBedToSection(sectionID SectionID, b *Bed) error {
	if l.unitType != Field {
		return ErrInvalidUnitType
	}
	for _, s := range l.sections {
		if s.ID() == sectionID {
			return s.AddBed(b)
		}
	}
	return ErrSectionNotFound
}

func (l *LandUnit) AddBedToGreenhouse(b *Bed) error {
	if l.unitType != Greenhouse {
		return ErrInvalidUnitType
	}
	if l.totalBedArea()+b.Area() > l.dimension.Area() {
		return ErrAreaExceeded
	}
	l.beds = append(l.beds, b)
	return nil
}

func (l *LandUnit) totalSectionArea() float64 {
	total := 0.0
	for _, s := range l.sections {
		total += s.dimension.Area()
	}
	return total
}

func (l *LandUnit) totalBedArea() float64 {
	total := 0.0
	for _, b := range l.beds {
		total += b.Area()
	}
	return total
}
func (l *LandUnit) ID() LandUnitID {
	return l.id
}

func (l *LandUnit) Name() string {
	return l.name
}

func (l *LandUnit) Type() LandUnitType {
	return l.unitType
}

func (l *LandUnit) Dimension() valueobject.Dimension {
	return l.dimension
}

func (l *LandUnit) Sections() []*Section {
	return append([]*Section(nil), l.sections...)
}

func (l *LandUnit) Beds() []*Bed {
	return append([]*Bed(nil), l.beds...)
}

func RehydrateLandUnit(
	id LandUnitID,
	name string,
	unitType LandUnitType,
	dim valueobject.Dimension,
	sections []*Section,
	beds []*Bed,
) *LandUnit {
	return &LandUnit{
		id:        id,
		name:      name,
		unitType:  unitType,
		dimension: dim,
		sections:  sections,
		beds:      beds,
	}
}

func RehydrateSection(id SectionID, name string, dim valueobject.Dimension) *Section {
	return &Section{
		id:        id,
		name:      name,
		dimension: dim,
		beds:      nil,
	}
}

func (l *LandUnit) rehydrateAddSection(s *Section) {
	l.sections = append(l.sections, s)
}

func (l *LandUnit) rehydrateAddBedToSection(sectionID SectionID, b *Bed) {
	for _, s := range l.sections {
		if s.ID() == sectionID {
			s.rehydrateAddBed(b)
			return
		}
	}
}

func (l *LandUnit) rehydrateAddBed(b *Bed) {
	l.beds = append(l.beds, b)
}
