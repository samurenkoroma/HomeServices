package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/spatial"
	"time"
)

// Block - участок внутри поля (существует только в рамках сезона)
type Block struct {
	aggregate.BaseAggregate

	ID            string
	FarmRefID     string // Ссылка на ID в farm (если есть)
	ParentFieldID string // ID поля из growing (FieldArea)
	Name          string
	Geometry      spatial.GeoJSON
	Area          float64

	Seasons         map[string]SeasonConfig
	CurrentSeasonID string
	ChildBeds       []string // Грядки внутри участка
}

func NewBlock(parentFieldID string, name string, geom spatial.GeoJSON) *Block {
	return &Block{
		ID:            generateID(),
		ParentFieldID: parentFieldID,
		Name:          name,
		Geometry:      geom,
		Area:          calculateArea(geom),
		Seasons:       make(map[string]SeasonConfig),
		ChildBeds:     []string{},
	}
}

// ConfigureForSeason - настроить участок на сезон
func (b *Block) ConfigureForSeason(
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
		Area:       calculateArea(geom),
		CropPlanID: &cropPlanID,
		BlockIDs:   []string{},
		Metadata:   metadata,
		ValidFrom:  time.Now(),
	}

	b.Seasons[seasonID] = config
	b.CurrentSeasonID = seasonID

	b.AddEvent(BlockConfigured{
		BlockID:    b.ID,
		SeasonID:   seasonID,
		CropPlanID: cropPlanID,
		Area:       config.Area,
	})

	return nil
}

// AddBed - добавить грядку в участок
func (b *Block) AddBed(seasonID string, bedID string) error {
	config, exists := b.Seasons[seasonID]
	if !exists {
		return ErrSeasonConfigNotFound
	}

	config.BlockIDs = append(config.BlockIDs, bedID)
	b.Seasons[seasonID] = config
	b.ChildBeds = append(b.ChildBeds, bedID)

	return nil
}

func (b *Block) GetID() string                { return b.ID }
func (b *Block) GetFarmRefID() string         { return b.FarmRefID }
func (b *Block) GetType() AreaType            { return AreaTypeBlock }
func (b *Block) GetName() string              { return b.Name }
func (b *Block) GetGeometry() spatial.GeoJSON { return b.Geometry }
func (b *Block) GetArea() float64             { return b.Area }
func (b *Block) GetCurrentSeasonID() string   { return b.CurrentSeasonID }
func (b *Block) HasBlocks() bool              { return false }
func (b *Block) GetBlocks() []string          { return []string{} }

func (b *Block) GetSeasonConfig(seasonID string) (*SeasonConfig, error) {
	if config, exists := b.Seasons[seasonID]; exists {
		return &config, nil
	}
	return nil, ErrSeasonConfigNotFound
}

func (b *Block) GetCropPlanForSeason(seasonID string) (string, error) {
	config, err := b.GetSeasonConfig(seasonID)
	if err != nil {
		return "", err
	}
	if config.CropPlanID == nil {
		return "", ErrNoCropPlanConfigured
	}
	return *config.CropPlanID, nil
}

func (b *Block) IsConfiguredForSeason(seasonID string) bool {
	_, ok := b.Seasons[seasonID]
	return ok
}

func (b *Block) ConfigureForSeason(seasonID string, config AreaConfig) error {
	if config.CropPlanID == nil {
		return ErrCropPlanRequiredForBlock
	}
	return b.ConfigureForSeason(seasonID, config.Name, config.Geometry, *config.CropPlanID, config.Metadata)
}
