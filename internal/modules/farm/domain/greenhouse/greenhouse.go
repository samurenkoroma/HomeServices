package greenhouse

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain"
	"samurenkoroma/services/internal/modules/farm/domain/bed"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
)

type Greenhouse struct {
	aggregate.Entity[types.GreenhouseId]

	Name      string
	Dimension valueobject.Dimension

	Beds []*bed.Bed
	Geom spatial.GeoJSON
}

func NewGreenhouse(id types.GreenhouseId, name string, dim valueobject.Dimension, geom spatial.GeoJSON) *Greenhouse {
	f := &Greenhouse{
		Entity:    aggregate.NewEntity(types.GreenhouseId(types.NewUUID())),
		Name:      name,
		Dimension: dim,
		Geom:      geom,
	}

	f.AddEvent(NewGreenhouseCreated(string(id)))
	return f
}

func (f *Greenhouse) AddBed(id types.BedId, name string, dim valueobject.Dimension) error {

	if f.containsBed(id) {
		return domain.ErrDuplicateArea
	}

	f.Beds = append(f.Beds, bed.NewBed(id, name, dim))
	return nil
}

func (f *Greenhouse) containsBed(id types.BedId) bool {
	for _, b := range f.Beds {
		if b.Id == id {
			return true
		}
	}
	return false
}

//func RehydrateGreenhouse(
//	id types.GreenhouseId,
//	name string,
//	dim valueobject.Dimension,
//	beds []*bed.Bed,
//) *Greenhouse {
//	return &Greenhouse{
//		id:        id,
//		name:      name,
//		dimension: dim,
//		beds:      beds,
//	}
//}
