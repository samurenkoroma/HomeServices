package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/spatial"
	"time"
)

// Bed - грядка (в теплице или участке)
type Bed struct {
	aggregate.BaseAggregate

	ID        string
	FarmRefID string
	ParentID  string // ID блока или теплицы (из growing)
	Name      string
	Geometry  spatial.GeoJSON
	Area      float64

	Seasons         map[string]SeasonConfig
	CurrentSeasonID string
}

func NewBed(parentID string, name string, geom spatial.GeoJSON) *Bed {
	return &Bed{
		ID:       generateID(),
		ParentID: parentID,
		Name:     name,
		Geometry: geom,
		Area:     calculateArea(geom),
		Seasons:  make(map[string]SeasonConfig),
	}
}

// ConfigureForSeason - настроить грядку на сезон
func (b *Bed) ConfigureForSeason(
	seasonID string,
	name string,
	geom spatial.GeoJSON,
	cropPlanID string,
	metadata map[string]interface{},
) error {
	if _, exists := b.Seasons[seasonID]; exists {
		return ErrSeasonAlreadyConfigured
	}

	config := SeasonConfig{
		SeasonID:   seasonID,
		Name:       name,
		Geometry:   geom,
		Area:       CalculateArea(geom),
		CropPlanID: &cropPlanID,
		Metadata:   metadata,
		ValidFrom:  time.Now(),
	}

	b.Seasons[seasonID] = config
	b.CurrentSeasonID = seasonID

	b.AddEvent(BedConfigured{
		BedID:      b.ID,
		SeasonID:   seasonID,
		CropPlanID: cropPlanID,
		Area:       config.Area,
	})

	return nil
}

func (b *Bed) GetID() string                { return b.ID }
func (b *Bed) GetFarmRefID() string         { return b.FarmRefID }
func (b *Bed) GetType() AreaType            { return AreaTypeBed }
func (b *Bed) GetName() string              { return b.Name }
func (b *Bed) GetGeometry() spatial.GeoJSON { return b.Geometry }
func (b *Bed) GetArea() float64             { return b.Area }
func (b *Bed) GetCurrentSeasonID() string   { return b.CurrentSeasonID }
func (b *Bed) HasBlocks() bool              { return false }
func (b *Bed) GetBlocks() []string          { return []string{} }

func (b *Bed) GetSeasonConfig(seasonID string) (*SeasonConfig, error) {
	if config, exists := b.Seasons[seasonID]; exists {
		return &config, nil
	}
	return nil, ErrSeasonConfigNotFound
}

func (b *Bed) GetCropPlanForSeason(seasonID string) (string, error) {
	config, err := b.GetSeasonConfig(seasonID)
	if err != nil {
		return "", err
	}
	if config.CropPlanID == nil {
		return "", ErrNoCropPlanConfigured
	}
	return *config.CropPlanID, nil
}

func (b *Bed) IsConfiguredForSeason(seasonID string) bool {
	_, ok := b.Seasons[seasonID]
	return ok
}

func (b *Bed) ConfigureForSeason(seasonID string, config AreaConfig) error {
	if config.CropPlanID == nil {
		return ErrCropPlanRequiredForBed
	}
	return b.ConfigureForSeason(seasonID, config.Name, config.Geometry, *config.CropPlanID, config.Metadata)
}
