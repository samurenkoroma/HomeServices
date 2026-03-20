package croptype

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"time"
)

type CropTypeID string
type CropCategory string

const (
	CategoryCereal     CropCategory = "cereal"     // Зерновые
	CategoryLegume     CropCategory = "legume"     // Бобовые
	CategoryOilseed    CropCategory = "oilseed"    // Масличные
	CategoryVegetable  CropCategory = "vegetable"  // Овощные
	CategoryFruit      CropCategory = "fruit"      // Плодовые
	CategoryIndustrial CropCategory = "industrial" // Технические
)

// CropType - тип сельскохозяйственной культуры
type CropType struct {
	aggregate.BaseAggregate

	ID             CropTypeID
	Name           string // Пшеница, Кукуруза
	ScientificName string // Triticum aestivum
	Category       CropCategory
	Description    string

	// Биологические характеристики
	RootDepth      int  // Глубина корней в см
	IsPerennial    bool // Многолетняя
	VegetationDays int  // Вегетационный период в днях

	// Экономические характеристики
	DefaultYield float64 // Средняя урожайность кг/га
	MarketPrice  float64 // Средняя рыночная цена

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCropType(
	name string,
	scientificName string,
	category CropCategory,
	vegetationDays int,
) (*CropType, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	if vegetationDays <= 0 {
		return nil, ErrInvalidVegetationDays
	}

	return &CropType{
		ID:             CropTypeID(generateID()),
		Name:           name,
		ScientificName: scientificName,
		Category:       category,
		VegetationDays: vegetationDays,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (c *CropType) Update(name string, scientificName string) {
	c.Name = name
	c.ScientificName = scientificName
	c.UpdatedAt = time.Now()
}

func (c *CropType) SetRootDepth(depth int) {
	c.RootDepth = depth
	c.UpdatedAt = time.Now()
}

func (c *CropType) GetID() CropTypeID { return c.ID }
func (c *CropType) GetName() string   { return c.Name }
