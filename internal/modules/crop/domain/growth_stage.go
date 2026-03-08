package domain

type GrowthStage struct {
	name        string
	order       int
	duration    int
	minTemp     float64
	maxTemp     float64
	waterPerDay float64
}

func (g GrowthStage) Name() string {
	return g.name
}

func (g GrowthStage) Order() int {
	return g.order
}

func (g GrowthStage) Duration() int {
	return g.duration
}

func (g GrowthStage) MinTemp() float64 {
	return g.minTemp
}

func (g GrowthStage) MaxTemp() float64 {
	return g.maxTemp
}

func (g GrowthStage) WaterPerDay() float64 {
	return g.waterPerDay
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
