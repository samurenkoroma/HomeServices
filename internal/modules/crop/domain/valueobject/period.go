package valueobject

import (
	"encoding/json"
	"time"
)

// Period — период времени
type Period struct {
	start time.Time
	end   time.Time
}

// NewPeriod создает новый период
func NewPeriod(start, end time.Time) (Period, error) {
	if start.After(end) {
		return Period{}, ErrInvalidPeriod
	}
	return Period{start: start, end: end}, nil
}

func (p Period) Start() time.Time { return p.start }
func (p Period) End() time.Time   { return p.end }

// Duration возвращает длительность периода в днях
func (p Period) Duration() int {
	return int(p.end.Sub(p.start).Hours() / 24)
}

// Contains проверяет, содержит ли период указанную дату
func (p Period) Contains(t time.Time) bool {
	return !t.Before(p.start) && !t.After(p.end)
}

// Overlaps проверяет пересечение с другим периодом
func (p Period) Overlaps(other Period) bool {
	return !p.end.Before(other.start) && !p.start.After(other.end)
}

// MarshalJSON реализует json.Marshaler
func (p Period) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}{
		Start: p.start.Format(time.RFC3339),
		End:   p.end.Format(time.RFC3339),
	})
}

// UnmarshalJSON реализует json.Unmarshaler
func (p *Period) UnmarshalJSON(data []byte) error {
	var aux struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	start, err := time.Parse(time.RFC3339, aux.Start)
	if err != nil {
		return err
	}
	end, err := time.Parse(time.RFC3339, aux.End)
	if err != nil {
		return err
	}

	period, err := NewPeriod(start, end)
	if err != nil {
		return err
	}
	*p = period
	return nil
}
