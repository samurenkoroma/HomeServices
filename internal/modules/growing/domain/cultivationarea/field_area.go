package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"time"
)

// FieldUsageType — тип использования поля
type FieldUsageType string

const (
	FieldUsageMonoculture FieldUsageType = "monoculture" // одна культура на всём поле
	FieldUsagePolyculture FieldUsageType = "polyculture" // несколько культур (с участками)
)

// FieldArea — поле как место выращивания
type FieldArea struct {
	aggregate.Entity[string]

	farmRefID string
	name      string
	geometry  spatial.GeoJSON
	totalArea float64
	fieldType FieldUsageType

	// Конфигурации по сезонам
	seasons         map[string]SeasonConfig
	currentSeasonID string
	childBlocks     []string // ID блоков в текущем сезоне
}

// NewFieldArea создаёт новое поле как место выращивания
func NewFieldArea(farmRefID, name string, geom spatial.GeoJSON) *FieldArea {
	return &FieldArea{
		Entity:      aggregate.NewEntity(types.NewUUID()),
		farmRefID:   farmRefID,
		name:        name,
		geometry:    geom,
		totalArea:   0,
		fieldType:   FieldUsageMonoculture,
		seasons:     make(map[string]SeasonConfig),
		childBlocks: []string{},
	}
}

// ConfigureAsMonoculture — настроить поле как монокультуру на сезон
func (f *FieldArea) ConfigureAsMonoculture(
	seasonID string,
	name string,
	geom spatial.GeoJSON,
	cropPlanID string,
	metadata map[string]interface{},
) error {
	if _, exists := f.seasons[seasonID]; exists {
		return ErrSeasonAlreadyConfigured
	}

	if len(f.childBlocks) > 0 && f.currentSeasonID == seasonID {
		return ErrFieldHasBlocks
	}

	config := SeasonConfig{
		SeasonID:   seasonID,
		Name:       name,
		Geometry:   geom,
		Area:       0,
		CropPlanID: &cropPlanID,
		BlockIDs:   []string{},
		Metadata:   metadata,
		ValidFrom:  time.Now(),
	}

	f.seasons[seasonID] = config
	f.fieldType = FieldUsageMonoculture
	f.currentSeasonID = seasonID
	f.Update()

	f.AddEvent(FieldConfiguredAsMonoculture{
		FieldID:    f.Id,
		SeasonID:   seasonID,
		CropPlanID: cropPlanID,
		Area:       config.Area,
	})

	return nil
}

// ConfigureAsPolyculture — настроить поле для разбивки на участки
func (f *FieldArea) ConfigureAsPolyculture(
	seasonID string,
	name string,
	geom spatial.GeoJSON,
	metadata map[string]interface{},
) error {
	if _, exists := f.seasons[seasonID]; exists {
		return ErrSeasonAlreadyConfigured
	}

	config := SeasonConfig{
		SeasonID:   seasonID,
		Name:       name,
		Geometry:   geom,
		Area:       0,
		CropPlanID: nil,
		BlockIDs:   []string{},
		Metadata:   metadata,
		ValidFrom:  time.Now(),
	}

	f.seasons[seasonID] = config
	f.fieldType = FieldUsagePolyculture
	f.currentSeasonID = seasonID
	f.Update()

	f.AddEvent(FieldConfiguredAsPolyculture{
		FieldID:  f.Id,
		SeasonID: seasonID,
	})

	return nil
}

// AddBlock — добавить участок к полю
func (f *FieldArea) AddBlock(seasonID string, blockID string) error {
	config, exists := f.seasons[seasonID]
	if !exists {
		return ErrSeasonConfigNotFound
	}

	if f.fieldType != FieldUsagePolyculture {
		return ErrFieldNotPolyculture
	}

	config.BlockIDs = append(config.BlockIDs, blockID)
	f.seasons[seasonID] = config
	f.childBlocks = append(f.childBlocks, blockID)
	f.Update()

	return nil
}

// GetCropPlanForSeason — получить план культуры для сезона (монокультура)
func (f *FieldArea) GetCropPlanForSeason(seasonID string) (string, error) {
	config, exists := f.seasons[seasonID]
	if !exists {
		return "", ErrSeasonConfigNotFound
	}

	if config.CropPlanID == nil {
		return "", ErrNotMonocultureField
	}

	return *config.CropPlanID, nil
}

// GetSeasonConfig — получить конфигурацию на сезон
func (f *FieldArea) GetSeasonConfig(seasonID string) (*SeasonConfig, error) {
	if config, exists := f.seasons[seasonID]; exists {
		return &config, nil
	}
	return nil, ErrSeasonConfigNotFound
}

// ConfigureForSeason — общий метод для настройки
func (f *FieldArea) ConfigureForSeason(seasonID string, config AreaConfig) error {
	if config.CropPlanID != nil {
		return f.ConfigureAsMonoculture(seasonID, config.Name, config.Geometry, *config.CropPlanID, config.Metadata)
	}
	return f.ConfigureAsPolyculture(seasonID, config.Name, config.Geometry, config.Metadata)
}

// Геттеры
func (f *FieldArea) GetID() string                        { return f.Id }
func (f *FieldArea) GetFarmRefID() string                 { return f.farmRefID }
func (f *FieldArea) GetType() AreaType                    { return AreaTypeField }
func (f *FieldArea) GetName() string                      { return f.name }
func (f *FieldArea) GetGeometry() spatial.GeoJSON         { return f.geometry }
func (f *FieldArea) GetArea() float64                     { return f.totalArea }
func (f *FieldArea) GetCurrentSeasonID() string           { return f.currentSeasonID }
func (f *FieldArea) IsConfiguredForSeason(id string) bool { _, ok := f.seasons[id]; return ok }
func (f *FieldArea) HasBlocks() bool                      { return len(f.childBlocks) > 0 }
func (f *FieldArea) GetBlocks() []string                  { return f.childBlocks }
func (f *FieldArea) GetHistory() []AreaSnapshot           { return []AreaSnapshot{} } // будет реализовано позже
