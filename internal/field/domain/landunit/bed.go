package landunit

import "samurenkoroma/services/internal/field/domain/valueobject"

type Bed struct {
	id        BedID
	name      string
	dimension valueobject.Dimension
}

func NewBed(id BedID, name string, dim valueobject.Dimension) *Bed {
	return &Bed{id: id, name: name, dimension: dim}
}

func (b *Bed) ID() BedID {
	return b.id
}

func (b *Bed) Area() float64 {
	return b.dimension.Area()
}

func (b *Bed) Name() string {
	return b.name
}

func (b *Bed) Dimension() valueobject.Dimension {
	return b.dimension
}
