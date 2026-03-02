package facility

import (
	"samurenkoroma/services/internal/common/domain"
	"samurenkoroma/services/internal/growing/domain/valueobject"
	"time"
)

type GrowingFacility struct {
	domain.BaseAggregate

	id           FacilityID
	name         string
	facilityType FacilityType
	dimension    valueobject.Dimension

	blocks []*FieldBlock
	beds   []*Bed
}

func NewFieldFacility(id FacilityID, name string, dim valueobject.Dimension) *GrowingFacility {
	f := &GrowingFacility{
		id:           id,
		name:         name,
		facilityType: FieldFacility,
		dimension:    dim,
	}

	f.AddEvent(FacilityCreatedEvent{
		FacilityID: string(id),
		Time:       time.Now(),
	})

	return f
}

func NewGreenhouseFacility(id FacilityID, name string, dim valueobject.Dimension) *GrowingFacility {
	f := &GrowingFacility{
		id:           id,
		name:         name,
		facilityType: GreenhouseFacility,
		dimension:    dim,
	}

	f.AddEvent(FacilityCreatedEvent{
		FacilityID: string(id),
		Time:       time.Now(),
	})
	return f
}

func (f *GrowingFacility) AddBlock(id GrowingAreaID, name string, dim valueobject.Dimension) error {

	if f.facilityType != FieldFacility {
		return ErrBlockNotAllowed
	}

	if f.containsBlock(id) {
		return ErrDuplicateArea
	}

	block := &FieldBlock{
		id:        id,
		name:      name,
		dimension: dim,
	}

	f.blocks = append(f.blocks, block)
	return nil
}

func (f *GrowingFacility) AddBedToBlock(
	blockID GrowingAreaID,
	bedID GrowingAreaID,
	name string,
	dim valueobject.Dimension,
) error {

	if f.facilityType != FieldFacility {
		return ErrBedMustHaveParentBlock
	}

	block := f.findBlock(blockID)
	if block == nil {
		return ErrAreaNotFound
	}

	if block.ContainsBed(bedID) {
		return ErrDuplicateArea
	}

	block.AddBed(NewBed(bedID, name, dim))

	return nil
}

func (f *GrowingFacility) AddBed(
	id GrowingAreaID,
	name string,
	dim valueobject.Dimension,
) error {

	if f.facilityType != GreenhouseFacility {
		return ErrBedNotAllowed
	}

	if f.containsBed(id) {
		return ErrDuplicateArea
	}

	f.beds = append(f.beds, NewBed(id, name, dim))
	return nil
}

func (f *GrowingFacility) containsBlock(id GrowingAreaID) bool {
	for _, b := range f.blocks {
		if b.id == id {
			return true
		}
	}
	return false
}

func (f *GrowingFacility) findBlock(id GrowingAreaID) *FieldBlock {
	for i := range f.blocks {
		if f.blocks[i].id == id {
			return f.blocks[i]
		}
	}
	return nil
}

func (f *GrowingFacility) containsBed(id GrowingAreaID) bool {
	for _, b := range f.beds {
		if b.id == id {
			return true
		}
	}
	return false
}

func (f *GrowingFacility) ID() FacilityID {
	return f.id
}

func (f *GrowingFacility) Name() string {
	return f.name
}

func (f *GrowingFacility) FacilityType() FacilityType {
	return f.facilityType
}

func (f *GrowingFacility) Dimension() valueobject.Dimension {
	return f.dimension
}

func (f *GrowingFacility) RehydrateAddBlock(b *FieldBlock) {
	f.blocks = append(f.blocks, b)
}
func (f *GrowingFacility) RehydrateAddBedToSection(blockID GrowingAreaID, bed *Bed) {
	for _, block := range f.blocks {
		if block.id == blockID {
			block.RehydrateAddBed(bed)
			return
		}
	}
}

func (f *GrowingFacility) RehydrateAddBed(bed *Bed) {
	f.beds = append(f.beds, bed)
}

func RehydrateGrowingFacility(
	id FacilityID,
	name string,
	facilityType FacilityType,
	dim valueobject.Dimension,
	sections []*FieldBlock,
	beds []*Bed,
) *GrowingFacility {
	return &GrowingFacility{
		id:           id,
		name:         name,
		facilityType: facilityType,
		dimension:    dim,
		blocks:       sections,
		beds:         beds,
	}
}
