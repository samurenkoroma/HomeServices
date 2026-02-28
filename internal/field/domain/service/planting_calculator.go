package service

import (
	"samurenkoroma/services/internal/field/domain/landunit"
	"samurenkoroma/services/internal/field/domain/valueobject"
)

type PlantingCalculator struct{}

func (p PlantingCalculator) CalculateForBed(
	bed *landunit.Bed,
	scheme valueobject.PlantingScheme,
) int {
	return 2
}
