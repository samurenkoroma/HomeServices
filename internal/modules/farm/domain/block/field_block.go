package block

import (
	"samurenkoroma/services/internal/modules/farm/domain"
	domain2 "samurenkoroma/services/internal/modules/farm/domain/bed"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
)

type FieldBlock struct {
	id        domain.GrowingAreaID
	name      string
	dimension valueobject.Dimension
	beds      []*domain2.Bed
}

func NewFieldBlock(id domain.GrowingAreaID, name string, dim valueobject.Dimension) FieldBlock {
	return FieldBlock{
		id:        id,
		name:      name,
		dimension: dim,
	}
}

func (b *FieldBlock) Dimension() valueobject.Dimension {
	return b.dimension
}
func (b *FieldBlock) AddBed(bed *domain2.Bed) {
	b.beds = append(b.beds, bed)
}

func (b *FieldBlock) ContainsBed(id domain.GrowingAreaID) bool {
	for _, bed := range b.beds {
		if bed.ID() == id {
			return true
		}
	}
	return false
}

func (b *FieldBlock) Name() string {
	return b.name
}
func (b *FieldBlock) ID() domain.GrowingAreaID {
	return b.id
}
func (b *FieldBlock) RehydrateAddBed(bed *domain2.Bed) {
	b.beds = append(b.beds, bed)
}

func RehydrateBlock(id domain.GrowingAreaID, name string, dim valueobject.Dimension) FieldBlock {
	return FieldBlock{
		id:        id,
		name:      name,
		dimension: dim,
	}
}
