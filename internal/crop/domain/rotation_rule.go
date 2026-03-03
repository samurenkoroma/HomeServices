package domain

type CropRotationRule struct {
	predecessor CropTypeID
	minYears    int
}

func NewCropRotationRule(
	predecessor CropTypeID,
	minYears int,
) (CropRotationRule, error) {

	if minYears <= 0 {
		return CropRotationRule{}, ErrInvalidDuration
	}

	return CropRotationRule{
		predecessor: predecessor,
		minYears:    minYears,
	}, nil
}
