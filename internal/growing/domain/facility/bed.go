package facility

import "samurenkoroma/services/internal/growing/domain/valueobject"

type Bed struct {
	id        GrowingAreaID
	name      string
	dimension valueobject.Dimension
}

func NewBed(id GrowingAreaID, name string, dim valueobject.Dimension) Bed {
	return Bed{id: id, name: name, dimension: dim}
}

func (b *Bed) ID() GrowingAreaID {
	return b.id
}

func (b *Bed) Name() string {
	return b.name
}

func (b *Bed) Dimension() valueobject.Dimension {
	return b.dimension
}
