package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/spatial"
	"time"
)

// GreenhouseArea - теплица как место выращивания
type GreenhouseArea struct {
	aggregate.BaseAggregate

	ID        string
	FarmRefID string
	Name      string
	Geometry  spatial.GeoJSON
	Area      float64

	Seasons         map[string]SeasonConfig
	CurrentSeasonID string
	ChildBeds       []string // Грядки внутри теплицы
}

func NewGreenhouseArea(farmRefID string, name string, geom spatial.GeoJSON) *GreenhouseArea {
	return &GreenhouseArea{
		ID:        generateID(),
		FarmRefID: farmRefID,
		Name:      name,
		Geometry:  geom,
		Area:      calculateArea(geom),
		Seasons:   make(map[string]SeasonConfig),
		ChildBeds: []string{},
	}
}

// ConfigureForSeason - настроить теплицу на сезон (всегда монокультура или грядки)
func (g *GreenhouseArea) ConfigureForSeason(
	seasonID string,
	name string,
	geom spatial.GeoJSON,
	metadata map[string]interface{},
) error {
	if _, exists := g.Seasons[seasonID]; exists {
		return ErrSeasonAlreadyConfigured
	}

	config := SeasonConfig{
		SeasonID:   seasonID,
		Name:       name,
		Geometry:   geom,
		Area:       calculateArea(geom),
		CropPlanID: nil, // Теплица может иметь несколько культур на грядках
		BlockIDs:   []string{},
		Metadata:   metadata,
		ValidFrom:  time.Now(),
	}

	g.Seasons[seasonID] = config
	g.CurrentSeasonID = seasonID

	g.AddEvent(GreenhouseConfigured{
		GreenhouseID: g.ID,
		SeasonID:     seasonID,
	})

	return nil
}

// AddBed - добавить грядку в теплицу
func (g *GreenhouseArea) AddBed(seasonID string, bedID string) error {
	config, exists := g.Seasons[seasonID]
	if !exists {
		return ErrSeasonConfigNotFound
	}

	config.BlockIDs = append(config.BlockIDs, bedID)
	g.Seasons[seasonID] = config
	g.ChildBeds = append(g.ChildBeds, bedID)

	return nil
}

func (g *GreenhouseArea) GetID() string                { return g.ID }
func (g *GreenhouseArea) GetFarmRefID() string         { return g.FarmRefID }
func (g *GreenhouseArea) GetType() AreaType            { return AreaTypeGreenhouse }
func (g *GreenhouseArea) GetName() string              { return g.Name }
func (g *GreenhouseArea) GetGeometry() spatial.GeoJSON { return g.Geometry }
func (g *GreenhouseArea) GetArea() float64             { return g.Area }
func (g *GreenhouseArea) GetCurrentSeasonID() string   { return g.CurrentSeasonID }
func (g *GreenhouseArea) HasBlocks() bool              { return len(g.ChildBeds) > 0 }
func (g *GreenhouseArea) GetBlocks() []string          { return g.ChildBeds }

func (g *GreenhouseArea) GetSeasonConfig(seasonID string) (*SeasonConfig, error) {
	if config, exists := g.Seasons[seasonID]; exists {
		return &config, nil
	}
	return nil, ErrSeasonConfigNotFound
}

func (g *GreenhouseArea) GetCropPlanForSeason(seasonID string) (string, error) {
	return "", ErrGreenhouseHasMultipleCrops
}

func (g *GreenhouseArea) IsConfiguredForSeason(seasonID string) bool {
	_, ok := g.Seasons[seasonID]
	return ok
}

func (g *GreenhouseArea) ConfigureForSeason(seasonID string, config AreaConfig) error {
	return g.ConfigureForSeason(seasonID, config.Name, config.Geometry, config.Metadata)
}
