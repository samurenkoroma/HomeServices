package croptype

import "errors"

// Общие ошибки
var (
	// Ошибки валидации
	ErrInvalidName           = errors.New("invalid crop type name")
	ErrInvalidScientificName = errors.New("invalid scientific name")
	ErrInvalidCategory       = errors.New("invalid crop category")
	ErrInvalidVegetationDays = errors.New("vegetation days must be greater than 0")
	ErrInvalidRootDepth      = errors.New("root depth must be greater than 0")
	ErrInvalidYield          = errors.New("default yield must be greater than 0")
	ErrInvalidPrice          = errors.New("market price must be greater than 0")

	// Ошибки существования
	ErrCropTypeNotFound      = errors.New("crop type not found")
	ErrCropTypeAlreadyExists = errors.New("crop type already exists")
	ErrCropTypeInUse         = errors.New("crop type is in use by plans or varieties")

	// Ошибки статуса
	ErrCropTypeInactive = errors.New("crop type is inactive")
	ErrCropTypeArchived = errors.New("crop type is archived")

	// Ошибки операций
	ErrCannotDeleteInUse    = errors.New("cannot delete crop type that has associated varieties or plans")
	ErrCannotModifyArchived = errors.New("cannot modify archived crop type")
)

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
