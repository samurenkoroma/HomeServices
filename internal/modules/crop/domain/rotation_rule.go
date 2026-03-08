package domain

type CropRotationRule struct {
	predecessor CropTypeID
	minYears    int
}

func (c CropRotationRule) Predecessor() CropTypeID {
	return c.predecessor
}

func (c CropRotationRule) MinYears() int {
	return c.minYears
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
