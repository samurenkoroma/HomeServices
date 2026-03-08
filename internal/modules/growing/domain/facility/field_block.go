package facility

import (
	"samurenkoroma/services/internal/modules/growing/domain/valueobject"
)

type FieldBlock struct {
	id        GrowingAreaID
	name      string
	dimension valueobject.Dimension
	beds      []*Bed
}

func NewFieldBlock(id GrowingAreaID, name string, dim valueobject.Dimension) FieldBlock {
	return FieldBlock{
		id:        id,
		name:      name,
		dimension: dim,
	}
}

func (b *FieldBlock) Dimension() valueobject.Dimension {
	return b.dimension
}
func (b *FieldBlock) AddBed(bed *Bed) {
	b.beds = append(b.beds, bed)
}

func (b *FieldBlock) ContainsBed(id GrowingAreaID) bool {
	for _, bed := range b.beds {
		if bed.id == id {
			return true
		}
	}
	return false
}

func (b *FieldBlock) Name() string {
	return b.name
}
func (b *FieldBlock) RehydrateAddBed(bed *Bed) {
	b.beds = append(b.beds, bed)
}

func RehydrateBlock(
	id GrowingAreaID,
	name string,
	dim valueobject.Dimension,
) FieldBlock {

	return FieldBlock{
		id:        id,
		name:      name,
		dimension: dim,
	}
}
