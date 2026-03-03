package domain

type GrowthStage struct {
	name        string
	order       int
	duration    int
	minTemp     float64
	maxTemp     float64
	waterPerDay float64
}

func NewGrowthStage(
	name string,
	order int,
	duration int,
	minTemp float64,
	maxTemp float64,
	waterPerDay float64,
) (GrowthStage, error) {

	if duration <= 0 {
		return GrowthStage{}, ErrInvalidDuration
	}

	return GrowthStage{
		name:        name,
		order:       order,
		duration:    duration,
		minTemp:     minTemp,
		maxTemp:     maxTemp,
		waterPerDay: waterPerDay,
	}, nil
}
