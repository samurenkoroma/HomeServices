package types

// FloatRange — общий тип для числовых диапазонов
type FloatRange struct {
	Min     float64 `json:"min"`
	Max     float64 `json:"max"`
	Optimal float64 `json:"optimal,omitempty"`
}

func (r FloatRange) IsValid() bool {
	return r.Min > 0 && r.Max >= r.Min
}

func (r FloatRange) Contains(value float64) bool {
	return value >= r.Min && value <= r.Max
}

// FloatRangeDTO — DTO для числового диапазона
type FloatRangeDTO struct {
	Min     float64 `json:"min" validate:"required,gt=0"`
	Max     float64 `json:"max" validate:"required,gtefield=Min"`
	Optimal float64 `json:"optimal,omitempty"`
}

// IntRange — диапазон дней (целые числа)
type IntRange struct {
	Min     int `json:"min"`
	Max     int `json:"max"`
	Optimal int `json:"optimal,omitempty"`
}

func (r IntRange) IsValid() bool {
	return r.Min > 0 && r.Max >= r.Min
}

func (r IntRange) Contains(days int) bool {
	return days >= r.Min && days <= r.Max
}

// IntRangeDTO — DTO для диапазона дней
type IntRangeDTO struct {
	Min     int `json:"min" validate:"required,gt=0"`
	Max     int `json:"max" validate:"required,gtefield=Min"`
	Optimal int `json:"optimal,omitempty"`
}
