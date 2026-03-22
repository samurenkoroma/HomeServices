package variety

import (
	"encoding/json"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"time"
)

type VarietyID string
type Attributes struct {
	// Характеристики сорта
	VegetationDays    string   // Вегетационный период
	YieldPotential    string   // Потенциальная урожайность кг/га
	DiseaseResistance []string // Устойчивость к болезням

	// Рекомендации
	RecommendedRegions []string
	PlantingDensity    int     // шт/га Норма посадки
	SeedRate           float64 // кг/га Норма высева семян

	// Общие
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func (a *Attributes) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Attributes) Unmarshal(data []byte) error {
	return json.Unmarshal(data, a)
}

// Variety - сорт культуры
type Variety struct {
	aggregate.Entity[VarietyID]
	CropTypeID  string
	Name        string
	Description string

	Attributes Attributes

	IsActive bool
}

func NewVariety(cropTypeID string, name string, vegetationDays string, yieldPotential string) (*Variety, error) {
	if cropTypeID == "" {
		return nil, ErrEmptyCropType
	}
	if !types.IsValidUUID(cropTypeID) {
		return nil, cropplan.ErrInvalidCropType
	}
	if name == "" {
		return nil, ErrInvalidVarietyName
	}
	//if vegetationDays <= 0 {
	//	return nil, ErrInvalidVegetationDays
	//}
	return &Variety{
		Entity:     aggregate.NewEntity(VarietyID(types.NewUUID())),
		CropTypeID: cropTypeID,
		Name:       name,
		Attributes: Attributes{
			VegetationDays: vegetationDays,
			YieldPotential: yieldPotential,
		},

		IsActive: true,
	}, nil
}

func (v *Variety) Deactivate() {
	v.IsActive = false
	v.UpdatedAt = time.Now()
}

func (v *Variety) GetID() VarietyID      { return v.Entity.Id }
func (v *Variety) GetName() string       { return v.Name }
func (v *Variety) GetCropTypeID() string { return v.CropTypeID }
