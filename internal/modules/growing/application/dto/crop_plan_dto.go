package dto

import (
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"time"
)

// CropPlanDTO ответ для запросов
type CropPlanDTO struct {
	ID          string `json:"id"`
	BedID       string `json:"bed_id"`
	Name        string `json:"name"`
	VarietyID   string `json:"variety_id"`
	VarietyName string `json:"variety_name"`
	CropName    string `json:"crop_name"`
	Status      string `json:"status"`
	StatusText  string `json:"status_text"`

	SeasonStart  time.Time `json:"season_start"`
	SeasonEnd    time.Time `json:"season_end"`
	PlantingDate time.Time `json:"planting_date"`

	SeedsPlanted  int     `json:"seeds_planted"`
	ExpectedYield float64 `json:"expected_yield"`
	HarvestKg     float64 `json:"harvest_kg"`

	Stages []StageDTO `json:"stages"`

	Progress     float64   `json:"progress"` // % выполнения
	CurrentStage *StageDTO `json:"current_stage,omitempty"`

	AssignedTo   string `json:"assigned_to"`
	AssignedName string `json:"assigned_name"`

	CreatedAt   time.Time  `json:"created_at"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// StageDTO этап в ответе
type StageDTO struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	Description  string     `json:"description"`
	Status       string     `json:"status"`
	StatusText   string     `json:"status_text"`
	Order        int        `json:"order"`
	BBCHStart    int        `json:"bbch_start"`
	BBCHEnd      int        `json:"bbch_end"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	IsApplicable bool       `json:"is_applicable"` // можно ли выполнять сейчас
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
