package query

import (
	"context"
	"encoding/json"
	"errors"
	"samurenkoroma/services/internal/modules/growing/application/dto"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
)

// GetCropPlanHandler запрос получения плана
type GetCropPlanHandler struct {
	PlanRepo cropplan.Repository
}

type GetCropPlanQuery struct {
	PlanID string `json:"plan_id"`
}

func DecodeGetCropPlan(data []byte) (any, error) {
	var q GetCropPlanQuery
	if err := json.Unmarshal(data, &q); err != nil {
		return nil, err
	}
	if q.PlanID == "" {
		return nil, errors.New("plan_id is required")
	}
	return q, nil
}

func (h *GetCropPlanHandler) Handle(ctx context.Context, query any) (any, error) {
	q, ok := query.(GetCropPlanQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	plan, err := h.PlanRepo.FindByID(ctx, q.PlanID)
	if err != nil {
		return nil, err
	}

	return toCropPlanDTO(plan), nil
}

func toCropPlanDTO(plan *cropplan.CropPlan) *dto.CropPlanDTO {
	// Расчет прогресса
	totalStages := len(plan.Stages())
	completedStages := 0
	for _, s := range plan.Stages() {
		if s.Status == cropplan.StageStatusCompleted {
			completedStages++
		}
	}
	progress := 0.0
	if totalStages > 0 {
		progress = float64(completedStages) / float64(totalStages) * 100
	}

	// Текущий этап
	var currentStageDTO *dto.StageDTO
	if currentStage := plan.GetNextStage(); currentStage != nil {
		stageDTO := dto.ToStageDTO(*currentStage, 0)
		currentStageDTO = &stageDTO
	}

	// Конвертируем этапы
	stages := make([]dto.StageDTO, len(plan.Stages()))
	for i, s := range plan.Stages() {
		stages[i] = dto.ToStageDTO(s, 0)
	}

	return &dto.CropPlanDTO{
		ID:            plan.ID(),
		BedID:         plan.BedID(),
		Name:          plan.Name(),
		VarietyID:     plan.VarietyID(),
		VarietyName:   plan.VarietyName(),
		CropName:      plan.CropName(),
		Status:        string(plan.Status()),
		StatusText:    getStatusText(plan.Status()),
		SeasonStart:   plan.SeasonStart(),
		SeasonEnd:     plan.SeasonEnd(),
		PlantingDate:  plan.PlantingDate(),
		SeedsPlanted:  plan.SeedsPlanted(),
		ExpectedYield: plan.ExpectedYield(),
		HarvestKg:     plan.HarvestKg(),
		Stages:        stages,
		Progress:      progress,
		CurrentStage:  currentStageDTO,
		AssignedTo:    plan.AssignedTo(),
		AssignedName:  plan.AssignedName(),
		CreatedAt:     plan.CreatedAt(),
		StartedAt:     plan.StartedAt(),
		CompletedAt:   plan.CompletedAt(),
	}
}

func getStatusText(status cropplan.Status) string {
	switch status {
	case cropplan.StatusDraft:
		return "Черновик"
	case cropplan.StatusActive:
		return "Активный"
	case cropplan.StatusCompleted:
		return "Завершен"
	case cropplan.StatusCancelled:
		return "Отменен"
	default:
		return "Неизвестно"
	}
}
