package cropplan

import (
	"errors"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"time"
)

// StageType — используем тип из catalog
type StageType = catalog.StageType

// Константы типов этапов (для удобства)
const (
	StageSoilPreparation = catalog.StageSoilPreparation
	StageSowing          = catalog.StageSowing
	StageFertilization   = catalog.StageFertilization
	StageProtection      = catalog.StageProtection
	StageIrrigation      = catalog.StageIrrigation
	StagePruning         = catalog.StagePruning
	StageHarvest         = catalog.StageHarvest
)

// StageStatus статус выполнения этапа
type StageStatus string

const (
	StageStatusPending    StageStatus = "pending"
	StageStatusInProgress StageStatus = "in_progress"
	StageStatusCompleted  StageStatus = "completed"
	StageStatusSkipped    StageStatus = "skipped"
)

// Stage этап выращивания (привязан к BBCH)
type Stage struct {
	// Идентификация
	ID     string `json:"id"`
	PlanID string `json:"plan_id"`

	// Тип и название
	Type        StageType `json:"type"`
	Name        string    `json:"name"`
	Description string    `json:"description"`

	// BBCH привязка (когда можно выполнять)
	BBCHStart int `json:"bbch_start"`
	BBCHEnd   int `json:"bbch_end"`

	// Статус
	Status StageStatus `json:"status"`
	Order  int         `json:"order"`

	// Даты
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// Дополнительно
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewStage создает новый этап
func NewStage(
	id, planID string,
	stageType StageType,
	name string,
	order int,
	bbchStart, bbchEnd int,
) (*Stage, error) {

	if id == "" {
		return nil, errors.New("stage id is required")
	}
	if planID == "" {
		return nil, errors.New("plan id is required")
	}
	if name == "" {
		return nil, errors.New("stage name is required")
	}
	if order < 0 {
		return nil, ErrInvalidStageOrder
	}
	if bbchStart < 0 || bbchEnd < 0 || bbchStart > bbchEnd {
		return nil, errors.New("invalid BBCH range")
	}

	now := time.Now()
	return &Stage{
		ID:        id,
		PlanID:    planID,
		Type:      stageType,
		Name:      name,
		Status:    StageStatusPending,
		Order:     order,
		BBCHStart: bbchStart,
		BBCHEnd:   bbchEnd,
		CreatedAt: now,
		UpdatedAt: now,
		Metadata:  make(map[string]interface{}),
	}, nil
}

// NewStageFromTemplate создает этап из шаблона catalog.StageTemplate
func NewStageFromTemplate(
	id, planID string,
	template catalog.StageTemplate,
	order int,
) (*Stage, error) {
	return NewStage(
		id, planID,
		template.Type,
		template.Name,
		order,
		template.BBCHStart,
		template.BBCHEnd,
	)
}

// IsApplicableForBBCH проверяет, можно ли выполнять этап при данном BBCH
func (s *Stage) IsApplicableForBBCH(bbchCode int) bool {
	return bbchCode >= s.BBCHStart && bbchCode <= s.BBCHEnd
}

// CanStart проверяет, можно ли начать этап
func (s *Stage) CanStart(currentBBCH int) error {
	if s.Status != StageStatusPending {
		return ErrStageAlreadyStarted
	}
	if !s.IsApplicableForBBCH(currentBBCH) {
		return errors.New("stage not applicable for current BBCH phase")
	}
	return nil
}

// Start начинает этап
func (s *Stage) Start(currentBBCH int) error {
	if err := s.CanStart(currentBBCH); err != nil {
		return err
	}

	now := time.Now()
	s.Status = StageStatusInProgress
	s.StartDate = &now
	s.UpdatedAt = now

	return nil
}

// Complete завершает этап
func (s *Stage) Complete() error {
	if s.Status != StageStatusInProgress {
		return ErrStageNotInProgress
	}

	now := time.Now()
	s.Status = StageStatusCompleted
	s.EndDate = &now
	s.UpdatedAt = now

	return nil
}

// Skip пропускает этап
func (s *Stage) Skip() error {
	if s.Status != StageStatusPending {
		return errors.New("only pending stages can be skipped")
	}

	s.Status = StageStatusSkipped
	s.UpdatedAt = time.Now()

	return nil
}

// IsFinished проверяет, завершен ли этап
func (s *Stage) IsFinished() bool {
	return s.Status == StageStatusCompleted || s.Status == StageStatusSkipped
}
