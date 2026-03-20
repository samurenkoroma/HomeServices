package cropplan

// RotationRule - правило севооборота
type RotationRule struct {
	PredecessorCropTypeID string `json:"predecessor_crop_type_id"` // Предшественник
	MinYears              int    `json:"min_years"`                // Минимальный перерыв в годах
	Recommended           bool   `json:"recommended"`              // Рекомендуемый или допустимый
	Notes                 string `json:"notes"`
}

func NewRotationRule(predecessorID string, minYears int, recommended bool) (RotationRule, error) {
	if minYears <= 0 {
		return RotationRule{}, ErrInvalidMinYears
	}
	if predecessorID == "" {
		return RotationRule{}, ErrInvalidPredecessor
	}

	return RotationRule{
		PredecessorCropTypeID: predecessorID,
		MinYears:              minYears,
		Recommended:           recommended,
	}, nil
}

// IsValidForRotation проверяет, подходит ли предшественник
func (r RotationRule) IsValidForRotation(lastCropID string, yearsSince int) bool {
	if lastCropID != r.PredecessorCropTypeID {
		return true // Другой предшественник - ок
	}
	return yearsSince >= r.MinYears
}
