package domain

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/modules/farm"
	bedDomain "samurenkoroma/services/internal/modules/farm/bed/domain"
	blockDomain "samurenkoroma/services/internal/modules/farm/block/domain"
	"samurenkoroma/services/internal/modules/farm/valueobject"
)

type Field struct {
	aggregate.BaseAggregate

	id        farm.GrowingAreaID
	name      string
	dimension valueobject.Dimension

	blocks []*blockDomain.FieldBlock
	beds   []*bedDomain.Bed
}

func NewField(id farm.GrowingAreaID, name string, dim valueobject.Dimension) *Field {
	f := &Field{
		id:        id,
		name:      name,
		dimension: dim,
	}

	f.AddEvent(NewFieldCreated(string(id)))

	return f
}

func (f *Field) AddBlock(id farm.GrowingAreaID, name string, dim valueobject.Dimension) error {
	if f.containsBlock(id) {
		return farm.ErrDuplicateArea
	}

	block := blockDomain.NewFieldBlock(id, name, dim)

	f.blocks = append(f.blocks, &block)
	return nil
}

func (f *Field) AddBed(id farm.GrowingAreaID, name string, dim valueobject.Dimension) error {
	if f.containsBed(id) {
		return farm.ErrDuplicateArea
	}

	f.beds = append(f.beds, bedDomain.NewBed(id, name, dim))
	return nil
}

func (f *Field) containsBlock(id farm.GrowingAreaID) bool {
	for _, b := range f.blocks {
		if b.ID() == id {
			return true
		}
	}
	return false
}

func (f *Field) findBlock(id farm.GrowingAreaID) *blockDomain.FieldBlock {
	for _, b := range f.blocks {
		if b.ID() == id {
			return b
		}
	}
	return nil
}

func (f *Field) containsBed(id farm.GrowingAreaID) bool {
	for _, b := range f.beds {
		if b.ID() == id {
			return true
		}
	}
	return false
}

func (f *Field) ID() farm.GrowingAreaID {
	return f.id
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) Dimension() valueobject.Dimension {
	return f.dimension
}

func (f *Field) RehydrateAddBlock(b *blockDomain.FieldBlock) {
	f.blocks = append(f.blocks, b)
}

func (f *Field) RehydrateAddBed(bed *bedDomain.Bed) {
	f.beds = append(f.beds, bed)
}

func RehydrateGrowingFacility(
	id farm.GrowingAreaID,
	name string,
	dim valueobject.Dimension,
	sections []*blockDomain.FieldBlock,
	beds []*bedDomain.Bed,
) *Field {
	return &Field{
		id:        id,
		name:      name,
		dimension: dim,
		blocks:    sections,
		beds:      beds,
	}
}
