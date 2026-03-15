package field_block

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain"
	domain2 "samurenkoroma/services/internal/modules/farm/domain/bed"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
	"time"
)

type FieldBlock struct {
	aggregate.Entity[types.FieldBlockId]
	Name      string
	Dimension valueobject.Dimension

	Beds     []*domain2.Bed
	Geometry spatial.GeoJSON
	ParentId types.FieldId
	valueobject.Additions
}

func NewFieldBlock(fieldId types.FieldId, name string, dim valueobject.Dimension, geom spatial.GeoJSON) *FieldBlock {
	b := &FieldBlock{
		Entity:    aggregate.NewEntity(types.FieldBlockId(types.NewUUID())),
		Name:      name,
		Dimension: dim,
		Geometry:  geom,
		ParentId:  fieldId,
		Additions: valueobject.Additions{},
	}
	b.AddEvent(Created{})
	return b
}

func (b *FieldBlock) AssignToCropCycle(cropCycleID, cropID, varietyID string, plantedAt time.Time) error {
	if b.Status != valueobject.AreaStatusEmpty && b.Status != valueobject.AreaStatusFallow {
		return domain.ErrBlockNotAvailable
	}

	b.CurrentCrop = &valueobject.CurrentCropInfo{
		CropCycleID: cropCycleID,
		CropID:      cropID,
		VarietyID:   varietyID,
		PlantedAt:   plantedAt,
	}

	b.Status = valueobject.AreaStatusPlanted
	b.UpdatedAt = time.Now()

	b.AddEvent(BlockAssignedToCrop{
		BlockID:     b.Id,
		CropCycleID: cropCycleID,
		CropID:      cropID,
		VarietyID:   varietyID,
		PlantedAt:   plantedAt,
	})

	return nil
}

func (b *FieldBlock) AddBed(bed *domain2.Bed) {
	b.Beds = append(b.Beds, bed)
}
func (b *FieldBlock) ContainsBed(id types.BedId) bool {
	for _, bed := range b.Beds {
		if bed.Id == id {
			return true
		}
	}
	return false
}
func (b *FieldBlock) RehydrateAddBed(bed *domain2.Bed) {
	b.Beds = append(b.Beds, bed)
}

//func RehydrateFieldBlock(id types.GrowingAreaID, name string, dim valueobject.Dimension) FieldBlock {
//	return FieldBlock{
//		Entity:    aggregate.NewEntity(id),
//		Name:      name,
//		Dimension: dim,
//	}
//}
