package croptype

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"time"
)

type CropTypeID string

// CropType - тип сельскохозяйственной культуры
type CropType struct {
	aggregate.Entity[CropTypeID]
	id             CropTypeID
	name           string
	scientificName string
	category       CropCategory
	description    string

	// Биологические характеристики
	rootDepth      int
	isPerennial    bool
	vegetationDays int

	// Экономические характеристики
	defaultYield float64
	marketPrice  float64

	isActive  bool
	createdAt time.Time
	updatedAt time.Time
}

// NewCropType создает новый тип культуры
func NewCropType(
	name string,
	scientificName string,
	category CropCategory,
	vegetationDays int,
) (*CropType, error) {
	if name == "" {
		return nil, NewValidationError("name", "name is required", ErrInvalidName)
	}
	if vegetationDays <= 0 {
		return nil, NewValidationError("vegetation_days", "must be greater than 0", ErrInvalidVegetationDays)
	}
	if !category.IsValid() {
		return nil, NewValidationError("category", "invalid category", ErrInvalidCategory)
	}

	return &CropType{
		id:             CropTypeID(types.NewUUID()),
		name:           name,
		scientificName: scientificName,
		category:       category,
		vegetationDays: vegetationDays,
		isActive:       true,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
	}, nil
}

// Getters
func (c *CropType) GetID() CropTypeID         { return c.id }
func (c *CropType) GetName() string           { return c.name }
func (c *CropType) GetScientificName() string { return c.scientificName }
func (c *CropType) GetCategory() CropCategory { return c.category }
func (c *CropType) GetDescription() string    { return c.description }
func (c *CropType) GetRootDepth() int         { return c.rootDepth }
func (c *CropType) IsPerennial() bool         { return c.isPerennial }
func (c *CropType) GetVegetationDays() int    { return c.vegetationDays }
func (c *CropType) GetDefaultYield() float64  { return c.defaultYield }
func (c *CropType) GetMarketPrice() float64   { return c.marketPrice }
func (c *CropType) IsActive() bool            { return c.isActive }
func (c *CropType) GetCreatedAt() time.Time   { return c.createdAt }
func (c *CropType) GetUpdatedAt() time.Time   { return c.updatedAt }

// Setters with validation
func (c *CropType) SetDescription(desc string) {
	c.description = desc
	c.updatedAt = time.Now()
}

func (c *CropType) SetRootDepth(depth int) {
	if depth > 0 {
		c.rootDepth = depth
		c.updatedAt = time.Now()
	}
}

func (c *CropType) SetDefaultYield(yield float64) {
	if yield >= 0 {
		c.defaultYield = yield
		c.updatedAt = time.Now()
	}
}

func (c *CropType) SetMarketPrice(price float64) {
	if price >= 0 {
		c.marketPrice = price
		c.updatedAt = time.Now()
	}
}

// Deactivate деактивирует тип культуры
func (c *CropType) Deactivate() error {
	if !c.isActive {
		return ForbiddenError("crop type is already inactive")
	}

	c.isActive = false
	c.updatedAt = time.Now()

	return nil
}

// Activate активирует тип культуры
func (c *CropType) Activate() error {
	if c.isActive {
		return ForbiddenError("crop type is already active")
	}

	c.isActive = true
	c.updatedAt = time.Now()

	return nil
}

// Rehydrate восстанавливает тип культуры из БД
func (c *CropType) Rehydrate(
	id CropTypeID,
	description string,
	rootDepth int,
	isPerennial bool,
	defaultYield float64,
	marketPrice float64,
	isActive bool,
	createdAt, updatedAt time.Time,
) {
	c.id = id
	c.description = description
	c.rootDepth = rootDepth
	c.isPerennial = isPerennial
	c.defaultYield = defaultYield
	c.marketPrice = marketPrice
	c.isActive = isActive
	c.createdAt = createdAt
	c.updatedAt = updatedAt
}
