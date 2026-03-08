package domain

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/modules/farm"
	bed "samurenkoroma/services/internal/modules/farm/bed/domain"
	block "samurenkoroma/services/internal/modules/farm/block/domain"
	"samurenkoroma/services/internal/modules/farm/valueobject"
)

type Greenhouse struct {
	aggregate.BaseAggregate

	id        farm.GrowingAreaID
	name      string
	dimension valueobject.Dimension

	blocks []*block.FieldBlock
	beds   []*bed.Bed
}

func NewGreenhouse(id farm.GrowingAreaID, name string, dim valueobject.Dimension) *Greenhouse {
	f := &Greenhouse{
		id:        id,
		name:      name,
		dimension: dim,
	}

	f.AddEvent(NewGreenhouseCreated(string(id)))
	return f
}

func (f *Greenhouse) AddBed(id farm.GrowingAreaID, name string, dim valueobject.Dimension) error {

	if f.containsBed(id) {
		return farm.ErrDuplicateArea
	}

	f.beds = append(f.beds, bed.NewBed(id, name, dim))
	return nil
}

func (f *Greenhouse) containsBed(id farm.GrowingAreaID) bool {
	for _, b := range f.beds {
		if b.ID() == id {
			return true
		}
	}
	return false
}

func (f *Greenhouse) ID() farm.GrowingAreaID {
	return f.id
}
func (f *Greenhouse) Name() string {
	return f.name
}
func (f *Greenhouse) Dimension() valueobject.Dimension {
	return f.dimension
}
func (f *Greenhouse) RehydrateAddBed(bed *bed.Bed) {
	f.beds = append(f.beds, bed)
}

func RehydrateGreenhouse(
	id farm.GrowingAreaID,
	name string,
	dim valueobject.Dimension,
	beds []*bed.Bed,
) *Greenhouse {
	return &Greenhouse{
		id:        id,
		name:      name,
		dimension: dim,
		beds:      beds,
	}
}
