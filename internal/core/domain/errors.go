package domain

// ValidationError — ошибка валидации с дополнительным контекстом
type ValidationError struct {
	Field   string
	Message string
	Err     error
}

func (e ValidationError) Error() string {
	if e.Err != nil {
		return e.Field + ": " + e.Message + " (" + e.Err.Error() + ")"
	}
	return e.Field + ": " + e.Message
}

func (e ValidationError) Unwrap() error {
	return e.Err
}

// NewValidationError создает новую ошибку валидации
func NewValidationError(field, message string, err error) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
		Err:     err,
	}
}

// DomainError — доменная ошибка с типом
type DomainError struct {
	Type    string
	Message string
	Err     error
}

func (e DomainError) Error() string {
	if e.Err != nil {
		return e.Type + ": " + e.Message + " (" + e.Err.Error() + ")"
	}
	return e.Type + ": " + e.Message
}

func (e DomainError) Unwrap() error {
	return e.Err
}

// NewDomainError создает новую доменную ошибку
func NewDomainError(errType, message string, err error) DomainError {
	return DomainError{
		Type:    errType,
		Message: message,
		Err:     err,
	}
}

// Предопределенные доменные ошибки
var (
	// BusinessRuleError — нарушение бизнес-правила
	BusinessRuleError = func(message string) DomainError {
		return DomainError{
			Type:    "BUSINESS_RULE",
			Message: message,
		}
	}

	// ConflictError — конфликт данных
	ConflictError = func(message string) DomainError {
		return DomainError{
			Type:    "CONFLICT",
			Message: message,
		}
	}

	// ForbiddenError — операция запрещена
	ForbiddenError = func(message string) DomainError {
		return DomainError{
			Type:    "FORBIDDEN",
			Message: message,
		}
	}
)
