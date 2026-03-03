package cropplan

import (
	"errors"
	"time"
)

type StageType string

const (
	StageSoilPreparation StageType = "soil_preparation"
	StageSowing          StageType = "sowing"
	StageFertilization   StageType = "fertilization"
	StageProtection      StageType = "protection"
	StageIrrigation      StageType = "irrigation"
	StageHarvest         StageType = "harvest"
)

type StageStatus string

const (
	StageStatusPending    StageStatus = "pending"
	StageStatusInProgress StageStatus = "in_progress"
	StageStatusCompleted  StageStatus = "completed"
	StageStatusSkipped    StageStatus = "skipped"
)

// Stage представляет этап выращивания в плане
type Stage struct {
	ID          string
	PlanID      string
	Type        StageType
	Name        string
	Description string
	Status      StageStatus
	StartDate   *time.Time
	EndDate     *time.Time
	Order       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Metadata    map[string]interface{}
}

// NewStage создает новый этап
func NewStage(id, planID string, stageType StageType, name string, order int) (*Stage, error) {
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
		return nil, errors.New("order must be non-negative")
	}

	now := time.Now()
	return &Stage{
		ID:        id,
		PlanID:    planID,
		Type:      stageType,
		Name:      name,
		Status:    StageStatusPending,
		Order:     order,
		CreatedAt: now,
		UpdatedAt: now,
		Metadata:  make(map[string]interface{}),
	}, nil
}

// Start начинает выполнение этапа
func (s *Stage) Start() error {
	if s.Status != StageStatusPending {
		return errors.New("only pending stages can be started")
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
		return errors.New("only in-progress stages can be completed")
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

	now := time.Now()
	s.Status = StageStatusSkipped
	s.UpdatedAt = now

	return nil
}

// UpdateMetadata обновляет метаданные этапа
func (s *Stage) UpdateMetadata(key string, value interface{}) {
	s.Metadata[key] = value
	s.UpdatedAt = time.Now()
}

// IsFinished проверяет, завершен ли этап (выполнен или пропущен)
func (s *Stage) IsFinished() bool {
	return s.Status == StageStatusCompleted || s.Status == StageStatusSkipped
}
