package croptemplate

import "errors"

var (
	// Общие ошибки
	ErrTemplateNotFound        = errors.New("crop template not found")
	ErrTemplateAlreadyExists   = errors.New("crop template already exists")
	ErrCannotModifyPublished   = errors.New("cannot modify published template")
	ErrAlreadyPublished        = errors.New("template already published")
	ErrAlreadyArchived         = errors.New("template already archived")
	ErrNoStages                = errors.New("template has no stages")
	ErrStageNotFound           = errors.New("stage not found")
	ErrStageOrderDuplicate     = errors.New("duplicate stage order")
	ErrInvalidTemplateData     = errors.New("invalid template data")
	ErrTemplateVersionConflict = errors.New("template version conflict")
)
