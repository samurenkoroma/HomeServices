package domain

type EnvironmentalRequirements struct {
	minTemp     float64
	maxTemp     float64
	minHumidity float64
	maxHumidity float64
	minPH       float64
	maxPH       float64
}

func (e EnvironmentalRequirements) MinTemp() float64 {
	return e.minTemp
}

func (e EnvironmentalRequirements) MaxTemp() float64 {
	return e.maxTemp
}

func (e EnvironmentalRequirements) MinHumidity() float64 {
	return e.minHumidity
}

func (e EnvironmentalRequirements) MaxHumidity() float64 {
	return e.maxHumidity
}

func (e EnvironmentalRequirements) MinPH() float64 {
	return e.minPH
}

func (e EnvironmentalRequirements) MaxPH() float64 {
	return e.maxPH
}
