package croptype

import (
	"samurenkoroma/services/internal/core/domain"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"time"
)

type CropTypeID string

// CropType - тип сельскохозяйственной культуры
type CropType struct {
	aggregate.Entity[CropTypeID]
	name        string
	category    CropCategory
	description string
	isPerennial bool
}

// NewCropType создает новый тип культуры
func NewCropType(
	name string,
	category CropCategory,
	description string,
	isPerennial bool,
) (*CropType, error) {
	if name == "" {
		return nil, domain.NewValidationError("name", "name is required", ErrInvalidName)
	}

	if !category.IsValid() {
		return nil, domain.NewValidationError("category", "invalid category", ErrInvalidCategory)
	}

	return &CropType{
		Entity:      aggregate.NewEntity(CropTypeID(types.NewUUID())),
		name:        name,
		category:    category,
		description: description,
		isPerennial: isPerennial,
	}, nil
}

// Getters
func (c *CropType) GetID() CropTypeID         { return c.Entity.Id }
func (c *CropType) GetName() string           { return c.name }
func (c *CropType) GetCategory() CropCategory { return c.category }
func (c *CropType) GetDescription() string    { return c.description }
func (c *CropType) IsPerennial() bool         { return c.isPerennial }
func (c *CropType) IsActive() bool            { return c.Entity.IsActive }
func (c *CropType) GetCreatedAt() time.Time   { return c.Entity.CreatedAt }
func (c *CropType) GetUpdatedAt() time.Time   { return c.Entity.UpdatedAt }

// Rehydrate восстанавливает тип культуры из БД
func Rehydrate(id, name, category, description string, isPerennial, isActive bool, createdAt, updatedAt time.Time) *CropType {
	return &CropType{
		Entity: aggregate.Entity[CropTypeID]{
			Id:        CropTypeID(id),
			IsActive:  isActive,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		name:        name,
		category:    CropCategory(category),
		isPerennial: isPerennial,
		description: description,
	}
}
