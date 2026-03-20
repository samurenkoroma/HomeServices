package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/spatial"
	"time"
)

// FieldArea - поле как место выращивания
type FieldArea struct {
	aggregate.BaseAggregate

	ID        string
	FarmRefID string
	Name      string
	Geometry  spatial.GeoJSON
	TotalArea float64
	FieldType FieldUsageType // monoculture или polyculture

	// Конфигурации по сезонам
	Seasons         map[string]SeasonConfig
	CurrentSeasonID string
	ChildBlocks     []string // ID блоков в текущем сезоне
}

type FieldUsageType string

const (
	FieldUsageMonoculture FieldUsageType = "monoculture"
	FieldUsagePolyculture FieldUsageType = "polyculture"
)

func NewFieldArea(farmRefID string, name string, geom spatial.GeoJSON) *FieldArea {
	return &FieldArea{
		ID:          generateID(),
		FarmRefID:   farmRefID,
		Name:        name,
		Geometry:    geom,
		TotalArea:   calculateArea(geom),
		FieldType:   FieldUsageMonoculture,
		Seasons:     make(map[string]SeasonConfig),
		ChildBlocks: []string{},
	}
}

// ConfigureAsMonoculture - настроить как монокультуру на сезон
func (f *FieldArea) ConfigureAsMonoculture(
	seasonID string,
	name string,
	geom spatial.GeoJSON,
	cropPlanID string,
	metadata map[string]interface{},
) error {
	if _, exists := f.Seasons[seasonID]; exists {
		return ErrSeasonAlreadyConfigured
	}

	if len(f.ChildBlocks) > 0 && f.CurrentSeasonID == seasonID {
		return ErrFieldHasBlocks
	}

	config := SeasonConfig{
		SeasonID:   seasonID,
		Name:       name,
		Geometry:   geom,
		Area:       calculateArea(geom),
		CropPlanID: &cropPlanID,
		BlockIDs:   []string{},
		Metadata:   metadata,
		ValidFrom:  time.Now(),
	}

	f.Seasons[seasonID] = config
	f.FieldType = FieldUsageMonoculture
	f.CurrentSeasonID = seasonID

	f.AddEvent(FieldConfiguredAsMonoculture{
		FieldID:    f.ID,
		SeasonID:   seasonID,
		CropPlanID: cropPlanID,
		Area:       config.Area,
	})

	return nil
}

// ConfigureAsPolyculture - настроить как поликультуру на сезон
func (f *FieldArea) ConfigureAsPolyculture(
	seasonID string,
	name string,
	geom spatial.GeoJSON,
	metadata map[string]interface{},
) error {
	if _, exists := f.Seasons[seasonID]; exists {
		return ErrSeasonAlreadyConfigured
	}

	config := SeasonConfig{
		SeasonID:   seasonID,
		Name:       name,
		Geometry:   geom,
		Area:       calculateArea(geom),
		CropPlanID: nil,
		BlockIDs:   []string{},
		Metadata:   metadata,
		ValidFrom:  time.Now(),
	}

	f.Seasons[seasonID] = config
	f.FieldType = FieldUsagePolyculture
	f.CurrentSeasonID = seasonID

	f.AddEvent(FieldConfiguredAsPolyculture{
		FieldID:  f.ID,
		SeasonID: seasonID,
	})

	return nil
}

// AddBlock - добавить участок к полю
func (f *FieldArea) AddBlock(seasonID string, blockID string) error {
	config, exists := f.Seasons[seasonID]
	if !exists {
		return ErrSeasonConfigNotFound
	}

	if f.FieldType != FieldUsagePolyculture {
		return ErrFieldNotPolyculture
	}

	config.BlockIDs = append(config.BlockIDs, blockID)
	f.Seasons[seasonID] = config
	f.ChildBlocks = append(f.ChildBlocks, blockID)

	return nil
}

// GetCropPlanForSeason - получить план культуры для сезона (монокультура)
func (f *FieldArea) GetCropPlanForSeason(seasonID string) (string, error) {
	config, exists := f.Seasons[seasonID]
	if !exists {
		return "", ErrSeasonConfigNotFound
	}

	if config.CropPlanID == nil {
		return "", ErrNotMonocultureField
	}

	return *config.CropPlanID, nil
}

func (f *FieldArea) GetID() string                        { return f.ID }
func (f *FieldArea) GetFarmRefID() string                 { return f.FarmRefID }
func (f *FieldArea) GetType() AreaType                    { return AreaTypeField }
func (f *FieldArea) GetName() string                      { return f.Name }
func (f *FieldArea) GetGeometry() spatial.GeoJSON         { return f.Geometry }
func (f *FieldArea) GetArea() float64                     { return f.TotalArea }
func (f *FieldArea) GetCurrentSeasonID() string           { return f.CurrentSeasonID }
func (f *FieldArea) IsConfiguredForSeason(id string) bool { _, ok := f.Seasons[id]; return ok }
func (f *FieldArea) HasBlocks() bool                      { return len(f.ChildBlocks) > 0 }
func (f *FieldArea) GetBlocks() []string                  { return f.ChildBlocks }

func (f *FieldArea) GetSeasonConfig(seasonID string) (*SeasonConfig, error) {
	if config, exists := f.Seasons[seasonID]; exists {
		return &config, nil
	}
	return nil, ErrSeasonConfigNotFound
}

func (f *FieldArea) ConfigureForSeason(seasonID string, config AreaConfig) error {
	if config.CropPlanID != nil {
		return f.ConfigureAsMonoculture(seasonID, config.Name, config.Geometry, *config.CropPlanID, config.Metadata)
	}
	return f.ConfigureAsPolyculture(seasonID, config.Name, config.Geometry, config.Metadata)
}
