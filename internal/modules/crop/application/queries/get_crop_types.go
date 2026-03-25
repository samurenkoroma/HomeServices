package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"
	"time"

	"github.com/olekukonko/errors"
)

// GetCropTypesQuery — параметры запроса списка типов культур
type GetCropTypesQuery struct {
	Category   string `json:"category"`
	ActiveOnly bool   `json:"active_only"`
}

// CropTypeDTO — DTO для ответа
type CropTypeDTO struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ScientificName string `json:"scientific_name"`
	Category       string `json:"category"`
	CategoryName   string `json:"category_name"`
	Description    string `json:"description"`
	IsPerennial    bool   `json:"is_perennial"`
	IsActive       bool   `json:"is_active"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// GetCropTypesHandler — обработчик запроса
type getCropTypesHandler struct {
	repo croptype.Repository
}

func (h *getCropTypesHandler) Name() string {
	return "GetCropTypes"
}

func NewGetCropTypesHandler(repo croptype.Repository) query.QueryHandler {
	return &getCropTypesHandler{repo: repo}
}

func (h *getCropTypesHandler) Handle(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetCropTypesQuery)
	if !ok {
		return nil, errors.New("invalid payload type")
	}
	// Получаем данные из репозитория
	var cropTypes []*croptype.CropType
	var err error

	if q.Category != "" {
		cropTypes, err = h.repo.FindByCategory(ctx, croptype.CropCategory(q.Category))
	} else {
		if q.ActiveOnly {
			cropTypes, err = h.repo.FindActive(ctx)
		} else {
			cropTypes, err = h.repo.FindAll(ctx)
		}
	}

	if err != nil {
		return nil, err
	}

	// Конвертируем в DTO
	result := make([]CropTypeDTO, len(cropTypes))
	for i, ct := range cropTypes {
		result[i] = CropTypeDTO{
			ID:             string(ct.GetID()),
			Name:           ct.GetName(),
			ScientificName: "latina",
			Category:       string(ct.GetCategory()),
			CategoryName:   ct.GetCategory().DisplayName(),
			Description:    ct.GetDescription(),
			IsPerennial:    ct.IsPerennial(),
			IsActive:       ct.IsActive(),
			CreatedAt:      ct.GetCreatedAt().Format(time.RFC3339),
			UpdatedAt:      ct.GetUpdatedAt().Format(time.RFC3339),
		}
	}

	return result, nil
}
