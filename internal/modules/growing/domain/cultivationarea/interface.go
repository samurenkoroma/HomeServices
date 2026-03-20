package cultivationarea

import (
	"samurenkoroma/services/internal/core/spatial"
	"time"
)

type AreaType string

const (
	AreaTypeField      AreaType = "field"
	AreaTypeBlock      AreaType = "block"
	AreaTypeGreenhouse AreaType = "greenhouse"
	AreaTypeBed        AreaType = "bed"
)

// CultivationArea - интерфейс для всех мест выращивания
type CultivationArea interface {
	// Базовые методы
	GetID() string
	GetFarmRefID() string // Ссылка на объект в farm модуле
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
	GetCropPlanForSeason(seasonID string) (string, error) // Для монокультуры
	HasBlocks() bool
	GetBlocks() []string
}

// SeasonConfig - конфигурация на сезон
type SeasonConfig struct {
	SeasonID   string
	Name       string
	Geometry   spatial.GeoJSON
	Area       float64
	CropPlanID *string  // Для монокультуры
	BlockIDs   []string // Для поликультуры (ссылки на блоки)
	Metadata   map[string]interface{}
	ValidFrom  time.Time
	ValidUntil *time.Time
}

// AreaConfig - DTO для настройки места
type AreaConfig struct {
	Name       string
	Geometry   spatial.GeoJSON
	CropPlanID *string
	Metadata   map[string]interface{}
}

// AreaSnapshot - снимок состояния для истории
type AreaSnapshot struct {
	SeasonID  string
	Config    SeasonConfig
	ChangedAt time.Time
	Reason    string
}
