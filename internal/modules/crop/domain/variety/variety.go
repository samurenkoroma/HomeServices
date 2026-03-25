package variety

import (
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"strconv"
	"strings"
)

type (
	VarietyID  string
	Attributes struct {
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
)

type MinMax struct {
	Min int
	Max int
}

func (m *MinMax) String() string {
	return fmt.Sprintf("%d..%d", m.Min, m.Max)
}

func parseRange(data string) MinMax {
	parse := strings.Split(data, "..")
	minValue, _ := strconv.Atoi(parse[0])
	maxValue, _ := strconv.Atoi(parse[1])
	return MinMax{
		Min: minValue,
		Max: maxValue,
	}
}

func (a *Attributes) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Attributes) Unmarshal(data []byte) error {
	return json.Unmarshal(data, a)
}
func (a *Attributes) VD() MinMax {
	return parseRange(a.VegetationDays)
}

func (a *Attributes) YP() MinMax {
	return parseRange(a.YieldPotential)
}

// Variety - сорт культуры
type Variety struct {
	aggregate.Entity[VarietyID]
	cropTypeID  string
	name        string
	description string

	attributes Attributes

	isActive bool
}

func (v *Variety) Deactivate() {
	v.isActive = false
	v.Update()
}

func (v *Variety) ID() VarietyID       { return v.Entity.Id }
func (v *Variety) CropTypeID() string  { return v.cropTypeID }
func (v *Variety) Name() string        { return v.name }
func (v *Variety) Description() string { return v.description }

func (v *Variety) Attributes() Attributes       { return v.attributes }
func (v *Variety) IsActive() bool               { return v.isActive }
func (v *Variety) YearReleased() int            { return 2026 }
func (v *Variety) Breeder() string              { return "Bearer" }
func (v *Variety) SeedRate() float64            { return v.attributes.SeedRate }
func (v *Variety) PlantingDensity() int         { return v.attributes.PlantingDensity }
func (v *Variety) RecommendedRegions() []string { return v.attributes.RecommendedRegions }

func (v *Variety) DiseaseResistance() []string { return v.attributes.DiseaseResistance }

func (v *Variety) VegetationDays() MinMax {
	return v.attributes.VD()
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
		cropTypeID: cropTypeID,
		name:       name,
		attributes: Attributes{
			VegetationDays: vegetationDays,
			YieldPotential: yieldPotential,
		},

		isActive: true,
	}, nil
}
