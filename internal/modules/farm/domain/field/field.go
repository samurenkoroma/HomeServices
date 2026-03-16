package field

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain"
	bedDomain "samurenkoroma/services/internal/modules/farm/domain/bed"
	blockDomain "samurenkoroma/services/internal/modules/farm/domain/field_block"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
	"time"
)

type Field struct {
	aggregate.Entity[types.FieldId]

	Name   string
	Beds   []BedReference
	Blocks []BlockReference

	Geom      spatial.GeoJSON
	TotalArea types.Area
	valueobject.Additions
}

type BlockReference struct {
	ID   types.FieldBlockId
	Name string
	Area types.Area
	Geom spatial.GeoJSON
}

func newBlockReference(id types.FieldBlockId, name string, json spatial.GeoJSON) BlockReference {
	return BlockReference{
		ID:   id,
		Name: name,
		Geom: json,
	}
}

type BedReference struct {
	ID        types.BedId
	Name      string
	Area      types.Area
	Dimension valueobject.Dimension
}

func newBedReference(id types.BedId, name string, dim valueobject.Dimension) BedReference {
	return BedReference{
		ID:        id,
		Name:      name,
		Dimension: dim,
	}
}
func NewField(name string, geom spatial.GeoJSON) *Field {
	f := &Field{
		Entity:    aggregate.NewEntity(types.FieldId(types.NewUUID())),
		Name:      name,
		Geom:      geom,
		Additions: valueobject.DefaultAdditions(),
		Blocks:    []BlockReference{},
		Beds:      []BedReference{},
	}

	f.AddEvent(NewFieldCreated(string(f.Id)))

	return f
}

func (f *Field) AddBlock(blockID types.FieldBlockId, blockNumber int, area types.Area) error {
	// Бизнес-правило: сумма площадей блоков не может превышать площадь поля
	totalBlocksArea := types.Area(0)
	for _, b := range f.Blocks {
		totalBlocksArea += b.Area
	}

	if totalBlocksArea+area > f.TotalArea {
		return domain.ErrAreaExceeded
	}

	f.Blocks = append(f.Blocks, BlockReference{
		ID:   blockID,
		Area: area,
	})

	f.UpdatedAt = time.Now()
	f.AddEvent(BlockAddedToField{
		FieldID:     f.Id,
		BlockID:     blockID,
		BlockNumber: blockNumber,
		Area:        area,
	})

	return nil
}

func (f *Field) AddBed(id types.BedId, name string, dim valueobject.Dimension) error {
	//if f.containsBed(id) {
	//	return domain.ErrDuplicateArea
	//}

	f.Beds = append(f.Beds, BedReference{ID: id, Name: name, Dimension: dim})
	return nil
}

func (f *Field) RehydrateAddBlock(b *blockDomain.FieldBlock) {
	f.Blocks = append(f.Blocks, newBlockReference(b.Id, b.Name, b.Geometry))
}

func (f *Field) RehydrateAddBed(bed *bedDomain.Bed) {
	f.Beds = append(f.Beds, newBedReference(bed.Id, bed.Name, bed.Dimension))
}

//func RehydrateGrowingFacility(
//	id types.GrowingAreaID,
//	name string,
//	dim valueobject.Dimension,
//	blocks []*blockDomain.FieldBlock,
//	beds []*bedDomain.Bed,
//) *Field {
//	blockRef =
//	return &Field{
//		id:        id,
//		name:      name,
//		dimension: dim,
//		blocks:    blocks,
//		beds:      beds,
//	}
//}
