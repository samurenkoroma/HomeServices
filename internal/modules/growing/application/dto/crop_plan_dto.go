package dto

import (
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"time"
)

// CropPlanDTO ответ для запросов
type CropPlanDTO struct {
	ID          string `json:"id"`
	BedID       string `json:"bedId"`
	Name        string `json:"name"`
	VarietyID   string `json:"varietyId"`
	VarietyName string `json:"varietyName"`
	CropName    string `json:"cropName"`
	Status      string `json:"status"`
	StatusText  string `json:"statusText"`

	SeasonStart  time.Time `json:"seasonStart"`
	SeasonEnd    time.Time `json:"seasonEnd"`
	PlantingDate time.Time `json:"plantingDate"`

	SeedsPlanted  int     `json:"seedsPlanted"`
	ExpectedYield float64 `json:"expectedYield"`
	HarvestKg     float64 `json:"harvestKg"`

	Stages []StageDTO `json:"stages"`

	Progress     float64   `json:"progress"` // % выполнения
	CurrentStage *StageDTO `json:"currentStage,omitempty"`

	AssignedTo   string `json:"assignedTo"`
	AssignedName string `json:"assignedName"`

	CreatedAt   time.Time  `json:"createdAt"`
	StartedAt   *time.Time `json:"startedAt,omitempty"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
}

// StageDTO этап в ответе
type StageDTO struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	Description  string     `json:"description"`
	Status       string     `json:"status"`
	StatusText   string     `json:"statusText"`
	Order        int        `json:"order"`
	BBCHStart    int        `json:"bbchStart"`
	BBCHEnd      int        `json:"bbchEnd"`
	StartDate    *time.Time `json:"startDate,omitempty"`
	EndDate      *time.Time `json:"endDate,omitempty"`
	IsApplicable bool       `json:"IsApplicable"` // можно ли выполнять сейчас
}

// ToStageDTO конвертирует доменный Stage в DTO
func ToStageDTO(stage cropplan.Stage, currentBBCH int) StageDTO {
	return StageDTO{
		ID:          stage.ID,
		Name:        stage.Name,
		Type:        string(stage.Type),
		Description: stage.Description,
		Status:      string(stage.Status),
		StatusText:  getStageStatusText(stage.Status),
		Order:       stage.Order,
		BBCHStart:   stage.BBCHStart,
		BBCHEnd:     stage.BBCHEnd,
		StartDate:   stage.StartDate,
		EndDate:     stage.EndDate,
		IsApplicable: stage.Status == cropplan.StageStatusPending &&
			currentBBCH >= stage.BBCHStart &&
			currentBBCH <= stage.BBCHEnd,
	}
}

func getStageStatusText(status cropplan.StageStatus) string {
	switch status {
	case cropplan.StageStatusPending:
		return "Ожидает"
	case cropplan.StageStatusInProgress:
		return "В процессе"
	case cropplan.StageStatusCompleted:
		return "Выполнен"
	case cropplan.StageStatusSkipped:
		return "Пропущен"
	default:
		return "Неизвестно"
	}
}
