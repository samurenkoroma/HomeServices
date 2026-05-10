package cultivationarea

import (
	"context"
	"time"
)

// AreaFilter — фильтр для списка мест
type AreaFilter struct {
	ObjectId string `json:"objectId"`
}

// AreaListItem — DTO для списка мест
type AreaListItem struct {
	ID       string  `json:"id"`
	ObjectId string  `json:"objectId"`
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Area     float64 `json:"area"`
}

// AreaDetail — DTO для детальной страницы места выращивания
type AreaDetail struct {
	// Основная информация
	ID        string  `json:"id"`
	FarmRefID string  `json:"farm_ref_id"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Geometry  any     `json:"geometry"`
	Area      float64 `json:"area"`
	ParentID  *string `json:"parent_id,omitempty"`

	// Атрибуты (для теплиц, полей)
	Attributes *AreaAttributes `json:"attributes,omitempty"`

	// Конфигурация на текущий сезон
	CurrentSeasonID   *string `json:"current_season_id,omitempty"`
	CurrentSeasonName *string `json:"current_season_name,omitempty"`
	CropPlanID        *string `json:"crop_plan_id,omitempty"`
	CropPlanName      *string `json:"crop_plan_name,omitempty"`
	UsageType         *string `json:"usage_type,omitempty"`

	// Статистика
	ActiveCyclesCount    int     `json:"active_cycles_count"`
	CompletedCyclesCount int     `json:"completed_cycles_count"`
	TotalYield           float64 `json:"total_yield"`

	// Метаданные
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AreaAttributes — атрибуты места выращивания (для теплиц)
type AreaAttributes struct {
	// Для теплицы
	GreenhouseType string  `json:"greenhouse_type,omitempty"` // film, glass, polycarbonate
	Width          float64 `json:"width,omitempty"`
	Length         float64 `json:"length,omitempty"`
	Height         float64 `json:"height,omitempty"`
	HasHeating     bool    `json:"has_heating,omitempty"`
	HasVentilation bool    `json:"has_ventilation,omitempty"`
	HasLighting    bool    `json:"has_lighting,omitempty"`

	// Для поля
	SoilType string `json:"soil_type,omitempty"`

	// Общие
	Description string `json:"description,omitempty"`
}

// Projection — интерфейс read-модели для мест выращивания
type Projections interface {
	GetList(ctx context.Context, filter AreaFilter) ([]AreaListItem, error)
	GetByID(ctx context.Context, id string) (*AreaDetail, error)
}
