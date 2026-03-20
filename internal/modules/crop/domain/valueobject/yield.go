package valueobject

import (
	"encoding/json"
	"fmt"
	"time"
)

// YieldPotential — потенциал урожайности
type YieldPotential struct {
	min      float64 // кг/га
	max      float64 // кг/га
	expected float64 // кг/га
}

// NewYieldPotential создает новый объект потенциала урожайности
func NewYieldPotential(min, max, expected float64) (YieldPotential, error) {
	if min < 0 || max < 0 || expected < 0 {
		return YieldPotential{}, ErrNegativeYield
	}
	if min > max {
		return YieldPotential{}, ErrInvalidYieldRange
	}
	if expected < min || expected > max {
		return YieldPotential{}, ErrExpectedOutOfRange
	}

	return YieldPotential{
		min:      min,
		max:      max,
		expected: expected,
	}, nil
}

func (y YieldPotential) Min() float64      { return y.min }
func (y YieldPotential) Max() float64      { return y.max }
func (y YieldPotential) Expected() float64 { return y.expected }

// Range возвращает диапазон урожайности
func (y YieldPotential) Range() float64 {
	return y.max - y.min
}

// MarshalJSON реализует json.Marshaler
func (y YieldPotential) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Min      float64 `json:"min"`
		Max      float64 `json:"max"`
		Expected float64 `json:"expected"`
	}{
		Min:      y.min,
		Max:      y.max,
		Expected: y.expected,
	})
}

// UnmarshalJSON реализует json.Unmarshaler
func (y *YieldPotential) UnmarshalJSON(data []byte) error {
	var aux struct {
		Min      float64 `json:"min"`
		Max      float64 `json:"max"`
		Expected float64 `json:"expected"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	potential, err := NewYieldPotential(aux.Min, aux.Max, aux.Expected)
	if err != nil {
		return err
	}
	*y = potential
	return nil
}
