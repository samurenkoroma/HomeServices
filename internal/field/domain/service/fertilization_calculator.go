package service

import (
	"samurenkoroma/services/internal/field/domain/landunit"
	"samurenkoroma/services/internal/field/domain/valueobject"
)

type FertilizationCalculator struct{}

func (f FertilizationCalculator) CalculateForBed(
	bed *landunit.Bed,
	norm valueobject.FertilizerNorm,
) float64 {
	return 4
}
