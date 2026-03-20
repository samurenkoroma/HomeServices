package valueobject

import (
	"encoding/json"
	"fmt"
)

// Area — площадь
type Area float64

func NewArea(value float64) (Area, error) {
	if value <= 0 {
		return 0, fmt.Errorf("area must be greater than 0: %w", ErrNegativeValue)
	}
	return Area(value), nil
}

func (a Area) Value() float64 { return float64(a) }

// InHectares возвращает площадь в гектарах
func (a Area) InHectares() float64 {
	return float64(a) / 10000
}

// InSquareMeters возвращает площадь в квадратных метрах
func (a Area) InSquareMeters() float64 {
	return float64(a)
}

// MarshalJSON реализует json.Marshaler
func (a Area) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(a))
}

// UnmarshalJSON реализует json.Unmarshaler
func (a *Area) UnmarshalJSON(data []byte) error {
	var value float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	area, err := NewArea(value)
	if err != nil {
		return err
	}
	*a = area
	return nil
}

// Density — плотность посадки (растений/га)
type Density int

func NewDensity(value int) (Density, error) {
	if value <= 0 {
		return 0, fmt.Errorf("density must be greater than 0: %w", ErrNegativeValue)
	}
	return Density(value), nil
}

func (d Density) Value() int { return int(d) }

// PlantsPerHectare возвращает количество растений на гектар
func (d Density) PlantsPerHectare() int {
	return int(d)
}

// PlantsForArea возвращает количество растений для заданной площади
func (d Density) PlantsForArea(area Area) int {
	return int(float64(d) * area.InHectares())
}

// Temperature — температура
type Temperature float64

func NewTemperature(value float64) Temperature {
	return Temperature(value)
}

func (t Temperature) Celsius() float64 { return float64(t) }
