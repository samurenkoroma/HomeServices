package landunit

import (
	"samurenkoroma/services/internal/domain/shared"
	"samurenkoroma/services/internal/field/domain/valueobject"
)

type GrowingFacility struct {
	shared.BaseAggregate
	id        LandUnitID
	name      string
	spaceType LandSpaceType
	dimension valueobject.Dimension
	sections  []*Section
	beds      []*Bed
}

func NewField(id LandUnitID, name string, dim valueobject.Dimension) *GrowingFacility {
	return &GrowingFacility{id: id, name: name, spaceType: Field, dimension: dim}
}

func NewGreenhouse(id LandUnitID, name string, dim valueobject.Dimension) *GrowingFacility {
	return &GrowingFacility{id: id, name: name, spaceType: Greenhouse, dimension: dim}
}

func (l *GrowingFacility) AddSection(s *Section) error {
	if l.spaceType != Field {
		return ErrInvalidSpaceType
	}
	if l.totalSectionArea()+s.dimension.Area() > l.dimension.Area() {
		return ErrAreaExceeded
	}
	l.sections = append(l.sections, s)
	return nil
}

func (l *GrowingFacility) AddBedToSection(sectionID SectionID, b *Bed) error {
	if l.spaceType != Field {
		return ErrInvalidSpaceType
	}
	for _, s := range l.sections {
		if s.ID() == sectionID {
			return s.AddBed(b)
		}
	}
	return ErrSectionNotFound
}

func (l *GrowingFacility) AddBedToGreenhouse(b *Bed) error {
	if l.spaceType != Greenhouse {
		return ErrInvalidSpaceType
	}
	if l.totalBedArea()+b.Area() > l.dimension.Area() {
		return ErrAreaExceeded
	}
	l.beds = append(l.beds, b)
	return nil
}

func (l *GrowingFacility) totalSectionArea() float64 {
	total := 0.0
	for _, s := range l.sections {
		total += s.dimension.Area()
	}
	return total
}

func (l *GrowingFacility) totalBedArea() float64 {
	total := 0.0
	for _, b := range l.beds {
		total += b.Area()
	}
	return total
}
func (l *GrowingFacility) ID() LandUnitID {
	return l.id
}

func (l *GrowingFacility) Name() string {
	return l.name
}

func (l *GrowingFacility) SpaceType() LandSpaceType {
	return l.spaceType
}

func (l *GrowingFacility) Dimension() valueobject.Dimension {
	return l.dimension
}

func (l *GrowingFacility) Sections() []*Section {
	return append([]*Section(nil), l.sections...)
}

func (l *GrowingFacility) Beds() []*Bed {
	return append([]*Bed(nil), l.beds...)
}

func RehydrateLandUnit(
	id LandUnitID,
	name string,
	spaceType LandSpaceType,
	dim valueobject.Dimension,
	sections []*Section,
	beds []*Bed,
) *GrowingFacility {
	return &GrowingFacility{
		id:        id,
		name:      name,
		spaceType: spaceType,
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

func (l *GrowingFacility) rehydrateAddSection(s *Section) {
	l.sections = append(l.sections, s)
}

func (l *GrowingFacility) rehydrateAddBedToSection(sectionID SectionID, b *Bed) {
	for _, s := range l.sections {
		if s.ID() == sectionID {
			s.rehydrateAddBed(b)
			return
		}
	}
}

func (l *GrowingFacility) rehydrateAddBed(b *Bed) {
	l.beds = append(l.beds, b)
}
