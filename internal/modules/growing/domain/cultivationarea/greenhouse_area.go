package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"time"
)

// GreenhouseArea — теплица как место выращивания
type GreenhouseArea struct {
	aggregate.Entity[string]

	farmRefID string
	name      string
	geometry  spatial.GeoJSON
	area      float64
	length    float64
	width     float64
	// Конфигурации по сезонам
	seasons         map[string]SeasonConfig
	currentSeasonID string
	childBeds       []string // Грядки внутри теплицы
}

func (g *GreenhouseArea) GetLength() float64 {
	return g.length
}

func (g *GreenhouseArea) GetWidth() float64 {
	return g.width
}

// NewGreenhouseArea создаёт новую теплицу как место выращивания
func NewGreenhouseArea(farmRefID, name string, dim types.Dimension, geom spatial.GeoJSON) *GreenhouseArea {
	return &GreenhouseArea{
		Entity:    aggregate.NewEntity(types.NewUUID()),
		farmRefID: farmRefID,
		name:      name,
		geometry:  geom,
		length:    *dim.Length,
		width:     *dim.Width,
		area:      dim.AreaInHectares(),
		seasons:   make(map[string]SeasonConfig),
		childBeds: []string{},
	}
}

// AddBed — добавить грядку в теплицу
func (g *GreenhouseArea) AddBed(seasonID string, bedID string) error {
	config, exists := g.seasons[seasonID]
	if !exists {
		return ErrSeasonConfigNotFound
	}

	config.BlockIDs = append(config.BlockIDs, bedID)
	g.seasons[seasonID] = config
	g.childBeds = append(g.childBeds, bedID)
	g.Update()

	return nil
}

// GetCropPlanForSeason — теплица может иметь несколько культур через грядки
func (g *GreenhouseArea) GetCropPlanForSeason(seasonID string) (string, error) {
	return "", ErrGreenhouseHasMultipleCrops
}

// GetSeasonConfig — получить конфигурацию на сезон
func (g *GreenhouseArea) GetSeasonConfig(seasonID string) (*SeasonConfig, error) {
	if config, exists := g.seasons[seasonID]; exists {
		return &config, nil
	}
	return nil, ErrSeasonConfigNotFound
}

// ConfigureForSeason — реализация интерфейса
func (g *GreenhouseArea) ConfigureForSeason(seasonID string, config AreaConfig) error {
	//if config.CropPlanID == nil {
	//	return ErrCropPlanRequiredForBlock
	//}

	if _, exists := g.seasons[seasonID]; exists {
		return ErrSeasonAlreadyConfigured
	}

	seasonConfig := SeasonConfig{
		SeasonID:   seasonID,
		Name:       config.Name,
		Geometry:   config.Geometry,
		Area:       0,
		CropPlanID: config.CropPlanID,
		BlockIDs:   []string{},
		Metadata:   config.Metadata,
		ValidFrom:  time.Now(),
	}

	g.seasons[seasonID] = seasonConfig
	g.currentSeasonID = seasonID
	g.name = config.Name
	g.geometry = config.Geometry
	g.area = seasonConfig.Area
	g.Update()

	g.AddEvent(GreenhouseConfigured{
		GreenhouseID: g.Id,
		SeasonID:     seasonID,
	})

	return nil
}

func (g *GreenhouseArea) Rehydrate(id string, createdAt, updatedAt time.Time, seasons []SeasonConfig) {
	g.Entity = aggregate.Entity[string]{Id: id, CreatedAt: createdAt, UpdatedAt: updatedAt}
	for _, season := range seasons {
		g.seasons[season.SeasonID] = season
	}
}

// SetFarmRefID устанавливает ссылку на farm модуль
func (g *GreenhouseArea) SetFarmRefID(farmRefID string) {
	g.farmRefID = farmRefID
	g.Update()
}

// Геттеры
func (g *GreenhouseArea) GetID() string                        { return g.Id }
func (g *GreenhouseArea) GetFarmRefID() string                 { return g.farmRefID }
func (g *GreenhouseArea) GetType() AreaType                    { return AreaTypeGreenhouse }
func (g *GreenhouseArea) GetName() string                      { return g.name }
func (g *GreenhouseArea) GetGeometry() spatial.GeoJSON         { return g.geometry }
func (g *GreenhouseArea) GetArea() float64                     { return g.area }
func (g *GreenhouseArea) GetCurrentSeasonID() string           { return g.currentSeasonID }
func (g *GreenhouseArea) IsConfiguredForSeason(id string) bool { _, ok := g.seasons[id]; return ok }
func (g *GreenhouseArea) HasBlocks() bool                      { return len(g.childBeds) > 0 }
func (g *GreenhouseArea) GetBlocks() []string                  { return g.childBeds }
func (g *GreenhouseArea) GetHistory() []AreaSnapshot           { return []AreaSnapshot{} }
