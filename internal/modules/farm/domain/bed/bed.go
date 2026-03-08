package bed

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/modules/farm/domain"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
)

type Bed struct {
	aggregate.BaseAggregate
	id        domain.GrowingAreaID
	name      string
	dimension valueobject.Dimension
}

func NewBed(id domain.GrowingAreaID, name string, dim valueobject.Dimension) *Bed {
	b := &Bed{id: id, name: name, dimension: dim}
	b.AddEvent(NewBedCreated(string(id)))
	return b
}

func (b *Bed) ID() domain.GrowingAreaID {
	return b.id
}

func (b *Bed) Name() string {
	return b.name
}

func (b *Bed) Dimension() valueobject.Dimension {
	return b.dimension
}

func RehydrateBed(
	id domain.GrowingAreaID,
	name string,
	dim valueobject.Dimension,
) *Bed {
	return &Bed{
		id:        id,
		name:      name,
		dimension: dim,
	}
}
