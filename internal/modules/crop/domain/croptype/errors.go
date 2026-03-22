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
	ErrVarietyAlreadyExists  = errors.New("variety already exists")
	ErrCropTypeInUse         = errors.New("crop type is in use by plans or varieties")

	// Ошибки статуса
	ErrCropTypeInactive = errors.New("crop type is inactive")
	ErrCropTypeArchived = errors.New("crop type is archived")

	// Ошибки операций
	ErrCannotDeleteInUse    = errors.New("cannot delete crop type that has associated varieties or plans")
	ErrCannotModifyArchived = errors.New("cannot modify archived crop type")
)
