package cropplan

import (
	"time"
)

type VarietyID string

// Variety - сорт культуры
type Variety struct {
	ID          VarietyID
	CropTypeID  string
	Name        string
	Description string

	// Характеристики сорта
	VegetationDays    int      // Вегетационный период
	YieldPotential    float64  // Потенциальная урожайность кг/га
	DiseaseResistance []string // Устойчивость к болезням

	// Рекомендации
	RecommendedRegions []string
	PlantingDensity    int     // шт/га
	SeedRate           float64 // кг/га

	// Производитель
	Breeder      string
	YearReleased int

	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewVariety(
	cropTypeID string,
	name string,
	vegetationDays int,
	yieldPotential float64,
) (*Variety, error) {
	if cropTypeID == "" {
		return nil, ErrInvalidCropType
	}
	if name == "" {
		return nil, ErrInvalidVarietyName
	}
	if vegetationDays <= 0 {
		return nil, ErrInvalidVegetationDays
	}

	return &Variety{
		ID:             VarietyID(generateID()),
		CropTypeID:     cropTypeID,
		Name:           name,
		VegetationDays: vegetationDays,
		YieldPotential: yieldPotential,
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (v *Variety) Deactivate() {
	v.IsActive = false
	v.UpdatedAt = time.Now()
}

func (v *Variety) GetID() VarietyID      { return v.ID }
func (v *Variety) GetName() string       { return v.Name }
func (v *Variety) GetCropTypeID() string { return v.CropTypeID }
