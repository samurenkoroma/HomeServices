package query

import (
	"context"
	"encoding/json"
	"errors"
	"samurenkoroma/services/internal/modules/growing/application/dto"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"time"
)

// ListCropPlansHandler запрос списка планов
type ListCropPlansHandler struct {
	PlanRepo cropplan.Repository
}

// ListCropPlansQuery параметры фильтрации
type ListCropPlansQuery struct {
	BedID      string     `json:"bed_id,omitempty"`
	VarietyID  string     `json:"variety_id,omitempty"`
	Status     string     `json:"status,omitempty"`
	AssignedTo string     `json:"assigned_to,omitempty"`
	DateFrom   *time.Time `json:"date_from,omitempty"`
	DateTo     *time.Time `json:"date_to,omitempty"`
	Limit      int        `json:"limit,omitempty"`
	Offset     int        `json:"offset,omitempty"`
}

// ListCropPlansResponse ответ со списком планов
type ListCropPlansResponse struct {
	Total int                `json:"total"`
	Plans []*dto.CropPlanDTO `json:"plans"`
}

// DecodeListCropPlans декодирует JSON в запрос
func DecodeListCropPlans(data []byte) (any, error) {
	var q ListCropPlansQuery
	if err := json.Unmarshal(data, &q); err != nil {
		return nil, err
	}
	return q, nil
}

// Handle выполняет запрос
func (h *ListCropPlansHandler) Handle(ctx context.Context, query any) (any, error) {
	q, ok := query.(ListCropPlansQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	// Собираем планы по разным критериям
	var allPlans []*cropplan.CropPlan
	var err error

	if q.BedID != "" {
		allPlans, err = h.PlanRepo.FindByBed(ctx, q.BedID)
	} else if q.VarietyID != "" {
		allPlans, err = h.PlanRepo.FindByVariety(ctx, q.VarietyID)
	} else if q.Status != "" {
		allPlans, err = h.PlanRepo.FindByStatus(ctx, cropplan.Status(q.Status))
	} else if q.AssignedTo != "" {
		allPlans, err = h.PlanRepo.FindByAssignedTo(ctx, q.AssignedTo)
	} else {
		// Если нет фильтров, возвращаем все планы (через FindByStatus с пустым?)
		// В реальном репозитории нужен метод FindAll
		allPlans, err = h.getAllPlans(ctx)
	}

	if err != nil {
		return nil, err
	}

	// Фильтрация по датам
	var filtered []*cropplan.CropPlan
	for _, plan := range allPlans {
		if q.DateFrom != nil && plan.CreatedAt().Before(*q.DateFrom) {
			continue
		}
		if q.DateTo != nil && plan.CreatedAt().After(*q.DateTo) {
			continue
		}
		filtered = append(filtered, plan)
	}

	// Пагинация
	total := len(filtered)
	start := q.Offset
	end := q.Offset + q.Limit

	if start > total {
		start = total
	}
	if end > total || q.Limit == 0 {
		end = total
	}

	paginated := filtered[start:end]

	// Конвертируем в DTO
	plansDTO := make([]*dto.CropPlanDTO, len(paginated))
	for i, plan := range paginated {
		plansDTO[i] = toCropPlanDTO(plan)
	}

	return &ListCropPlansResponse{
		Total: total,
		Plans: plansDTO,
	}, nil
}

// getAllPlans временный метод для получения всех планов
// В реальном репозитории должен быть метод FindAll
func (h *ListCropPlansHandler) getAllPlans(ctx context.Context) ([]*cropplan.CropPlan, error) {
	// Здесь нужна реализация в репозитории
	// Пока возвращаем пустой список
	return []*cropplan.CropPlan{}, nil
}
