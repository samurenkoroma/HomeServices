package valueobject

import (
	"fmt"
	"strconv"
	"strings"
)

type MinMax struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

func (m *MinMax) String() string {
	return fmt.Sprintf("%d..%d", m.Min, m.Max)
}

func ParseRange(data string) MinMax {
	parse := strings.Split(data, "..")
	minValue, _ := strconv.Atoi(parse[0])
	maxValue, _ := strconv.Atoi(parse[1])
	return MinMax{
		Min: minValue,
		Max: maxValue,
	}
}
