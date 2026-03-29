package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/spatial"
	"time"
)

// AreaType — тип места выращивания
type AreaType string

const (
	AreaTypeField      AreaType = "field"
	AreaTypeBlock      AreaType = "block"
	AreaTypeGreenhouse AreaType = "greenhouse"
	AreaTypeBed        AreaType = "bed"
)

// SeasonConfig — конфигурация места на сезон
type SeasonConfig struct {
	SeasonID   string                 `json:"season_id"`
	Name       string                 `json:"name"`
	Geometry   spatial.GeoJSON        `json:"geometry"`
	Area       float64                `json:"area"`
	CropPlanID *string                `json:"crop_plan_id,omitempty"`
	BlockIDs   []string               `json:"block_ids,omitempty"` // для полей с поликультурой
	Metadata   map[string]interface{} `json:"metadata"`
	ValidFrom  time.Time              `json:"valid_from"`
	ValidUntil *time.Time             `json:"valid_until,omitempty"`
}

// AreaConfig — DTO для настройки места
type AreaConfig struct {
	Name       string                 `json:"name"`
	Geometry   spatial.GeoJSON        `json:"geometry"`
	CropPlanID *string                `json:"crop_plan_id,omitempty"`
	Metadata   map[string]interface{} `json:"metadata"`
}

// AreaSnapshot — снимок состояния для истории
type AreaSnapshot struct {
	SeasonID  string       `json:"season_id"`
	Config    SeasonConfig `json:"config"`
	ChangedAt time.Time    `json:"changed_at"`
	Reason    string       `json:"reason"`
}

// CultivationArea — интерфейс для всех мест выращивания
type CultivationArea interface {
	aggregate.Aggregate
	// Базовые методы
	GetID() string
	GetFarmRefID() string
	GetType() AreaType
	GetName() string
	GetGeometry() spatial.GeoJSON
	GetArea() float64

	// Работа с сезонами
	GetCurrentSeasonID() string
	GetSeasonConfig(seasonID string) (*SeasonConfig, error)
	IsConfiguredForSeason(seasonID string) bool
	GetHistory() []AreaSnapshot

	// Конфигурация
	ConfigureForSeason(seasonID string, config AreaConfig) error
	GetCropPlanForSeason(seasonID string) (string, error)
	HasBlocks() bool
	GetBlocks() []string
}

// Для Block добавляем специфичные методы (опционально)
type BlockInterface interface {
	CultivationArea
	AddBed(seasonID string, bedID string) error
}

// Для Bed специфичные методы
type BedInterface interface {
	CultivationArea
}

// Для GreenhouseArea специфичные методы
type GreenhouseAreaInterface interface {
	CultivationArea
	AddBed(seasonID string, bedID string) error
	GetLength() float64
	GetWidth() float64
}
