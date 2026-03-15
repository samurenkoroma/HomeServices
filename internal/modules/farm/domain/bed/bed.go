package bed

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
)

type Bed struct {
	aggregate.Entity[types.BedId]
	Name      string
	Dimension valueobject.Dimension
	ParentId  string
	RootId    string
	Geom      spatial.GeoJSON
	valueobject.Additions
}

func NewBed(id types.BedId, name string, dim valueobject.Dimension) *Bed {
	b := &Bed{
		Entity: aggregate.NewEntity(types.BedId(types.NewUUID())),
		Name:   name, Dimension: dim,
	}
	b.AddEvent(NewBedCreated(string(id)))
	return b
}

//func RehydrateBed(id types.GrowingAreaID, name string, dim valueobject.Dimension) *Bed {
//	return &Bed{
//		Id:        id,
//		Name:      name,
//		Dimension: dim,
//	}
//}
