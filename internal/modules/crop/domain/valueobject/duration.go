package valueobject

import "time"

// Duration — длительность в днях
type Duration int

func NewDuration(days int) (Duration, error) {
	if days <= 0 {
		return 0, ErrInvalidDuration
	}
	return Duration(days), nil
}

func (d Duration) Days() int { return int(d) }

// ToPeriod преобразует длительность в период от заданной даты
func (d Duration) ToPeriod(from time.Time) Period {
	end := from.AddDate(0, 0, d.Days())
	period, _ := NewPeriod(from, end)
	return period
}

func (d Duration) Equal(total int) bool {
	return d.Days() == total
}
